package apiclient

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
)

// var deckNamesAndIdsBody string = `{ "action": "deckNamesAndIds", "version": 6 }`
var deckNamesAndIdsAction string = "deckNamesAndIds"

var uri = "http://localhost:8765"
var mime_type = "application/json"
var http_client HTTPClient = new(http.Client)

func DeckNamesAndIds() (string, error) {

	deckNamesAndIdsBody := createActionString(deckNamesAndIdsAction)
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

type request_body struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
}

func createActionString(action string) string {
	s := &request_body{Action: action, Version: 6}
	a, err := json.Marshal(s)
	if err != nil {
		panic("Error: could not convert action to json")
	}
	return string(a)
}
