package utils

import (
	"github.com/TanmayPatil105/go-chat/router"
	"github.com/robfig/cron/v3"
)

var job *cron.Cron

func SetupCronJob() {
	job = cron.New()

	job.AddFunc("@hourly", func() { router.CleanUp() })

	// Start cron job
	job.Start()
}

func StopCronJob() {
	// Stop cron job
	job.Stop()
}
