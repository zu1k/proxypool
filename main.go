package main

import (
	"flag"
	"fmt"
	"net/http"

	"github.com/zu1k/proxypool/api"
	"github.com/zu1k/proxypool/app"
)

func main() {
	go pprof()

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

func pprof() {
	ip := "127.0.0.1:6060"
	if err := http.ListenAndServe(ip, nil); err != nil {
		fmt.Printf("start pprof failed on %s\n", ip)
	}
}
