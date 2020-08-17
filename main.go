package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"

	_ "github.com/mkevac/debugcharts"
	"github.com/zu1k/proxypool/api"
	"github.com/zu1k/proxypool/app"
	"github.com/zu1k/proxypool/proxy"
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
	} else {
		err := app.InitConfigAndGetters(configFilePath)
		if err != nil {
			fmt.Println(err)
		}
	}
	proxy.InitGeoIpDB()

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
