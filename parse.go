package main

import (
	"fmt"
	"log"
	"net"

	"github.com/oschwald/geoip2-golang"
)

var db *geoip2.Reader

func init() {
	var err error
	db, err = geoip2.Open("geoip.mmdb")
	if err != nil {
		log.Fatal(err)
	}
}

func area(ipstring string) string {
	// If you are using strings that may be invalid, check that ip is not nil
	ip := net.ParseIP(ipstring)
	record, err := db.Country(ip)
	if err != nil {
		fmt.Println("Groip error:", err.Error())
	}
	country := record.Country.Names["zh-CN"]
	area, found := closeCountry[country]
	if found {
		return area
	} else {
		return "香港"
	}
}

var closeCountry = map[string]string{
	"日本":   "东京",
	"新加坡":  "新加坡",
	"韩国":   "东京",
	"德国":   "法兰克福",
	"巴西":   "洛杉矶",
	"澳大利亚": "香港",
	"印度":   "香港",
	"美国":   "洛杉矶",
	"荷兰":   "法兰克福",
}
