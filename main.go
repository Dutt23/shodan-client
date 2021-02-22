package main

import (
	"fmt"
	"log"
	"os"
	"shodan-client/shodan"
)

func main() {
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	fmt.Printf("%+v\n", info)
	fmt.Printf(
		"Query Credits: %d\nScan Credits: %d\n\n", info.QueryCredits,
		info.ScanCredits)
	// if len(os.Args) != 2 {
	// 	log.Fatalln("Usage : shodan searchterm")
	// }
	// for item, _ := range os.Args {
	// 	fmt.Println(item)
	// }

	hosts, err := s.HostSearch("linux")
	if err != nil {
		log.Panicln(err)
	}
	for _, host := range hosts.Matches {
		fmt.Printf("%18s%8d\n", host.IPString, host.Port)
	}
}
