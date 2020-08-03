package ankitools

import (
	"github.com/ChickieDrake/ankitools/apiclient"
	"github.com/ChickieDrake/ankitools/convert"
)

func DeckNames() (decks []string, err error) {
	c := convert.NewConverter()
	m, _ := c.ToRequestMessage(convert.DecksAction)
	body, _ := apiclient.NewApiClient().DoAction(m)
	return c.ToDeckList(body)
}
