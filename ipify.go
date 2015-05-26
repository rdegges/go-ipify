package ipify

import (
	"errors"
	"github.com/jpillora/backoff"
	"io/ioutil"
	"net/http"
	"strconv"
	"time"
)

// GetIp queries the ipify service (http://www.ipify.org) to retrieve this
// machine's public IP address.
func GetIp() (string, error) {
	b := &backoff.Backoff{
		Jitter: true,
	}
	client := &http.Client{}

	req, err := http.NewRequest("GET", API_URI, nil)
	if err != nil {
		return "", errors.New("Received an invalid status code from ipify: 500. The service might be experiencing issues.")
	}

	req.Header.Add("User-Agent", USER_AGENT)

	for tries := 0; tries < MAX_TRIES; tries++ {
		resp, err := client.Do(req)
		if err != nil {
			d := b.Duration()
			time.Sleep(d)
			continue
		}

		defer resp.Body.Close()

		ip, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", errors.New("Received an invalid status code from ipify: 500. The service might be experiencing issues.")
		}

		if resp.StatusCode != 200 {
			return "", errors.New("Received an invalid status code from ipify: " + strconv.Itoa(resp.StatusCode) + ". The service might be experiencing issues.")
		}

		return string(ip), nil
	}

	return "", errors.New("The request failed because it wasn't able to reach the ipify service. This is most likely due to a networking error of some sort.")
}
