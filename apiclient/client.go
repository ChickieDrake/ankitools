/*
Package apiclient provides a simple interface for making http calls to the AnkiConnect API.

It assumes that AnkiConnect is running on localhost:8765. It does not support all methods at this time.
*/
package apiclient

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

const wrongStatusErrorFormat = "apiclient: Received non-200 status code from api endpoint: %d"

// ApiClient provides a wrapper for making http calls to AnkiConnect.
type ApiClient struct {
	uri        string
	mimeType   string
	httpClient httpClient
	bodyReader bodyReader
}

// New creates a pointer to a new instance of ApiClient.
func New() *ApiClient {
	return &ApiClient{
		uri:        "http://localhost:8765",
		mimeType:   "application/json",
		httpClient: new(http.Client),
		bodyReader: new(defaultBodyReader),
	}
}

// DoAction takes a well-formated JSON message and sends it to AnkiConnect.
func (a *ApiClient) DoAction(body string) (message string, err error) {

	resp, err := a.httpClient.Post(a.uri, a.mimeType, bytes.NewBufferString(body))
	if err != nil {
		return
	}

	if resp.StatusCode != 200 {
		errMessage := fmt.Sprintf(wrongStatusErrorFormat, resp.StatusCode)
		err = errors.New(errMessage)
		return
	}

	if resp.Body != nil {
		defer resp.Body.Close()
		body, readErr := a.bodyReader.ReadAll(resp.Body)
		if readErr != nil {
			err = readErr
			return
		}

		message = string(body)
	}

	return
}

type httpClient interface {
	Post(url, contentType string, body io.Reader) (resp *http.Response, err error)
}

type bodyReader interface {
	ReadAll(r io.Reader) ([]byte, error)
}

type defaultBodyReader struct{}

func (*defaultBodyReader) ReadAll(r io.Reader) ([]byte, error) {
	return ioutil.ReadAll(r)
}
