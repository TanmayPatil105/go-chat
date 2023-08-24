package utils

import (
	"github.com/TanmayPatil105/go-chat/router"
	"github.com/robfig/cron/v3"
)

func SetupCronJob() {
	c := cron.New()

	c.AddFunc("@hourly", router.CleanUp())
}
