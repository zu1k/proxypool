package cloudflare

import (
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/zu1k/proxypool/config"
)

func test() {
	api, err := cloudflare.New(config.Config.CFKey, config.Config.CFKey)
	if err != nil {
		log.Fatal(err)
	}

	// Fetch the zone ID
	id, err := api.ZoneIDByName(config.Config.Domain)
	if err != nil {
		log.Fatal(err)
	}

	// Fetch zone details
	zone, err := api.ZoneDetails(id)
	if err != nil {
		log.Fatal(err)
	}
	// Print zone details
	fmt.Println(zone)
}
