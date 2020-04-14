package main

import (
	"fmt"
	"log"

	"github.com/cloudflare/cloudflare-go"
)

func setDnsRecords(ips []string) {
	api, err := cloudflare.NewWithAPIToken(config.Cloudflare.Token)
	if err != nil {
		log.Fatal(err)
	}

	// Fetch the zone ID
	zone, err := api.ZoneIDByName(config.Cloudflare.Zone)
	if err != nil {
		log.Fatal(err)
	}

	originRecords, err := api.DNSRecords(zone, dnsRecord4Delete)
	for _, r := range originRecords {
		_ = api.DeleteDNSRecord(zone, r.ID)
	}

	for _, ip := range ips {
		dnsRecord := cloudflare.DNSRecord{
			Type:     "A",
			Name:     config.Cloudflare.Record,
			Content:  ip,
			Proxied:  false,
			TTL:      120,
			Priority: 0,
		}
		_, err := api.CreateDNSRecord(zone, dnsRecord)
		if err != nil {
			fmt.Println(err.Error())
		} else {
			fmt.Println("Cloudflare new record:", ip)
		}
	}
}

var dnsRecord4Delete = cloudflare.DNSRecord{
	Type: "A",
	Name: config.Cloudflare.Record,
}
