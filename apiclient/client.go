package apiclient

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

var uri = "http://localhost:8765"
var mime_type = "application/json"
var Http_client HTTPClient = new(http.Client)

func callUriAndReturnBody(body string) (message string, err error) {

	resp, err := Http_client.Post(uri, mime_type, bytes.NewBufferString(body))
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		err_message := fmt.Sprintf("apiclient: Received non-200 status code from api endpoint: %d", resp.StatusCode)
		err = errors.New(err_message)
		return
	}

	if resp.Body != nil {
		defer resp.Body.Close()
		body, readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			err = readErr
			return
		}

		message = string(body)
	}

	return
}

type HTTPClient interface {
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
}
