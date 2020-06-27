package api

import (
	"fmt"
	"net/http"
)

// Card provide information of a Hearstone card
type Card struct {
	ID            int    `json:"id"`
	Collectible   int    `json:"collectible"`
	Slug          string `json:"slug"`
	ClassID       int    `json:"classId"`
	MultiClassIds []int  `json:"multiClassIds"`
	CardTypeID    int    `json:"cardTypeId"`
	CardSetID     int    `json:"cardSetId"`
	RarityID      int    `json:"rarityId"`
	ArtistName    string `json:"artistName"`
	Health        int    `json:"health"`
	Attack        int    `json:"attack"`
	ManaCost      int    `json:"manaCost"`
	Name          string `json:"name"`
	Text          string `json:"text"`
	Image         string `json:"image"`
	ImageGold     string `json:"imageGold"`
	FlavorText    string `json:"flavorText"`
	CropImage     string `json:"cropImage"`
	KeywordIds    []int  `json:"keywordIds"`
}

type cardSearch struct {
	// Required Parameters
	url    string
	id     string
	locale string
	// Optional Parameters
	optional map[string]string
}

// NewCardSearch acts as a constructor for cardSearch
func (client *HeartstoneAPI) newCardSearch(id, locale string) cardSearch {
	// Required parameters
	return cardSearch{
		url:      client.apiURL,
		id:       id,
		locale:   locale,
		optional: make(map[string]string),
	}
}

// SetID update the current id value for cardSearch
func (search *cardSearch) SetID(id string) {
	search.id = id
}

// SetLocale set the optional parameter of locale for cardSearch
func (search *cardSearch) SetLocale(locale string) {
	search.locale = locale
}

// SetGameMode set the optional parameter of game mode for cardSearch
func (search *cardSearch) SetGameMode(gameMode string) {
	search.optional["gameMode"] = gameMode
}

func (search *cardSearch) test() {

}

func (search *cardSearch) execute(client *http.Client, token string) interface{} {
	url := search.url +
		"hearthstone/cards/" +
		search.id + "?locale=" +
		search.locale + "&" +
		"access_token=" + token

	for key, element := range search.optional {
		fmt.Println("Key:", key, "=>", "Element:", element)
		url += key + "=" + element + "&"
	}

	card := Card{}
	err := get(client, url, &card)

	if err != nil {
		panic(err)
	}

	print(card)

	return card
}

// CardCollectionSearch provides parameters for a card collection search
type CardCollectionSearch struct {
	// Required Parameters
	locale string

	// Optional Parameters
	optionalString map[string]string
	optionalInt    map[string]int
}

// NewCardCollectionSearch acts as a constructor for CardsSearch
func (client *HeartstoneAPI) NewCardCollectionSearch() CardCollectionSearch {
	return CardCollectionSearch{}
}

// // SetLocale set the optional parameter of locale for CardsSearch
// func (cardsSearch *CardsSearch) SetLocale(locale string) {
// 	cardSearch.optionalString["locale"] = locale
// }

// // SetGameMode set the optional parameter of game mode for CardsSearch
// func (cardsSearch *CardsSearch) SetGameMode(gameMode string) {
// 	cardSearch.optionalString["gameMode"] = gameMode
// }

// // SetCardSet set the optional parameter of card set for CardsSearch
// func (cardsSearch *CardsSearch) SetCardSet(set string) {
// 	cardSearch.optionalString["set"] = set
// }

// // SetClass set the optional parameter of hero class for CardsSearch
// func (cardsSearch *CardsSearch) SetClass(class string) {
// 	cardSearch.optionalString["class"] = class
// }

// // SetManaCost set the optional parameter of card mana cost for CardsSearch
// func (cardsSearch *CardsSearch) SetManaCost(manaCost int) {
// 	cardSearch.optionalInt["manaCost"] = manaCost
// }

// // SetAttack set the optional parameter of minion attack for CardsSearch
// func (cardsSearch *CardsSearch) SetAttack(attack int) {
// 	cardSearch.optionalInt["SetAttack"] = SetAttack
// }

// // SetHealth set the optional parameter of minion health for CardsSearch
// func (cardsSearch *CardsSearch) SetHealth(health int) {
// 	cardSearch.optionalInt["health"] = health
// }

// // SetCollectible set the optional parameter of collectible for CardsSearch
// func (cardsSearch *CardsSearch) SetCollectible(collectible bool) {
// 	cardSearch.optionalInt["collectible"] = collectible
// }
