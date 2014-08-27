package main

import ("encoding/json"
	"net/http"
	"fmt"
	"log"
	"flag"
)

const wtfipurl = "https://wtfismyip.com/json"

type IPInfo struct {
	YourFuckingIPAddress string
	YourFuckingLocation string
	YourFuckingHostname string
	YourFuckingISP string
}


func main() {
	var doAddr, doHost, doLoc, doISP, doAll bool

	flag.BoolVar(&doAddr, "i", false, "Show my fucking IP Address")
	flag.BoolVar(&doHost, "h", false, "Show my fucking Hostname")
	flag.BoolVar(&doLoc, "l", false, "Show my fucking Location")
	flag.BoolVar(&doISP, "p", false, "Show my fucking ISP")
	flag.BoolVar(&doAll, "a", false, "Show me fucking Everything")
	flag.Parse()

	client := &http.Client{}
	req, err := http.NewRequest("GET", wtfipurl, nil)

	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Add("Accept", "application/json")
	resp, err := client.Do(req)

	if err != nil {
		log.Fatalln(err)
	}

	defer resp.Body.Close()


	ipinfo := IPInfo{}
	decoder := json.NewDecoder(resp.Body)
	err = decoder.Decode(&ipinfo)

	if err != nil {
		log.Fatalln(err)
	}

	if !doAddr && !doHost && !doISP && !doLoc {
		doAll = true
	}

	if doAll {
		fmt.Println("Fucking Addr:\t", ipinfo.YourFuckingIPAddress)
		fmt.Println("Fucking Host:\t", ipinfo.YourFuckingHostname)
		fmt.Println("Fucking ISP :\t", ipinfo.YourFuckingISP)
		fmt.Println("Fucking Loc :\t", ipinfo.YourFuckingLocation)
		return
	}

	if doAddr {
		fmt.Println(ipinfo.YourFuckingIPAddress)
	}
	if doHost {
		fmt.Println(ipinfo.YourFuckingHostname)
	}
	if doISP {
		fmt.Println(ipinfo.YourFuckingISP)
	}
	if doLoc {
		fmt.Println(ipinfo.YourFuckingLocation)
	}
}

