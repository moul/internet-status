package captiveportal

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const expectedBody = "<HTML><HEAD><TITLE>Success</TITLE></HEAD><BODY>Success</BODY></HTML>"

// Check checks for the presence of captive portal
func Check() error {
	resp, err := http.Get("http://www.apple.com/library/test/success.html")
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	if string(body) != expectedBody {
		return fmt.Errorf("body differs from expected result")
	}

	return nil
}
