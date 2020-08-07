package ankitools

import (
	"github.com/ChickieDrake/ankitools/apiclient"
	"github.com/ChickieDrake/ankitools/convert"
	"github.com/ChickieDrake/ankitools/types"
)

type Tools struct {
	ac apiClient
	cv converter
}

func New() *Tools {
	return &Tools{
		ac: apiclient.New(),
		cv: convert.New(),
	}
}

func (t *Tools) DeckNames() ([]string, error) {
	m, err := t.cv.ToRequestMessage(convert.DecksAction, nil)
	if err != nil {
		return nil, err
	}

	body, err := t.ac.DoAction(m)
	if err != nil {
		return nil, err
	}

	return t.cv.ToDeckList(body)
}

func (t *Tools) QueryNotes(query string) ([]*types.Note, error) {
	params := &convert.Params{Query: query}

	m, err := t.cv.ToRequestMessage(convert.QueryNotesAction, params)
	if err != nil {
		return nil, err
	}

	body, err := t.ac.DoAction(m)
	if err != nil {
		return nil, err
	}

	return t.cv.ToNoteList(body)
}

//go:generate mockery -name converter -filename mock_converter_test.go -structname MockConverter -output . -inpkg
type converter interface {
	ToRequestMessage(action convert.Action, params *convert.Params) (message string, err error)
	ToDeckList(message string) (decks []string, err error)
	ToNoteList(message string) (notes []*types.Note, err error)
}

//go:generate mockery -name apiClient -filename mock_apiClient_test.go -structname MockApiClient -output . -inpkg
type apiClient interface {
	DoAction(body string) (message string, err error)
}
