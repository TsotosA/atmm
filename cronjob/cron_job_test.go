package cronjob

//func TestRestartCronJobs(t *testing.T) {
//	x := []model.CronJob{
//		{
//			ScheduleToUse: config.Conf.Cron,
//			MethodToRun: func() {
//
//			},
//		},
//		{
//			ScheduleToUse: config.Conf.CheckForUpdatesInterval,
//			MethodToRun: func() {
//
//			},
//		},
//		{
//			ScheduleToUse: config.Conf.ScanForMovieInterval,
//			MethodToRun: func() {
//
//			},
//		},
//		{
//			ScheduleToUse: config.Conf.DbBucketsCleanupInterval,
//			MethodToRun: func() {
//
//			},
//		},
//	}
//	mwg := sync.WaitGroup{}
//	C := cron.New()
//	config.ConfigInit(C, &mwg, x)
//	cronJobs := x
//	initialNumber := 0
//	afterRestartNumber := 0
//
//	t.Run("same number of jobs", func(t *testing.T) {
//		for _, job := range cronJobs {
//			if job.ScheduleToUse != "" {
//				_, err := C.AddFunc(job.ScheduleToUse, job.MethodToRun)
//				if err != nil {
//					//return
//				}
//			}
//		}
//		initialNumber = len(C.Entries())
//		RestartCronJobs(C, &mwg, cronJobs)
//		C.Stop()
//		afterRestartNumber = len(C.Entries())
//		if initialNumber != afterRestartNumber {
//			t.Errorf("same # = %+v, want %+v", afterRestartNumber, initialNumber)
//		}
//	})
//
//	t.Run("different ids", func(t *testing.T) {
//		for _, job := range cronJobs {
//			if job.ScheduleToUse != "" {
//				_, err := C.AddFunc(job.ScheduleToUse, job.MethodToRun)
//				if err != nil {
//					//return
//				}
//			}
//		}
//		RestartCronJobs(C, &mwg, cronJobs)
//		C.Stop()
//		t.Logf("%v", C.Entries())
//		//if C.Entries()[0].ID != 3 {
//		//	t.Errorf("same # = %+v, want %+v", afterRestartNumber, initialNumber)
//		//}
//	})
//
//}
