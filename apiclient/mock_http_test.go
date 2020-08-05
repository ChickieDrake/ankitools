// Code generated by mockery v1.0.0. DO NOT EDIT.

package apiclient

import (
	io "io"
	http "net/http"

	mock "github.com/stretchr/testify/mock"
)

// MockHTTPClient is an autogenerated mock type for the httpClient type
type MockHTTPClient struct {
	mock.Mock
}

// Post provides a mock function with given fields: url, contentType, body
func (_m *MockHTTPClient) Post(url string, contentType string, body io.Reader) (*http.Response, error) {
	ret := _m.Called(url, contentType, body)

	var r0 *http.Response
	if rf, ok := ret.Get(0).(func(string, string, io.Reader) *http.Response); ok {
		r0 = rf(url, contentType, body)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*http.Response)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string, io.Reader) error); ok {
		r1 = rf(url, contentType, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}