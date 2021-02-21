package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type ApiInfo struct {
	ScanCredits  int         `json:"scan_credits"`
	QueryCredits int         `json:"query_credits"`
	MonitoredIps int         `json:"monitored_ips"`
	Telnet       bool        `json:"telnet"`
	Plan         string      `json:"plan"`
	Https        bool        `json:"https"`
	Unlocked     int         `json:"unlocked_left"`
	Usage        UsageLimits `json:"usage_limits"`
}

type UsageLimits struct {
	ScanCredits  int `json:"scan_credits"`
	QueryCredits int `json:"query_credits"`
	MonitoredIps int `json:"monitored_ips"`
}

func (client *Client) APIInfo() (*ApiInfo, error) {
	resp, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseUrl, client.apiKey))
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var res ApiInfo
	if err := json.NewDecoder(resp.Body).Decode(&res); err != nil {
		return nil, err
	}
	fmt.Println(resp.Body)
	return &res, err
}
