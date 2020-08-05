// Code generated by mockery v1.0.0. DO NOT EDIT.

package ankitools

import (
	convert "github.com/ChickieDrake/ankitools/convert"
	mock "github.com/stretchr/testify/mock"
)

// MockConverter is an autogenerated mock type for the converter type
type MockConverter struct {
	mock.Mock
}

// ToDeckList provides a mock function with given fields: message
func (_m *MockConverter) ToDeckList(message string) ([]string, error) {
	ret := _m.Called(message)

	var r0 []string
	if rf, ok := ret.Get(0).(func(string) []string); ok {
		r0 = rf(message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(message)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// ToRequestMessage provides a mock function with given fields: action
func (_m *MockConverter) ToRequestMessage(action convert.Action) (string, error) {
	ret := _m.Called(action)

	var r0 string
	if rf, ok := ret.Get(0).(func(convert.Action) string); ok {
		r0 = rf(action)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(convert.Action) error); ok {
		r1 = rf(action)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
