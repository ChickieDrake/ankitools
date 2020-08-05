package ankitools

import (
	"errors"
	"github.com/ChickieDrake/ankitools/apiclient"
	"github.com/ChickieDrake/ankitools/convert"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
	"reflect"
	"testing"
)

var expectedErr = errors.New("expected error")

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

const request = `{"action":"fake","version":6}`
const response = `{ "result": {"Default": 1}, "error": null }`

const toRequestMessage = "ToRequestMessage"
const doAction = "DoAction"
const toDeckList = "ToDeckList"

func TestTools_DeckNames_SUCCESS(t *testing.T) {
	// setup
	tools, ac, cv := toolsWithMocks()

	cv.On(toRequestMessage, convert.DecksAction).Return(request, nil).Once()
	ac.On(doAction, request).Return(response, nil).Once()
	expectedNames := []string{"Deck1", "Deck2"}
	cv.On(toDeckList, response).Return(expectedNames, nil).Once()

	// execute
	names, err := tools.DeckNames()

	// verify
	assert.Nil(t, err)
	require.NotNil(t, expectedNames)
	assert.Equal(t, expectedNames, names)

	cv.AssertExpectations(t)
	ac.AssertExpectations(t)
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
	tools, ac, cv := toolsWithMocks()

	cv.On(toRequestMessage, convert.DecksAction).Return("", expectedErr).Once()
	// execute
	names, err := tools.DeckNames()

	// verify
	assert.Nil(t, names)
	require.NotNil(t, err)
	assert.Equal(t, expectedErr, err)

	cv.AssertExpectations(t)

	ac.AssertNotCalled(t, doAction, mock.Anything)
	cv.AssertNotCalled(t, "ToDeckList", mock.Anything)
}

func TestTools_DeckNames_ERR_FROM_TO_DOACTION_FAIL(t *testing.T) {

	// setup
	tools, ac, cv := toolsWithMocks()

	cv.On(toRequestMessage, convert.DecksAction).Return(request, nil).Once()
	ac.On(doAction, request).Return("", expectedErr).Once()

	// execute
	names, err := tools.DeckNames()

	// verify
	assert.Nil(t, names)
	require.NotNil(t, err)
	assert.Equal(t, expectedErr, err)

	cv.AssertExpectations(t)
	ac.AssertExpectations(t)
	cv.AssertNotCalled(t, "ToDeckList", mock.Anything)
}

func TestTools_DeckNames_ERR_FROM_TODECKLIST_FAIL(t *testing.T) {
	// setup
	tools, ac, cv := toolsWithMocks()

	cv.On(toRequestMessage, convert.DecksAction).Return(request, nil).Once()
	ac.On(doAction, request).Return(response, nil).Once()
	cv.On(toDeckList, response).Return(nil, expectedErr).Once()

	// execute
	names, err := tools.DeckNames()

	// verify
	assert.Nil(t, names)
	require.NotNil(t, err)
	assert.Equal(t, expectedErr, err)

	cv.AssertExpectations(t)
	ac.AssertExpectations(t)
}
