package ankitools

import (
	"errors"
	"fmt"
	"github.com/ChickieDrake/ankitools/apiclient"
	"github.com/ChickieDrake/ankitools/convert"
	"github.com/ChickieDrake/ankitools/types"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

var expectedErr = errors.New("expected error")
var emptyParams *convert.Params

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *Tools
	}{
		{"Happy path", &Tools{apiclient.New(), convert.New()}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

const requestTemplate = `{"action":"%s","version":6}`
const requestTemplateWithParams = `{"action":"%s","version":6,"params":%s}`

// Valid HTTP requests (for mocking)
var deckNamesRequest = fmt.Sprintf(requestTemplate, convert.DecksAction)
var queryNotesRequest = fmt.Sprintf(requestTemplateWithParams, convert.QueryNotesAction, `{"query":"deck:current"}`)
var notesInfoRequest = fmt.Sprintf(requestTemplateWithParams, convert.NotesInfoAction, `{"notes":[1483959289817,1483959291695]}`)

// Valid HTTP responses (for mocking)
const deckNamesResponse = `{"result": ["Deck1","Deck2"],"error": null}`
const queryNotesResponse = `{"result": [1483959289817, 1483959291695],"error": null}`
const notesInfoResponse = `{
		"result": [
			{
				"noteId":1483959289817
			},
			{
				"noteId":1483959291695
			}
		],
		"error": null
	}`

// Valid mock object method names
const toRequestMessage = "ToRequestMessage"
const doAction = "DoPost"
const toDeckList = "ToDeckNameList"
const toNoteIDList = "ToNoteIDList"
const toNoteList = "ToNoteList"

// Notes query string (and struct)
const notesQuery = "deck:current"

var notesQueryParams = &convert.Params{Query: notesQuery}
var notesInfoParams = &convert.Params{Notes: expectedNoteIDs}

var expectedNoteIDs = []int{1483959289817, 1483959291695}
var expectedNotes = []*types.Note{
	{NoteID: 1483959289817},
	{NoteID: 1483959291695},
}

func TestTools_DeckNames_SUCCESS(t *testing.T) {
	// setup
	tools := New()
	ac := &MockApiClient{}
	tools.ac = ac

	//cv.On(toRequestMessage, convert.DecksAction, emptyParams).Return(deckNamesRequest, nil).Once()
	ac.On(doAction, ankiURI, deckNamesRequest).Return(deckNamesResponse, nil).Once()

	// execute
	names, err := tools.DeckNames()

	// verify
	expectedNames := []string{"Deck1", "Deck2"}
	assert.Nil(t, err)
	require.NotNil(t, names)
	assert.Equal(t, expectedNames, names)
}

func toolsWithMocks() (*Tools, *MockApiClient, *MockConverter) {
	tools := New()
	ac := &MockApiClient{}
	tools.ac = ac
	cv := &MockConverter{}
	tools.cv = cv
	return tools, ac, cv
}

func TestTools_DeckNames_ERR_FROM_TO_REQUEST_MESSAGE_FAIL(t *testing.T) {
	// setup
	tools, _, cv := toolsWithMocks()

	cv.On(toRequestMessage, convert.DecksAction, emptyParams).Return("", expectedErr).Once()
	// execute
	names, err := tools.DeckNames()

	// verify
	assert.Nil(t, names)
	require.NotNil(t, err)
	assert.Equal(t, expectedErr, err)
}

func TestTools_DeckNames_ERR_FROM_TO_DOACTION_FAIL(t *testing.T) {

	// setup
	tools, ac, cv := toolsWithMocks()

	cv.On(toRequestMessage, convert.DecksAction, emptyParams).Return(deckNamesRequest, nil).Once()
	ac.On(doAction, ankiURI, deckNamesRequest).Return("", expectedErr).Once()

	// execute
	names, err := tools.DeckNames()

	// verify
	assert.Nil(t, names)
	require.NotNil(t, err)
	assert.Equal(t, expectedErr, err)
}

func TestTools_DeckNames_ERR_FROM_TODECKNAMELIST_FAIL(t *testing.T) {
	// setup
	tools, ac, cv := toolsWithMocks()

	cv.On(toRequestMessage, convert.DecksAction, emptyParams).Return(deckNamesRequest, nil).Once()
	ac.On(doAction, ankiURI, deckNamesRequest).Return(deckNamesResponse, nil).Once()
	cv.On(toDeckList, deckNamesResponse).Return(nil, expectedErr).Once()

	// execute
	names, err := tools.DeckNames()

	// verify
	assert.Nil(t, names)
	require.NotNil(t, err)
	assert.Equal(t, expectedErr, err)
}

func TestTools_QueryNotes_SUCCESS(t *testing.T) {
	// setup
	tools := New()
	ac := &MockApiClient{}
	tools.ac = ac

	ac.On(doAction, ankiURI, queryNotesRequest).Return(queryNotesResponse, nil).
		On(doAction, ankiURI, notesInfoRequest).Return(notesInfoResponse, nil)

	// execute
	notes, err := tools.QueryNotes(notesQuery)

	// verify
	assert.Nil(t, err)
	require.NotNil(t, notes)
	assert.Equal(t, expectedNotes, notes)

}

func TestTools_QueryNotes_DUPLICATE_VALUES_SUCCESS(t *testing.T) {
	// setup
	tools := New()
	ac := &MockApiClient{}
	tools.ac = ac

	// test the condition when the query returns duplicate not IDs (only get unique)
	const dupQueryNotesResponse = `{"result": [1483959289817, 1483959289817],"error": null}`
	var dupNotesInfoRequest = fmt.Sprintf(requestTemplateWithParams, convert.NotesInfoAction, `{"notes":[1483959289817]}`)
	const dupNotesInfoResponse = `{
		"result": [
			{
				"noteId":1483959289817
			}
		],
		"error": null
	}`

	ac.On(doAction, ankiURI, queryNotesRequest).Return(dupQueryNotesResponse, nil).
		On(doAction, ankiURI, dupNotesInfoRequest).Return(dupNotesInfoResponse, nil)

	// execute
	notes, err := tools.QueryNotes(notesQuery)

	// verify
	assert.Nil(t, err)
	require.NotNil(t, notes)

	var expectedNotes = []*types.Note{
		{NoteID: 1483959289817},
	}
	assert.Equal(t, expectedNotes, notes)

	// make sure that the correct requests were sent to the api client
	ac.AssertExpectations(t)

}

func TestTools_QueryNotes_ERR_FROM_TO_REQUEST_MESSAGE_FAIL(t *testing.T) {
	// setup
	tools, _, cv := toolsWithMocks()

	cv.On(
		toRequestMessage,
		convert.QueryNotesAction,
		notesQueryParams,
	).Return("", expectedErr).Once()

	// execute
	notes, err := tools.QueryNotes(notesQuery)

	// verify
	assert.Nil(t, notes)
	require.NotNil(t, err)
	assert.Equal(t, expectedErr, err)

}

func TestTools_QueryNotes_ERR_FROM_DOACTION_FAIL(t *testing.T) {

	// setup
	tools := New()
	ac := &MockApiClient{}
	tools.ac = ac

	ac.On(doAction, ankiURI, queryNotesRequest).Return("", expectedErr).Once()

	// execute
	notes, err := tools.QueryNotes(notesQuery)

	// verify
	assert.Nil(t, notes)
	require.NotNil(t, err)
	assert.Equal(t, expectedErr, err)
}

func TestTools_QueryNotes_ERR_FROM_TONOTEIDLIST_FAIL(t *testing.T) {
	// setup
	tools, ac, cv := toolsWithMocks()

	cv.On(toRequestMessage, convert.QueryNotesAction, notesQueryParams).Return(queryNotesRequest, nil).Once()
	ac.On(doAction, ankiURI, queryNotesRequest).Return(queryNotesResponse, nil).Once()
	cv.On(toNoteIDList, queryNotesResponse).Return(nil, expectedErr).Once()

	// execute
	notes, err := tools.QueryNotes(notesQuery)

	// verify
	assert.Nil(t, notes)
	require.NotNil(t, err)
	assert.Equal(t, expectedErr, err)
}

func TestTools_QueryNotes_ERR_FROM_SECOND_TOREQUESTMESSAGE_FAIL(t *testing.T) {
	tools, ac, cv := toolsWithMocks()

	// All this is passing behaviour
	cv.On(toRequestMessage, convert.QueryNotesAction, notesQueryParams).Return(queryNotesRequest, nil).Once()
	ac.On(doAction, ankiURI, queryNotesRequest).Return(queryNotesResponse, nil).Once()
	cv.On(toNoteIDList, queryNotesResponse).Return(expectedNoteIDs, nil).Once()

	// This is throwing the error we expect to see
	cv.On(toRequestMessage, convert.NotesInfoAction, notesInfoParams).Return("", expectedErr).Once()

	// execute
	notes, err := tools.QueryNotes(notesQuery)

	// verify
	assert.Nil(t, notes)
	require.NotNil(t, err)
	assert.Equal(t, expectedErr, err)
}

func TestTools_QueryNotes_ERR_FROM_SECOND_DOACTION_FAIL(t *testing.T) {
	tools, ac, cv := toolsWithMocks()

	// All this is passing behaviour
	cv.On(toRequestMessage, convert.QueryNotesAction, notesQueryParams).Return(queryNotesRequest, nil).Once()
	ac.On(doAction, ankiURI, queryNotesRequest).Return(queryNotesResponse, nil).Once()
	cv.On(toNoteIDList, queryNotesResponse).Return(expectedNoteIDs, nil).Once()
	cv.On(toRequestMessage, convert.NotesInfoAction, notesInfoParams).Return(notesInfoRequest, nil).Once()

	// This is throwing the error we expect to see
	ac.On(doAction, ankiURI, notesInfoRequest).Return("", expectedErr).Once()

	// execute
	notes, err := tools.QueryNotes(notesQuery)

	// verify
	assert.Nil(t, notes)
	require.NotNil(t, err)
	assert.Equal(t, expectedErr, err)
}

func TestTools_QueryNotes_ERR_FROM_TONOTELIST_FAIL(t *testing.T) {
	// setup
	tools, ac, cv := toolsWithMocks()

	// All this is passing behaviour
	cv.On(toRequestMessage, convert.QueryNotesAction, notesQueryParams).Return(queryNotesRequest, nil).Once()
	ac.On(doAction, ankiURI, queryNotesRequest).Return(queryNotesResponse, nil).Once()
	cv.On(toNoteIDList, queryNotesResponse).Return(expectedNoteIDs, nil).Once()

	cv.On(toRequestMessage, convert.NotesInfoAction, notesInfoParams).Return(notesInfoRequest, nil).Once()
	ac.On(doAction, ankiURI, notesInfoRequest).Return(notesInfoResponse, nil).Once()

	// This is throwing the error we expect to see
	cv.On(toNoteList, notesInfoResponse).Return(nil, expectedErr).Once()

	// execute
	notes, err := tools.QueryNotes(notesQuery)

	// verify
	assert.Nil(t, notes)
	require.NotNil(t, err)
	assert.Equal(t, expectedErr, err)
}
