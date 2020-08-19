/*
Package convert handles JSON conversion (both ways) for types used in AnkiConnect.

It assumes AnkiConnect version 6. It does not support all types at this time.
*/
package convert

import (
	"errors"
	"github.com/ChickieDrake/ankitools/types"
)

type Action string

const (
	DecksAction      Action = "deckNames"
	QueryNotesAction Action = "findNotes"
	NotesInfoAction  Action = "notesInfo"
	AddTagAction     Action = "addTags"
	UpdateNoteAction Action = "updateNoteFields"
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

It supports the following params. If you pass in the wrong params for your action, it
does not check (YMMV with respect to the API result).
	* query
*/
func (c *Converter) ToRequestMessage(action Action, params *Params) (string, error) {
	var m string
	s := requestBody{Action: string(action), Version: 6, Params: params}
	a, err := c.marshaler.Marshal(s)
	if err != nil {
		return m, err
	}
	m = string(a)
	return m, err
}

// ToDeckNameList interprets a result from the "deckNames" AnkiConnect action
func (c *Converter) ToDeckNameList(message string) ([]string, error) {
	r := &deckNamesResponse{}
	err := c.unmarshal(message, r)
	return r.Result, err
}

// ToNoteList interprets a result from the "findNotes" AnkiConnect action
func (c *Converter) ToNoteIDList(message string) ([]int, error) {
	r := &noteIDsResponse{}
	err := c.unmarshal(message, r)
	return r.Result, err
}

func (c *Converter) ToNoteList(message string) ([]*types.NoteInfo, error) {
	r := &notesResponse{}
	err := c.unmarshal(message, r)
	return r.Result, err
}

func (c *Converter) unmarshal(message string, v apiErrorHolder) error {
	err := c.unmarshaler.Unmarshal([]byte(message), v)
	if err != nil {
		return err
	}
	if v.apiError() != "" {
		return errors.New(v.apiError())
	}
	return err
}
