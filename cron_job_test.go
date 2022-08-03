package main

import (
	"testing"
)

func TestGetCronJobs(t *testing.T) {

	ConfigInit()

	res := []CronJob{
		{
			ScheduleToUse: Conf.Cron,
			MethodToRun:   RunTvShowsAsCronJob,
		},
		{
			ScheduleToUse: Conf.CheckForUpdatesInterval,
			MethodToRun:   RunUpdateAsCronJob,
		},
		{
			ScheduleToUse: Conf.ScanForMovieInterval,
			MethodToRun:   RunMoviesAsCronJob,
		},
		{
			ScheduleToUse: Conf.DbBucketsCleanupInterval,
			MethodToRun:   RunDbCleanupAsCronJob,
		},
	}

	t.Run("GetCronJobs", func(t *testing.T) {
		if got := GetCronJobs(); len(res) != len(got) {
			t.Errorf("GetCronJobs() = %+v, want %+v", got, res)
		}
	})
}

func TestRestartCronJobs(t *testing.T) {
	ConfigInit()
	cronJobs := GetCronJobs()
	initialNumber := 0
	afterRestartNumber := 0

	t.Run("same number of jobs", func(t *testing.T) {
		for _, job := range cronJobs {
			if job.ScheduleToUse != "" {
				_, err := C.AddFunc(job.ScheduleToUse, job.MethodToRun)
				if err != nil {
					//return
				}
			}
		}
		initialNumber = len(C.Entries())
		RestartCronJobs(C)
		C.Stop()
		afterRestartNumber = len(C.Entries())
		if initialNumber != afterRestartNumber {
			t.Errorf("same # = %+v, want %+v", afterRestartNumber, initialNumber)
		}
	})

}
