package shodan

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type UsageLimits struct {
	ScanCredits   int `json:"scan_credits"`
	QueryCreadits int `json:"query_credits"`
	MonitoredIps  int `json:"monitored_ips"`
}

type APIInfo struct {
	ScanCredits   int         `json:"scan_credits"`
	UsageLimits   UsageLimits `json:"usage_limits"`
	Plan          string      `json:"plan"`
	Https         bool        `json:"https"`
	Unlocked      bool        `json:"unlocked"`
	QueryCreadits int         `json:"query_credits"`
	MonitoredIps  int         `json:"monitored_ips"`
	UnlockedLeft  int         `json:"unlocked_left"`
	Telnet        bool        `json:"telnet"`
}

// Returns information about your API key
func (s *Client) APIInfo() (*APIInfo, error) {
	res, err := http.Get(fmt.Sprintf("%s/api-info?key=%s", BaseURL, s.apiKey))
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	var ret APIInfo
	if err := json.NewDecoder(res.Body).Decode(&ret); err != nil {
		return nil, err
	}

	return &ret, nil
}
