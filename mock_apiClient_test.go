// Code generated by mockery v1.0.0. DO NOT EDIT.

package ankitools

import mock "github.com/stretchr/testify/mock"

// MockApiClient is an autogenerated mock type for the apiClient type
type MockApiClient struct {
	mock.Mock
}

// DoAction provides a mock function with given fields: body
func (_m *MockApiClient) DoPost(body string) (string, error) {
	ret := _m.Called(body)

	var r0 string
	if rf, ok := ret.Get(0).(func(string) string); ok {
		r0 = rf(body)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(body)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
