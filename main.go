package main

import (
	"flag"
	"fmt"

	"github.com/zu1k/proxypool/api"
	"github.com/zu1k/proxypool/app"
)

func main() {
	filePath := flag.String("c", "", "path to config file: source.yaml")
	flag.Parse()
	if *filePath == "" {
		app.NeedFetchNewConfigFile = true
	} else {
		app.InitConfigAndGetters(*filePath)
	}
	go app.Cron()
	fmt.Println("Do the first crawl...")
	app.CrawlGo()
	api.Run()
}
