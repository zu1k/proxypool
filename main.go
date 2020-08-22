package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strings"

	"github.com/zu1k/proxypool/config"

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

	envConfigFilePath := os.Getenv("CONFIG_FILE")
	if envConfigFilePath == "" {
		envConfigFilePath = "source.yaml"
	}

	if configFilePath != "" {
		initConfigFile(configFilePath)
	} else {
		initConfigFile(envConfigFilePath)
	}

	proxy.InitGeoIpDB()

	go cron.Cron()
	fmt.Println("Do the first crawl...")
	go app.CrawlGo()
	api.Run()
}

func initConfigFile(path string) {
	fmt.Println("Config file:", path)
	if strings.HasPrefix(path, "http") {
		config.Url = path
		config.NeedFetch = true
	} else {
		err := app.InitConfigAndGetters(configFilePath)
		if err != nil {
			fmt.Errorf("Config file not found")
			os.Exit(2)
		}
	}
}

func pprof() {
	ip := "127.0.0.1:6060"
	if err := http.ListenAndServe(ip, nil); err != nil {
		fmt.Printf("start pprof failed on %s\n", ip)
	}
}
