package api

import (
	"net/http"
)

type cardBackSearch struct {
	url    string
	id     string
	locale string
}

// CardBack provide information of a Hearthstone card backs
type CardBack struct {
	ID           int    `json:"id"`
	SortCategory int    `json:"sortCategory"`
	Text         string `json:"text"`
	Name         string `json:"name"`
	Image        string `json:"image"`
	Slug         string `json:"slug"`
}

func (client *HearthstoneAPI) newCardBackSearch(id string) cardBackSearch {
	// Required parameters
	return cardBackSearch{
		url:    client.apiURL,
		id:     id,
		locale: client.locale,
	}
}

func (search *cardBackSearch) execute(client *http.Client, token string) interface{} {
	url := search.url +
		"hearthstone/cardbacks/" +
		search.id + "?locale=" +
		search.locale + "&" +
		"access_token=" + token

	cardBack := CardBack{}
	err := get(client, url, &cardBack)

	if err != nil {
		panic(err)
	}

	// print(cardBack)

	return cardBack
}

// SearchCardBack make a API call to search for a card back with the given id
func (client *HearthstoneAPI) SearchCardBack(id string) CardBack {
	search := client.newCardBackSearch(id)

	if output, ok := client.execute(&search).(CardBack); ok {
		return output
	}

	return CardBack{}
}
