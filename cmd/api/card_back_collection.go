package api

import (
	"net/http"
	"strconv"
)

type cardBackCollectionSearch struct {
	// Required Parameters
	url    string
	locale string

	// Optional Parameters
	optionalString map[string]string
	optionalInt    map[string]int
}

// CardBackCollection provide information of a set of Card Backs
type CardBackCollection struct {
	CardBacks []CardBack `json:"cardBacks"`
	CardCount int        `json:"cardCount"`
	PageCount int        `json:"pageCount"`
	Page      int        `json:"page"`
}

func (client *HearthstoneAPI) newCardBackCollectionSearch() *cardBackCollectionSearch {
	// Required parameters
	return &cardBackCollectionSearch{
		url:            client.apiURL,
		locale:         client.locale,
		optionalString: make(map[string]string),
		optionalInt:    make(map[string]int),
	}
}

func (search *cardBackCollectionSearch) execute(client *http.Client, token string) interface{} {
	url := search.url +
		"hearthstone/cardbacks/?locale=" +
		search.locale + "&"

	for key, element := range search.optionalString {
		url += key + "=" + element + "&"
	}

	for key, element := range search.optionalInt {
		url += key + "=" + strconv.Itoa(element) + "&"
	}

	url += "access_token=" + token

	cardBack := CardBackCollection{}
	err := get(client, url, &cardBack)

	if err != nil {
		panic(err)
	}

	// print(cardBack)

	return cardBack
}

// BeginCardBackCollectionSearch start the process of card collection search
func (client *HearthstoneAPI) BeginCardBackCollectionSearch() {
	client.cardBackSearch = client.newCardBackCollectionSearch()
}

// EndCardBackCollectionSearch execute the current card back collection search
func (client *HearthstoneAPI) EndCardBackCollectionSearch() *CardBackCollection {
	search := client.cardBackSearch

	if output, ok := client.execute(search).(CardBackCollection); ok {
		page := output.Page
		totalPage := output.PageCount

		for i := page + 1; i <= totalPage; i++ {
			search.setPage(i)
			if cardBacks, ok := client.execute(search).(CardBackCollection); ok {
				output.CardBacks = append(output.CardBacks, cardBacks.CardBacks...)
			} else {
				return nil
			}
		}

		search.removePage()

		return &output
	}

	return nil
}

// SetCardBackCategory filtered the card back search to a specific category
func (client *HearthstoneAPI) SetCardBackCategory(category string) {
	if client.cardBackSearch != nil {
		client.cardBackSearch.optionalString["cardBackCategory"] = category
	}
}

// SetCardBackTextFilter filtered cards search by the name of the card
func (client *HearthstoneAPI) SetCardBackTextFilter(textFilter string) {
	if client.cardBackSearch != nil {
		client.cardBackSearch.optionalString["textFilter"] = textFilter
	}
}

func (search *cardBackCollectionSearch) setPage(page int) {
	search.optionalInt["page"] = page
}

func (search *cardBackCollectionSearch) removePage() {
	delete(search.optionalInt, "page")
}
