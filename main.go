package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"github.com/zu1k/proxypool/internal/cron"

	_ "github.com/mkevac/debugcharts"
	"github.com/zu1k/proxypool/api"
	"github.com/zu1k/proxypool/internal/app"
	"github.com/zu1k/proxypool/pkg/proxy"
)

var (
	debugMode      = false
	configFilePath = ""
)

func main() {
	flag.StringVar(&configFilePath, "c", "", "path to config file: source.yaml")
	flag.BoolVar(&debugMode, "d", false, "debug mode")
	flag.Parse()

	if debugMode {
		go pprof()
	}

	if configFilePath == "" {
		app.NeedFetchNewConfigFile = true
		app.FetchNewConfigFileThenInit()
	} else {
		err := app.InitConfigAndGetters(configFilePath)
		if err != nil {
			fmt.Println(err)
		}
	}
	proxy.InitGeoIpDB()

	go cron.Cron()
	fmt.Println("Do the first crawl...")
	go app.CrawlGo()
	api.Run()
}

func pprof() {
	ip := "127.0.0.1:6060"
	if err := http.ListenAndServe(ip, nil); err != nil {
		fmt.Printf("start pprof failed on %s\n", ip)
	}
}
