package ipify

import (
	"fmt"
	"testing"
)

func TestGetIp(t *testing.T) {
	originalApiUri := API_URI

	ip, err := GetIp()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ip)

	API_URI = "https://api.ipifyyyyyyyyyyyy.org"

	ip, err = GetIp()
	if err == nil || ip != "" {
		t.Error("Request to https://api.ipifyyyyyyyyyyyy.org should have failed, but succeeded.")
	} else {
		fmt.Println(err)
	}

	API_URI = originalApiUri
}
