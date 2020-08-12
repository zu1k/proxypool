package main

import (
	"fmt"

	"github.com/zu1k/proxypool/api"
	"github.com/zu1k/proxypool/app"
)

func main() {
	go app.Cron()
	fmt.Println("Do the first crawl...")
	app.CrawlTGChannel()
	api.Run()
}
