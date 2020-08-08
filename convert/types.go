package convert

import (
	"encoding/json"
	"github.com/ChickieDrake/ankitools/types"
)

/*
Params represents the basic structure used by AnkiConnect for request parameters.

Not all params are used by requests, but this represents the param types that are supported.
*/
type Params struct {
	Query string `json:"query,omitempty"`
	Notes []int  `json:"notes,omitempty"`
}

type requestBody struct {
	Action  string  `json:"action"`
	Version int     `json:"version"`
	Params  *Params `json:"params,omitempty"`
}

type errorResponse struct {
	Error string
}

func (e *errorResponse) apiError() string {
	return e.Error
}

type apiErrorHolder interface {
	apiError() string
}

type deckNamesResponse struct {
	Result []string
	errorResponse
}

type noteIDsResponse struct {
	Result []int
	errorResponse
}

type notesResponse struct {
	Result []*types.Note
	errorResponse
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
