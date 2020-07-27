package apiclient

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func Test_callUriAndReturnBody_success(t *testing.T) {

	// setup
	expected_action := `{"action":"deckNamesAndIds","version":6}`
	mock_response := create_default_response()
	var mock_err error
	mockHTTPClient := setup_mocks(t)

	// verify (wrong order because of mocking)
	// check to make sure that the method calls the api with the correct request body
	mockHTTPClient.EXPECT().Post(uri, mime_type, bytes.NewBufferString(expected_action)).Times(1).Return(mock_response, mock_err)

	// execute
	callUriAndReturnBody(expected_action)

}

func Test_callUriAndReturnBody_err(t *testing.T) {

	// setup
	expected_action := `{"action":"deckNamesAndIds","version":6}`
	mock_response := create_default_response()
	var mock_err error = errors.New("Simple error for testing")
	mockHTTPClient := setup_mocks(t)

	// assert (wrong order because of mocking)
	// check to make sure that the method calls the api with the correct request body
	mockHTTPClient.EXPECT().Post(uri, mime_type, bytes.NewBufferString(expected_action)).Times(1).Return(mock_response, mock_err)

	// execute
	message, err := callUriAndReturnBody(expected_action)

	// assert
	if message != "" {
		t.Error("Expected message to be null if http client returned an error")
	}
	if err == nil {
		t.Error("Expected to get the http error passed through")
	}

}

func Test_callUriAndReturnBody_not_200(t *testing.T) {

	// setup
	expected_action := `{"action":"deckNamesAndIds","version":6}`
	mock_response := create_default_response()
	mock_response.StatusCode = 500
	var mock_err error
	mockHTTPClient := setup_mocks(t)

	// assert (wrong order because of mocking)
	// check to make sure that the method calls the api with the correct request body
	mockHTTPClient.EXPECT().Post(uri, mime_type, bytes.NewBufferString(expected_action)).Times(1).Return(mock_response, mock_err)

	// execute
	message, err := callUriAndReturnBody(expected_action)

	// assert
	if message != "" {
		t.Errorf("apiclient: Expected message to be empty if http client returned an error, received: %s", message)
	}
	if err == nil {
		t.Error("Expected an error if the http response code was not 200")
	}
	expected_error_message := fmt.Sprintf("apiclient: Received non-200 status code from api endpoint: %d", mock_response.StatusCode)
	if err.Error() != expected_error_message {
		t.Error("Expected a different error message")
	}

}

// checks that the expected action is the body of the reqest that goes out.
func setup_mocks(t *testing.T) (mockHTTPClient *MockHTTPClient) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockHTTPClient = NewMockHTTPClient(mockCtrl)
	Http_client = mockHTTPClient
	return
}

func create_default_response() (response *http.Response) {
	response_json := `{ "result": {"Default": 1}, "error": null }`
	body := ioutil.NopCloser(bytes.NewReader([]byte(response_json)))
	response = &http.Response{Body: body, StatusCode: 200}
	return
}
