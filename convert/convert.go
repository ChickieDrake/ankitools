/*
Package convert handles JSON conversion (both ways) for types used in AnkiConnect.

It assumes AnkiConnect version 6. It does not support all types at this time.
*/
package convert

import (
	"encoding/json"
	"errors"
)

type Action string

const (
	DecksAction Action = "deckNames"
)

// Converter provides methods that can be used to create and interpret messages for AnkiConnect
type Converter struct {
	// Internal parameters that need to be mocked in tests
	marshaler   marshaler
	unmarshaler unmarshaler
}

// New creates a new converter instance.
func New() *Converter {
	return &Converter{
		marshaler:   new(defaultMarshaler),
		unmarshaler: new(defaultUnmarshaler),
	}
}

/*
ToRequestMessage takes the name of an action and returns a message for AnkiConnect.

It does not support parametrization at this time.
*/
func (c *Converter) ToRequestMessage(action Action) (message string, err error) {
	s := requestBody{Action: string(action), Version: 6}
	a, err := c.marshaler.Marshal(s)
	if err != nil {
		return
	}
	return string(a), err
}

// ToDeckList interprets a response from the "deckNames" AnkiConnect action
func (c *Converter) ToDeckList(message string) (decks []string, err error) {
	r := &deckNamesResponse{}
	err = c.unmarshaler.Unmarshal([]byte(message), r)
	if r.Error != nil {
		return nil, errors.New(*r.Error)
	}
	return r.Result, err
}

type requestBody struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
}

type deckNamesResponse struct {
	Result []string
	Error  *string
}

type marshaler interface {
	Marshal(v interface{}) ([]byte, error)
}

type unmarshaler interface {
	Unmarshal(data []byte, v interface{}) error
}

type defaultMarshaler struct{}

type defaultUnmarshaler struct{}

func (*defaultMarshaler) Marshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func (*defaultUnmarshaler) Unmarshal(data []byte, v interface{}) error {
	return json.Unmarshal(data, v)
}
