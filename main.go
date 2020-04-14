package main

import (
	"fmt"
	"net"
	"time"

	uerr "github.com/ucloud/ucloud-sdk-go/ucloud/error"
)

func main() {
	config.getConf()
	initUcloud()
	for {
		process()
		time.Sleep(time.Hour)
	}
}

func process() {
	fmt.Println("Start to process...")

	fmt.Println("Delete old GlobalSSH...")
	for _, id := range config.Instances {
		deleteGlobalSSH(id)
	}

	fmt.Println("Checking Github ip...")
	ips, _ := fetchIPs()

	fmt.Println("Creating ucloud GlobalSSH...")
	DomainsNew := make([]string, 0)
	instances := make([]string, 0)
	for _, ip := range ips {
		ip = parseIP(ip)
		if ip != "" {
			domain, instance, err := newGlobalSSH(ip, area(ip))
			if err != nil {
				e := err.(uerr.Error)
				if e.Code() == 33981 {
					fmt.Println("Ucloud ip", ip, "already exists")
					instances = append(instances, instance)
				} else {
					fmt.Println("Ucloud error:", err.Error())
				}
				continue
			} else {
				fmt.Println("Ucloud new GlobalSSH:", domain)
				DomainsNew = append(DomainsNew, domain)
				instances = append(instances, instance)
			}
		}
	}
	config.Instances = instances
	config.save()

	fmt.Println("\nSleeping 2min for dns ...")
	time.Sleep(time.Minute * 2)
	fmt.Println("\nNow Lookup Domain...")
	newDnsRecordIps := make([]string, 0)
	for _, domain := range DomainsNew {
		ips, err := net.LookupIP(domain)
		if err != nil {
			fmt.Println("DNS lookup error:", err.Error())
		} else {
			for _, newip := range ips {
				newDnsRecordIps = append(newDnsRecordIps, newip.String())
			}
		}
	}
	setDnsRecords(newDnsRecordIps)
	fmt.Println("End...")
}
