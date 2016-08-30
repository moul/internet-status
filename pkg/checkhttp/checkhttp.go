package checkhttp

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Check checks for the presence of captive portal
func CheckFacebookHTTP() error {
	// configure HTTP client
	timeout := time.Duration(2 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	// perform HTTP request
	resp, err := client.Get("http://www.facebook.com/")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 302 {
		return fmt.Errorf("Invalid status code: %d", resp.StatusCode)
	}

	// read the Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// check for expected body
	if len(body) > 0 {
		return fmt.Errorf("http://www.facebook.com, body should be empty")
	}

	return nil
}
