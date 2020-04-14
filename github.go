package main

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"net/http"
	"time"
)

func fetchIPs() (ips []string, err error) {
	http.DefaultClient.Timeout = time.Second * 10
	//proxy := func(_ *http.Request) (*url.URL, error) {
	//	return url.Parse("http://127.0.0.1:7890")
	//}
	//http.DefaultClient.Transport = &http.Transport{
	//	Proxy: proxy,
	//}

	resp, err := http.Get("https://api.github.com/meta")
	if err != nil {
		fmt.Println("Error:", err.Error())
		return
	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	//fmt.Println(string(s))
	data := make(map[string]interface{})
	err = json.Unmarshal(s, &data)
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}
	ipsi := data["web"].([]interface{})
	ips = make([]string, 0)
	for _, ipr := range ipsi {
		ips = append(ips, ipr.(string))
	}
	return
}

func parseIP(ipr string) string {
	ip, ipnet, err := net.ParseCIDR(ipr)
	if err != nil {
		return ""
	}
	if hex.EncodeToString(ipnet.Mask) == "ffffffff" {
		return ip.String()
	}
	return ""
}
