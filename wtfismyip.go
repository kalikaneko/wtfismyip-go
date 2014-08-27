package main

import ("encoding/json"
	"net/http"
	"fmt"
	"log"
)

const wtfipurl = "https://wtfismyip.com/json"

type IPInfo struct {
	YourFuckingIPAddress string
	YourFuckingLocation string
	YourFuckingHostname string
	YourFuckingISP string
}

func main() {
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
	fmt.Println("Fucking Addr:\t", ipinfo.YourFuckingIPAddress)
	fmt.Println("Fucking Host:\t", ipinfo.YourFuckingHostname)
	fmt.Println("Fucking ISP :\t", ipinfo.YourFuckingISP)
	fmt.Println("Fucking Loc :\t", ipinfo.YourFuckingLocation)
}

