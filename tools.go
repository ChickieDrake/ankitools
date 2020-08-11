// ankitools provides API methods that call AnkiConnect (if it is running locally).
package ankitools

import (
	"github.com/ChickieDrake/ankitools/apiclient"
	"github.com/ChickieDrake/ankitools/convert"
	"github.com/ChickieDrake/ankitools/types"
)

const ankiURI = "http://localhost:8765"

// Tools provides the methods to interact with AnkiConnect.
type Tools struct {
	ac apiClient
	cv converter
}

// New creates a new instance of Tools.
func New() *Tools {
	return &Tools{
		ac: apiclient.New(ankiURI),
		cv: convert.New(),
	}
}

// DeckNames gets a list of decks in the current collection.
func (t *Tools) DeckNames() ([]string, error) {
	m, err := t.cv.ToRequestMessage(convert.DecksAction, nil)
	if err != nil {
		return nil, err
	}

	body, err := t.ac.DoPost(m)
	if err != nil {
		return nil, err
	}

	return t.cv.ToDeckNameList(body)
}

// QueryNotes finds the info for the notes that match the query string passed in.
func (t *Tools) QueryNotes(query string) ([]*types.Note, error) {

	noteIDs, err := t.findNoteIDsByQuery(query)
	if err != nil {
		return nil, err
	}

	var notes []*types.Note
	notes, err = t.findNotesByID(noteIDs)

	return notes, err
}

//func (t *Tools) AddTag(ids []int, tag string) error {
//}

//func (t *Tools) UpdateNoteFields(note *types.Note) error {
//}

func (t *Tools) findNoteIDsByQuery(query string) ([]int, error) {
	params := &convert.Params{Query: query}

	m, err := t.cv.ToRequestMessage(convert.QueryNotesAction, params)
	if err != nil {
		return nil, err
	}

	body, err := t.ac.DoPost(m)
	if err != nil {
		return nil, err
	}
	list, err := t.cv.ToNoteIDList(body)

	// only want unique values
	var u []int
	for _, id := range list {
		found := false
		for _, uid := range u {
			if id == uid {
				found = true
				break
			}
		}
		if !found {
			u = append(u, id)
		}
	}
	return u, err
}

func (t *Tools) findNotesByID(ids []int) ([]*types.Note, error) {
	params := &convert.Params{Notes: ids}
	m, err := t.cv.ToRequestMessage(convert.NotesInfoAction, params)
	if err != nil {
		return nil, err
	}
	body, err := t.ac.DoPost(m)
	if err != nil {
		return nil, err
	}
	return t.cv.ToNoteList(body)
}

//go:generate mockery -name converter -filename mock_converter_test.go -structname MockConverter -output . -inpkg
type converter interface {
	ToRequestMessage(action convert.Action, params *convert.Params) (message string, err error)
	ToDeckNameList(message string) (decks []string, err error)
	ToNoteIDList(message string) (notes []int, err error)
	ToNoteList(message string) ([]*types.Note, error)
}

//go:generate mockery -name apiClient -filename mock_apiClient_test.go -structname MockApiClient -output . -inpkg
type apiClient interface {
	DoPost(body string) (message string, err error)
	//DoGet(url string) (*http.Response, error)
}
