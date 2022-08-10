package main

import (
	"github.com/robfig/cron/v3"
	"github.com/tsotosa/atmm/config"
	"github.com/tsotosa/atmm/gconst"
	"github.com/tsotosa/atmm/global"
	"github.com/tsotosa/atmm/model"
	"github.com/tsotosa/atmm/web/api"
	ui_serve "github.com/tsotosa/atmm/web/ui-serve"
	"go.uber.org/zap"
	"sync"
)

var (
	mwg     = sync.WaitGroup{}
	Version = "development"
	C       = cron.New()
)

func main() {
	config.ConfigInit(C, &mwg, GetCronJobs())

	logger, err := InitLogger(config.Conf.LogOutputPath)
	if err != nil {
		panic(err)
	}
	defer func() {
		err := Sync()
		if err != nil {
			zap.S().Errorf("%s", "failed to flush logs before exiting")
		}
	}()
	undo := zap.ReplaceGlobals(logger)

	zap.S().Infof("::: app start :::")
	zap.S().Infof("version:[%s]", Version)
	defer zap.S().Infof("::: app exit :::\n")
	defer undo()

	// >>bbolt
	err = InitBolt("atmm.db", []string{gconst.TvShowEpisodeFilesBucket, gconst.MovieFilesBucket})
	if err != nil {
		zap.S().Fatalf("failed to init db: error is: %v\n", err)
		return
	}
	defer Close()
	// bbolt<<

	cronJobs := GetCronJobs()

	for _, job := range cronJobs {
		if job.ScheduleToUse != "" {
			mwg.Add(1)
			_, err := C.AddFunc(job.ScheduleToUse, job.MethodToRun)
			if err != nil {
				mwg.Done()
				//return
			}
		}
	}

	if config.Conf.ApiPort != 0 {
		mwg.Add(1)
		go api.Up()
	}

	if config.Conf.UiPort != 0 {
		mwg.Add(1)
		go ui_serve.Up()
	}

	C.Start()
	mwg.Wait()
	if err != nil {
		return
	}
}

func RunTvShowsAsCronJob() {
	if global.WaitingSeriesToFinishCopying {
		zap.S().Infof("skipping tv show run due to copying files")
		return
	}
	err := HandleTvShows()
	if err != nil {
		return
	}
}
func RunMoviesAsCronJob() {
	if global.WaitingMoviesToFinishCopying {
		zap.S().Infof("skipping movies run due to copying files")
		return
	}
	err := HandleMovies()
	if err != nil {
		return
	}
}

func RunUpdateAsCronJob() {
	if global.WaitingMoviesToFinishCopying || global.WaitingSeriesToFinishCopying {
		zap.S().Infof("skipping update run due to copying files")
		return
	}
	err := HandleUpdate()
	if err != nil {
		return
	}
}

func RunDbCleanupAsCronJob() {
	err := HandleDbCleanup()
	if err != nil {
		return
	}

}

func GetCronJobs() []model.CronJob {
	res := []model.CronJob{
		{
			ScheduleToUse: config.Conf.Cron,
			MethodToRun:   RunTvShowsAsCronJob,
		},
		{
			ScheduleToUse: config.Conf.CheckForUpdatesInterval,
			MethodToRun:   RunUpdateAsCronJob,
		},
		{
			ScheduleToUse: config.Conf.ScanForMovieInterval,
			MethodToRun:   RunMoviesAsCronJob,
		},
		{
			ScheduleToUse: config.Conf.DbBucketsCleanupInterval,
			MethodToRun:   RunDbCleanupAsCronJob,
		},
	}
	return res
}
