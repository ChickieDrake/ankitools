package apiclient

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"io"
	"io/ioutil"
	"net/http"
	"testing"
)

const expectedAction = `{"action":"deckNamesAndIds","version":6}`
const bodyReadingErrorMessage = "error reading body"
const mockErrorMessage = "mock: simple error for testing"

func Test_callUriAndReturnBody_success(t *testing.T) {

	// setup
	a := NewApiClient()
	mockResponse := createDefaultResponse()
	var mockErr error
	mockHTTPClient := setupMocks(t, a)

	// verification (wrong order because of mocking)
	// check to make sure that the method calls the api with the correct request body
	mockHTTPClient.EXPECT().Post(a.uri, a.mimeType, bytes.NewBufferString(expectedAction)).Times(1).Return(mockResponse, mockErr)

	// execution
	_, _ = a.DoAction(expectedAction)

}

func Test_callUriAndReturnBody_error_reading_body(t *testing.T) {

	// setup
	a := NewApiClient()
	a.bodyReader = new(errorBodyReader)

	mockResponse := createDefaultResponse()
	var mockErr error
	mockHTTPClient := setupMocks(t, a)

	// verify (wrong order because of mocking)
	// check to make sure that the method calls the api with the correct request body
	mockHTTPClient.EXPECT().Post(a.uri, a.mimeType, bytes.NewBufferString(expectedAction)).Times(1).Return(mockResponse, mockErr)

	// execute
	// execute
	message, err := a.DoAction(expectedAction)

	// assert
	assert.Emptyf(t, message, "Expected message to be empty if ioutil.Readall returned an error, received: %s", message)
	require.NotNil(t, err, "Expected an error if ioutil.Readall returned an error reading the body")
	assert.Equalf(t, bodyReadingErrorMessage, err.Error(), "Expected a different error message, received: %s", err.Error())
}

func Test_callUriAndReturnBody_err(t *testing.T) {

	// setup
	a := NewApiClient()
	mockResponse := createDefaultResponse()
	var mockErr = errors.New(mockErrorMessage)
	mockHTTPClient := setupMocks(t, a)

	// assert (wrong order because of mocking)
	// check to make sure that the method calls the api with the correct request body
	mockHTTPClient.EXPECT().Post(a.uri, a.mimeType, bytes.NewBufferString(expectedAction)).Times(1).Return(mockResponse, mockErr)

	// execute
	message, err := a.DoAction(expectedAction)

	// assert
	assert.Empty(t, message, "Expected message to be null if http client returned an error")
	require.NotNil(t, err, "Expected to get the http error passed through")
	assert.Equalf(t, mockErrorMessage, err.Error(), "Expected a different error message, received: %s", err.Error())
}

func Test_callUriAndReturnBody_not_200(t *testing.T) {
	// setup
	a := NewApiClient()
	mockResponse := createDefaultResponse()
	mockResponse.StatusCode = 500
	var mockErr error
	mockHTTPClient := setupMocks(t, a)

	// assert (wrong order because of mocking)
	// check to make sure that the method calls the api with the correct request body
	mockHTTPClient.EXPECT().Post(a.uri, a.mimeType, bytes.NewBufferString(expectedAction)).Times(1).Return(mockResponse, mockErr)

	// execute
	message, err := a.DoAction(expectedAction)

	// assert
	assert.Empty(t, message, "Expected message to be empty if http client returned an error")
	require.NotNil(t, err, "Expected an error if the http response code was not 200")

	expectedErrorMessage := fmt.Sprintf(wrongStatusErrorFormat, mockResponse.StatusCode)
	assert.Equalf(t, expectedErrorMessage, err.Error(), "Expected a different error message, received: %s", err.Error())
}

// checks that the expected action is the body of the request that goes out.
func setupMocks(t *testing.T, client *ApiClient) (mockHTTPClient *MockHTTPClient) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mockHTTPClient = NewMockHTTPClient(mockCtrl)
	client.httpClient = mockHTTPClient
	return
}

func createDefaultResponse() (response *http.Response) {
	responseJson := `{ "result": {"Default": 1}, "error": null }`
	body := ioutil.NopCloser(bytes.NewReader([]byte(responseJson)))
	response = &http.Response{Body: body, StatusCode: 200}
	return
}

type errorBodyReader struct{}

func (*errorBodyReader) ReadAll(io.Reader) ([]byte, error) {
	return nil, errors.New(bodyReadingErrorMessage)
}
