package apiclient

import (
	"bytes"
	"io"
	"io/ioutil"
	"net/http"
)

var deckNamesAndIdsBody string = `{ "action": "deckNamesAndIds", "version": 6 }`
var uri = "http://localhost:8765"
var mime_type = "application/json"
var http_client HTTPClient = new(http.Client)

func DeckNamesAndIds() (string, error) {
	return callUriAndReturnBody(deckNamesAndIdsBody)
}

func callUriAndReturnBody(action string) (string, error) {
	var message string

	resp, err := http_client.Post(uri, mime_type, bytes.NewBufferString(action))
	if err != nil {
		return message, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
		body, readErr := ioutil.ReadAll(resp.Body)
		if readErr != nil {
			return message, readErr
		}

		message = string(body)
	}

	return message, err
}

type HTTPClient interface {
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
}
