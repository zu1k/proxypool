package app

import (
	"github.com/jasonlvhit/gocron"
)

func Cron() {
	gocron.Every(10).Minutes().Do(CrawlTGChannel)

	<-gocron.Start()
}
