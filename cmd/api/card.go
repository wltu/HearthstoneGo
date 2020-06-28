package api

import (
	"fmt"
	"net/http"
)

// Card provide information of a Hearthstone card
type Card struct {
	ID            int    `json:"id"`
	Collectible   int    `json:"collectible"`
	Slug          string `json:"slug"`
	ClassID       int    `json:"classId"`
	MultiClassIds []int  `json:"multiClassIds"`
	MinionTypeID  int    `json:"minionTypeId"`
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
	ChildIds      []int  `json:"childIds"`
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
func (client *HearthstoneAPI) newCardSearch(id, locale string) cardSearch {
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

// CardCollection provide information of a Hearthstone Card Collection
type CardCollection struct {
	Cards     []Card `json:"cards"`
	CardCount int    `json:"cardCount"`
	PageCount int    `json:"pageCount"`
	Page      int    `json:"page"`
}

// cardCollectionSearch provides parameters for a card collection search
type cardCollectionSearch struct {
	// Required Parameters
	url    string
	locale string

	// Optional Parameters
	optionalString map[string]string
	optionalInt    map[string]int
}

// NewCardCollectionSearch acts as a constructor for CardsSearch
func (client *HearthstoneAPI) newCardCollectionSearch(locale string) cardCollectionSearch {
	return cardCollectionSearch{
		url:    client.apiURL,
		locale: locale,
	}
}

// SetGameMode set the optional parameter of game mode for CardsSearch
func (search *cardCollectionSearch) SetGameMode(gameMode string) {
	search.optionalString["gameMode"] = gameMode
}

// SetCardSet set the optional parameter of card set for CardsSearch
func (search *cardCollectionSearch) SetCardSet(set string) {
	search.optionalString["set"] = set
}

// SetClass set the optional parameter of hero class for CardsSearch
func (search *cardCollectionSearch) SetClass(class string) {
	search.optionalString["class"] = class
}

// SetManaCost set the optional parameter of card mana cost for CardsSearch
func (search *cardCollectionSearch) SetManaCost(manaCost int) {
	search.optionalInt["manaCost"] = manaCost
}

// SetAttack set the optional parameter of minion attack for CardsSearch
func (search *cardCollectionSearch) SetAttack(attack int) {
	search.optionalInt["SetAttack"] = attack
}

// SetHealth set the optional parameter of minion health for CardsSearch
func (search *cardCollectionSearch) SetHealth(health int) {
	search.optionalInt["health"] = health
}

// SetCollectible set the optional parameter of collectible for CardsSearch
func (search *cardCollectionSearch) SetCollectible(collectible int) {
	search.optionalInt["collectible"] = collectible
}

func (search *cardCollectionSearch) execute(client *http.Client, token string) interface{} {
	url := search.url +
		"hearthstone/cards/?locale=" +
		search.locale + "&" +
		"access_token=" + token

	for key, element := range search.optionalString {
		fmt.Println("Key:", key, "=>", "Element:", element)
		url += key + "=" + element + "&"
	}

	for key, element := range search.optionalInt {
		fmt.Println("Key:", key, "=>", "Element:", element)
		url += key + "=" + string(element) + "&"
	}

	cardCollection := CardCollection{}
	err := get(client, url, &cardCollection)

	if err != nil {
		panic(err)
	}

	print(cardCollection)

	return cardCollection
}
