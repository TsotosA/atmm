package main

import (
	"github.com/robfig/cron/v3"
	"go.uber.org/zap"
	"sync"
)

var (
	mwg     = sync.WaitGroup{}
	Version = "development"
	C       = cron.New()
)

func main() {
	ConfigInit()

	logger, err := InitLogger(Conf.LogOutputPath)
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
	err = InitBolt("atmm.db", []string{TvShowEpisodeFilesBucket, MovieFilesBucket})
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

	C.Start()
	mwg.Wait()
	if err != nil {
		return
	}
}

func RunTvShowsAsCronJob() {
	if WaitingSeriesToFinishCopying {
		zap.S().Infof("skipping tv show run due to copying files")
		return
	}
	err := HandleTvShows()
	if err != nil {
		return
	}
}
func RunMoviesAsCronJob() {
	if WaitingMoviesToFinishCopying {
		zap.S().Infof("skipping movies run due to copying files")
		return
	}
	err := HandleMovies()
	if err != nil {
		return
	}
}

func RunUpdateAsCronJob() {
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
