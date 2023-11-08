package ipinfo

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type IPInfo struct {
	Ip       string `json:"ip"`
	Hostname string `json:"hostname"`
	City     string `json:"city"`
	Region   string `json:"region"`
	Country  string `json:"country"`
	Loc      string `json:"loc"`
	Org      string `json:"org"`
	Postal   string `json:"postal"`
	Timezone string `json:"timezone"`
	Readme   string `json:"readme"`
}

func New() *IPInfo {
	return &IPInfo{}
}

// Get retrieves all available ip info, populates the underlying struct, and returns the string containing the actual IP.
func (ip *IPInfo) Get() (string, error) {

	resp, err := http.Get("http://ipinfo.io/json")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(ip); err != nil {
		return "", err
	}

	return fmt.Sprintf("%s\n", ip.Ip), nil
}
