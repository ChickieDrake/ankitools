package ankitools

import (
	"github.com/ChickieDrake/ankitools/apiclient"
	"github.com/ChickieDrake/ankitools/convert"
)

type Tools struct {
	apiClient := apiclient.New()

}

func DeckNames() (decks []string, err error) {
	c := convert.New()
	m, _ := c.ToRequestMessage(convert.DecksAction)
	body, _ := apiclient.New().DoAction(m)
	return c.ToDeckList(body)
}


