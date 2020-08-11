package main

import (
	"github.com/zu1k/proxypool/api"
	"github.com/zu1k/proxypool/app"
)

func main() {
	app.CrawlTGChannel()
	api.Run()
}
