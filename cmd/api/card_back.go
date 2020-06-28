package api

import (
	"fmt"
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

	print(cardBack)

	return cardBack
}

type cardBackCollectionSearch struct {
	// Required Parameters
	url    string
	locale string

	// Optional Parameters
	optional map[string]string
}

// CardBackCollection provide information of a Hearthstone Card Back Collection
type CardBackCollection struct {
	CardBacks []CardBack `json:"cardBacks"`
	CardCount int        `json:"cardCount"`
	PageCount int        `json:"pageCount"`
	Page      int        `json:"page"`
}

func (client *HearthstoneAPI) newCardBackCollectionSearch() cardBackCollectionSearch {
	// Required parameters
	return cardBackCollectionSearch{
		url:    client.apiURL,
		locale: client.locale,
	}
}

// SetCategory set the optional parameter of Category for CardBackCollectionSearch
func (search *cardBackCollectionSearch) SetCategory(category string) {
	search.optional["cardBackCategory"] = category
}

// SetTextFilter set the optional parameter of text filter for CardBackCollectionSearch
func (search *cardBackCollectionSearch) SetTextFilter(textFilter string) {
	search.optional["textFilter"] = textFilter
}

// SetCategory set the optional parameter of the field to use to sort for CardBackCollectionSearch
func (search *cardBackCollectionSearch) SetSort(sort string) {
	search.optional["sort"] = sort
}

// SetOrder set the optional parameter of how to use the field to sort CardBackCollectionSearch
func (search *cardBackCollectionSearch) SetOrder(order string) {
	search.optional["order"] = order
}

func (search *cardBackCollectionSearch) execute(client *http.Client, token string) interface{} {
	url := search.url +
		"hearthstone/cardbacks/?locale=" +
		search.locale + "&" +
		"access_token=" + token

	for key, element := range search.optional {
		fmt.Println("Key:", key, "=>", "Element:", element)
		url += key + "=" + string(element) + "&"
	}

	cardBack := CardBackCollection{}
	err := get(client, url, &cardBack)

	if err != nil {
		panic(err)
	}

	print(cardBack)

	return cardBack
}
