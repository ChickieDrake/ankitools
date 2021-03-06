// Code generated by mockery v1.0.0. DO NOT EDIT.

package ankitools

import mock "github.com/stretchr/testify/mock"

// MockApiClient is an autogenerated mock type for the apiClient type
type MockApiClient struct {
	mock.Mock
}

// DoPost provides a mock function with given fields: uri, body
func (_m *MockApiClient) DoPost(uri string, body string) (string, error) {
	ret := _m.Called(uri, body)

	var r0 string
	if rf, ok := ret.Get(0).(func(string, string) string); ok {
		r0 = rf(uri, body)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(uri, body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
