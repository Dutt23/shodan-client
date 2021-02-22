package shodan

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
)

type Host struct {
	Info         string       `json:"info"`
	Product      string       `json:"product"`
	Hash         int          `json:"hash"`
	IP           int          `json:"ip"`
	Isp          string       `json:"isp"`
	Transport    string       `json:"transport"`
	HostNames    []string     `json:"hostnames"`
	Cpe          []string     `json:"cpe"`
	Data         string       `json:"data"`
	Asn          string       `json:"asn"`
	Port         int          `json:"port"`
	Version      string       `json:"version"`
	TimeStamp    string       `json:"timestamp"`
	Domains      []string     `json:"domains"`
	Org          string       `json:"org"`
	OS           string       `json:"os"`
	IPString     string       `json:"ip_str"`
	Shodan       Shodan       `json:"shodan"`
	HostLocation HostLocation `json:"location"`
}

type Shodan struct {
	Crawler string        `json:"cralwer"`
	Ptr     string        `json:"ptr"`
	ID      string        `json:"id"`
	Module  string        `json:"module"`
	Options []interface{} `json:"options"`
}

type HostLocation struct {
	City         string  `json:"city"`
	RegionCode   string  `json:"region_code"`
	AreaCode     string  `json:"area_code"`
	Longitude    float32 `json:"longitude"`
	CountryCode3 string  `json:"country_code3"`
	Latitude     float32 `json:"latitude"`
	PostalCode   string  `json:"postal_code"`
	DmaCode      int     `json:"dma_code"`
	CountryCode  string  `json:"country_code"`
	CountryName  string  `json:"country_name"`
}

type Matches struct {
	Matches []Host `json:"matches"`
	Total   int    `json:"total"`
}

func (client *Client) HostSearch(q string) (*Matches, error) {
	resp, err := http.Get(fmt.Sprintf("%s/shodan/host/search?key=%s&query=%s", BaseUrl, client.apiKey, q))
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println(resp.Request)
		return nil, errors.New("Expected status code 200 got " + strconv.Itoa(resp.StatusCode))
	}
	defer resp.Body.Close()
	var response Matches
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	return &response, err
}
