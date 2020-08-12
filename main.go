package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/zu1k/proxypool/config"

	"github.com/zu1k/proxypool/api"
	"github.com/zu1k/proxypool/app"
)

func main() {
	filePath := flag.String("c", "source.yaml", "path to config file: source.yaml")
	flag.Parse()
	c, err := config.Parse(*filePath)
	if err != nil {
		fmt.Println("Error: ", err.Error())
		os.Exit(1)
	}
	if c == nil {
		fmt.Println("Error: no sources")
		os.Exit(2)
	}

	app.InitGetters(c.Sources)
	go app.Cron()
	fmt.Println("Do the first crawl...")
	app.CrawlGo()
	api.Run()
}
