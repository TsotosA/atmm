package cronjob

import (
	"github.com/robfig/cron/v3"
	"github.com/tsotosa/atmm/model"
	"sync"
)

func RestartCronJobs(c *cron.Cron, mwg *sync.WaitGroup, jobs []model.CronJob) {
	//zap.S().Debugf("RestartCronJobs() before stopping cron jobs. mwg: %v, cronJobs: %v", mwg, len(C.Entries()))
	c.Stop()
	var oldEntries []cron.EntryID
	for _, entry := range c.Entries() {
		oldEntries = append(oldEntries, entry.ID)
	}
	cronJobs := jobs
	for _, job := range cronJobs {
		if job.ScheduleToUse != "" {
			mwg.Add(1)
			_, err := c.AddFunc(job.ScheduleToUse, job.MethodToRun)
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
	c.Start()
}
