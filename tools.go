package ankitools

import (
	"github.com/ChickieDrake/ankitools/apiclient"
	"github.com/ChickieDrake/ankitools/convert"
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
	m, err := t.cv.ToRequestMessage(convert.DecksAction)
	if err != nil {
		return nil, err
	}

	body, err := t.ac.DoAction(m)
	if err != nil {
		return nil, err
	}

	return t.cv.ToDeckList(body)
}

//go:generate mockery -name converter
type converter interface {
	ToRequestMessage(action convert.Action) (message string, err error)
	ToDeckList(message string) (decks []string, err error)
}

//go:generate mockery -name apiClient
type apiClient interface {
	DoAction(body string) (message string, err error)
}
