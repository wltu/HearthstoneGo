package api

import (
	"fmt"
	"net/http"
)

type deckSearch struct {
	url    string
	id     string
	locale string
}

// Hero conatiner information about the Hearthstone Hero
type Hero Card

// String function for Hero
func (hero Hero) String() string {
	return fmt.Sprintf("%s", hero.Name)
}

// HeroPower conatiner information about the Hero's power in Hearthstone
type HeroPower Card

// String function for HeroPower
func (heroPower HeroPower) String() string {
	return fmt.Sprintf("%s: %s", heroPower.Name, heroPower.Text)
}

// Deck provides card information within a deck
type Deck struct {
	DeckCode  string    `json:"deckCode"`
	Version   int       `json:"version"`
	Format    string    `json:"format"`
	Hero      Hero      `json:"hero"`
	HeroPower HeroPower `json:"heroPower"`
	Class     Class     `json:"class"`
	Cards     []Card    `json:"cards"`
	CardCount int       `json:"cardCount"`
}

func (client *HearthstoneAPI) newDeckSearch(id string) deckSearch {
	// Required parameters
	return deckSearch{
		url:    client.apiURL,
		id:     id,
		locale: client.locale,
	}
}

func (search *deckSearch) execute(client *http.Client, token string) interface{} {
	url := search.url +
		"hearthstone/deck/" +
		search.id + "?locale=" +
		search.locale + "&" +
		"access_token=" + token

	deck := Deck{}
	err := get(client, url, &deck)

	if err != nil {
		panic(err)
	}

	// print(deck)

	return deck
}
