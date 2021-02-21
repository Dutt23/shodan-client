package main

import (
	"fmt"
	"os"
	"shodan-client/shodan"
)

func main() {
	apiKey := os.Getenv("SHODAN_API_KEY")
	s := shodan.New(apiKey)
	info, err := s.APIInfo()
	fmt.Printf(
		"Query Credits: %d\nScan Credits: %d\n\n", info.QueryCredits,
		info.ScanCredits)
	fmt.Println(err)
}
