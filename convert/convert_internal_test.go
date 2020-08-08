package convert

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const mockErrorMessage string = "mock: error message"

func TestConverter_ToRequestMessage_ERROR_FROM_MARSHAL_FAIL(t *testing.T) {
	// setup
	converter := New()
	converter.marshaler = new(errorMarshaler)

	// execution
	got, err := converter.ToRequestMessage("any string", nil)

	// validation
	assert.Emptyf(
		t,
		got,
		"Expected converted value to be empty if json.Marshal returned an error, received: %s",
		got,
	)
	require.NotNil(t, err, "Expected an error if the json.Marshal returned an error")
	assert.Equalf(t, mockErrorMessage, err.Error(), "Expected a different error message, received: %s", err.Error())
}

func TestConverter_ToDeckList_ERROR_FROM_UNMARSHAL_FAIL(t *testing.T) {
	// setup
	converter := New()
	converter.unmarshaler = new(errorUnmarshaler)

	// execution
	got, err := converter.ToDeckNameList("any string")

	// validation
	assert.Emptyf(
		t,
		got,
		"Expected converted value to be empty if json.Unmarshal returned an error, received: %s",
		got,
	)
	require.NotNil(t, err, "Expected an error if the json.Unmarshal returned an error")
	assert.Equalf(t, mockErrorMessage, err.Error(), "Expected a different error message, received: %s", err.Error())
}

func TestConverter_ToNoteIDList_ERROR_FROM_UNMARSHAL_FAIL(t *testing.T) {
	// setup
	converter := New()
	converter.unmarshaler = new(errorUnmarshaler)

	// execution
	got, err := converter.ToNoteIDList("any string")

	// validation
	assert.Emptyf(
		t,
		got,
		"Expected converted value to be empty if json.Unmarshal returned an error, received: %s",
		got,
	)
	require.NotNil(t, err, "Expected an error if the json.Unmarshal returned an error")
	assert.Equalf(t, mockErrorMessage, err.Error(), "Expected a different error message, received: %s", err.Error())
}

func TestConverter_ToNoteList_ERROR_FROM_UNMARSHAL_FAIL(t *testing.T) {
	// setup
	converter := New()
	converter.unmarshaler = new(errorUnmarshaler)

	// execution
	got, err := converter.ToNoteList("any string")

	// validation
	assert.Emptyf(
		t,
		got,
		"Expected converted value to be empty if json.Unmarshal returned an error, received: %s",
		got,
	)
	require.NotNil(t, err, "Expected an error if the json.Unmarshal returned an error")
	assert.Equalf(t, mockErrorMessage, err.Error(), "Expected a different error message, received: %s", err.Error())
}

type errorMarshaler struct{}

type errorUnmarshaler struct{}

func (*errorMarshaler) Marshal(interface{}) ([]byte, error) {
	err := errors.New(mockErrorMessage)
	return nil, err
}

func (*errorUnmarshaler) Unmarshal([]byte, interface{}) error {
	return errors.New(mockErrorMessage)
}
