package ankitools

import "encoding/json"

var deckNamesAndIdsAction string = "deckNamesAndIds"

func DeckNamesAndIds() (string, error) {

	deckNamesAndIdsBody, err := createActionString(deckNamesAndIdsAction)
	if err != nil {
		return nil, err
	}
	return apiclient.callUriAndReturnBody(deckNamesAndIdsBody)
}

type request_body struct {
	Action  string `json:"action"`
	Version int    `json:"version"`
}

func createActionString(action string) (string, error) {
	s := &request_body{Action: action, Version: 6}
	a, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	return string(a, err)
}
