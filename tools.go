package ankitools

var deckNamesAndIdsAction string = "deckNamesAndIds"

func DeckNamesAndIds() (string, error) {

	deckNamesAndIdsBody, err := convert.ConvertToRequestMessage(deckNamesAndIdsAction)
	if err != nil {
		return nil, err
	}
	return apiclient.callUriAndReturnBody(deckNamesAndIdsBody)
}
