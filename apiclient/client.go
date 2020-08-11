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
	URI        string
	httpClient httpClient
	bodyReader bodyReader
}

// New creates a pointer to a new instance of ApiClient.
func New(uri string) *ApiClient {
	return &ApiClient{
		URI:        uri,
		httpClient: new(http.Client),
		bodyReader: new(defaultBodyReader),
	}
}

// DoPost takes a well-formatted JSON message and sends it to AnkiConnect.
func (a *ApiClient) DoPost(body string) (string, error) {
	return a.DoHTTP(http.MethodPost, body)
}

func (a *ApiClient) DoHTTP(method string, body string) (string, error) {
	var message string
	var resp *http.Response
	var err error

	if method == http.MethodPost {
		mimeType := "application/json"
		resp, err = a.httpClient.Post(a.URI, mimeType, bytes.NewBufferString(body))
	} else if method == http.MethodGet {
		resp, err = a.httpClient.Get(a.URI)
	}
	if err != nil {
		return message, err
	}

	if resp.StatusCode != 200 {
		errMessage := fmt.Sprintf(wrongStatusErrorFormat, resp.StatusCode)
		err = errors.New(errMessage)
		return message, err
	}

	if resp.Body != nil {
		defer resp.Body.Close()
		body, readErr := a.bodyReader.ReadAll(resp.Body)
		if readErr != nil {
			err = readErr
			return message, err
		}

		message = string(body)
	}

	return message, err
}

//go:generate mockery -name httpClient -filename mock_http_test.go -structname MockHTTPClient -output . -inpkg
type httpClient interface {
	Post(uri, contentType string, body io.Reader) (*http.Response, error)
	Get(uri string) (*http.Response, error)
}

type bodyReader interface {
	ReadAll(r io.Reader) ([]byte, error)
}

type defaultBodyReader struct{}

func (*defaultBodyReader) ReadAll(r io.Reader) ([]byte, error) {
	return ioutil.ReadAll(r)
}
