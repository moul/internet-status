package captiveportal

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const expectedBody = "<HTML><HEAD><TITLE>Success</TITLE></HEAD><BODY>Success</BODY></HTML>"

// Check checks for the presence of captive portal
func Check() error {
	// configure HTTP client
	timeout := time.Duration(2 * time.Second)
	client := http.Client{
		Timeout: timeout,
	}

	// perform HTTP request
	resp, err := client.Get("http://www.apple.com/library/test/success.html")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// read the Body
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	// check for expected body
	if string(body) != expectedBody {
		return fmt.Errorf("body differs from expected result")
	}

	return nil
}
