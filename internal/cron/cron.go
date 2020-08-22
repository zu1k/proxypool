package cron

import (
	"github.com/jasonlvhit/gocron"
	"github.com/zu1k/proxypool/internal/app"
)

func Cron() {
	_ = gocron.Every(15).Minutes().Do(crawlTask)
	<-gocron.Start()
}

func crawlTask() {
	_ = app.InitConfigAndGetters("")
	app.CrawlGo()
}
