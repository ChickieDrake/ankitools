package convert

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

const mockErrorMessage string = "mock: error message"

func TestToRequestMessage_Error_From_Marshal(t *testing.T) {
	// setup
	converter := NewConverter()
	converter.marshaler = new(errorMarshaler)

	// execution
	got, err := converter.ToRequestMessage("anystring")

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

func TestMessageToDeckList_Error_From_Unmarshal(t *testing.T) {
	// setup
	converter := NewConverter()
	converter.unmarshaler = new(errorUnmarshaler)

	// execution
	got, err := converter.ToDeckList("anystring")

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

func (*errorMarshaler) Marshal(v interface{}) ([]byte, error) {
	err := errors.New(mockErrorMessage)
	return nil, err
}

func (*errorUnmarshaler) Unmarshal(data []byte, v interface{}) error {
	return errors.New(mockErrorMessage)
}