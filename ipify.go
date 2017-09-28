// Package ipify provides a single function for retrieving your computer's
// public IP address from the ipify service: http://www.ipify.org
package ipify

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/jpillora/backoff"
)

// GetIp queries the ipify service (http://www.ipify.org) to retrieve this
// machine's public IP address.  Returns your public IP address as a string, and
// any error encountered.  By default, this function will run using exponential
// backoff -- if this function fails for any reason, the request will be retried
// up to 3 times.
//
// Usage:
//
//		package main
//
//		import (
//			"fmt"
//			"github.com/rdegges/go-ipify"
//		)
//
//		func main() {
//			ip, err := ipify.GetIp()
//			if err != nil {
//				fmt.Println("Couldn't get my IP address:", err)
//			} else {
//				fmt.Println("My IP address is:", ip)
//			}
//		}
func GetIp() (string, error) {
	b := &backoff.Backoff{
		Jitter: true,
	}
	client := &http.Client{}

	req, err := http.NewRequest("GET", API_URI, nil)
	if err != nil {
		return "", wrapError(err)
	}

	req.Header.Set("User-Agent", USER_AGENT)

	for tries := 0; tries < MAX_TRIES; tries++ {
		var resp *http.Response
		resp, err = client.Do(req)
		if err != nil {
			d := b.Duration()
			time.Sleep(d)
			continue
		}

		defer resp.Body.Close()

		var ip []byte
		ip, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			return "", wrapError(err)
		}

		if resp.StatusCode != http.StatusOK {
			description :=
				fmt.Sprintf("Received invalid status code %d from ipify: %s. The service might be experiencing issues.",
					resp.StatusCode, resp.Status)
			return "", errors.New(description)
		}

		return string(ip), nil
	}

	return "", wrapError(err)
}

func wrapError(inner error) error {
	description :=
		fmt.Sprintf("The request failed with error %s. This is most likely due to a networking error of some sort.",
			inner.Error())
	return errors.New(description)
}
