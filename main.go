package main

import (
	"fmt"

	"github.com/zu1k/proxypool/api"
	"github.com/zu1k/proxypool/app"
)

func main() {
	fmt.Println("Do the first crawl...")
	app.CrawlTGChannel()
	api.Run()
}
