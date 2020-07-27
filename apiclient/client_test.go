package apiclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestDeckNamesAndIds(t *testing.T) {

	// setup

	// expected values
	expected_action := `{"action":"deckNamesAndIds","version":6}`

	// set up mock client and response
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockHTTPClient := NewMockHTTPClient(mockCtrl)
	http_client = mockHTTPClient
	response_json := `{ "result": {"Default": 1}, "error": null }`
	body := ioutil.NopCloser(bytes.NewReader([]byte(response_json)))
	response := &http.Response{Body: body}
	action_bytes := bytes.NewBufferString(expected_action)
	mockHTTPClient.EXPECT().Post(uri, mime_type, action_bytes).Times(1).Return(response, nil)

	// execute

	DeckNamesAndIds()

}
