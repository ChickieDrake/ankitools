// Code generated by mockery v1.0.0. DO NOT EDIT.

package ankitools

import (
	convert "github.com/ChickieDrake/ankitools/convert"
	mock "github.com/stretchr/testify/mock"

	types "github.com/ChickieDrake/ankitools/types"
)

// MockConverter is an autogenerated mock type for the converter type
type MockConverter struct {
	mock.Mock
}

// ToDeckNameList provides a mock function with given fields: message
func (_m *MockConverter) ToDeckNameList(message string) ([]string, error) {
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

// ToNoteIDList provides a mock function with given fields: message
func (_m *MockConverter) ToNoteIDList(message string) ([]int, error) {
	ret := _m.Called(message)

	var r0 []int
	if rf, ok := ret.Get(0).(func(string) []int); ok {
		r0 = rf(message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]int)
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

// ToNoteList provides a mock function with given fields: message
func (_m *MockConverter) ToNoteList(message string) ([]*types.NoteInfo, error) {
	ret := _m.Called(message)

	var r0 []*types.NoteInfo
	if rf, ok := ret.Get(0).(func(string) []*types.NoteInfo); ok {
		r0 = rf(message)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*types.NoteInfo)
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

// ToRequestMessage provides a mock function with given fields: action, params
func (_m *MockConverter) ToRequestMessage(action convert.Action, params *convert.Params) (string, error) {
	ret := _m.Called(action, params)

	var r0 string
	if rf, ok := ret.Get(0).(func(convert.Action, *convert.Params) string); ok {
		r0 = rf(action, params)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(convert.Action, *convert.Params) error); ok {
		r1 = rf(action, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
