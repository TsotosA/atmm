package main

import "github.com/robfig/cron/v3"

func GetCronJobs() []CronJob {
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
	return res
}

func RestartCronJobs(c *cron.Cron) {
	//zap.S().Debugf("RestartCronJobs() before stoping cron jobs. mwg: %v, cronJobs: %v", mwg, len(C.Entries()))
	C.Stop()
	var oldEntries []cron.EntryID
	for _, entry := range c.Entries() {
		oldEntries = append(oldEntries, entry.ID)
	}
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
	for _, entry := range oldEntries {
		c.Remove(entry)
		mwg.Done()
	}
	//zap.S().Debugf("RestartCronJobs() before starting cron jobs. mwg: %v, cronJobs: %v", mwg, len(C.Entries()))
	C.Start()
}
