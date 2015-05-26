package ipify

import (
	"testing"
)

func TestGetIp(t *testing.T) {
	originalApiUri := API_URI

	_, err := GetIp()
	if err != nil {
		t.Error(err)
	}

	API_URI = "https://api.ipifyyyyyyyyyyyy.org"

	_, err = GetIp()
	if err == nil {
		t.Error("Request to https://api.ipifyyyyyyyyyyyy.org should have failed, but succeeded.")
	}

	API_URI = originalApiUri
}
