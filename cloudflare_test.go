package main

import (
	"fmt"
	"log"
	"testing"

	"github.com/cloudflare/cloudflare-go"
)

func TestListRecords(t *testing.T) {
	config.getConf()

	api, err := cloudflare.NewWithAPIToken(config.Cloudflare.Token)
	if err != nil {
		log.Fatal(err)
	}

	// Fetch the zone ID
	zone, err := api.ZoneIDByName(config.Cloudflare.Zone)
	if err != nil {
		log.Fatal(err)
	}

	originRecords, err := api.DNSRecords(zone, cloudflare.DNSRecord{
		Type: "A",
		Name: config.Cloudflare.Record,
	})
	for _, r := range originRecords {
		fmt.Println(r)
	}
}
