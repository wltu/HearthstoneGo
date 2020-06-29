package api

import (
	"fmt"
	"net/http"
	"strconv"
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

func (client *HearthstoneAPI) newCardSearch(id string) cardSearch {
	// Required parameters
	return cardSearch{
		url:      client.apiURL,
		id:       id,
		locale:   client.locale,
		optional: make(map[string]string),
	}
}

// String function for Card
func (card Card) String() string {
	return fmt.Sprintf("%s: %d mana %d/%d",
		card.Name, card.ManaCost,
		card.Attack, card.Health)
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
		url += key + "=" + element + "&"
	}

	card := Card{}
	err := get(client, url, &card)

	if err != nil {
		panic(err)
	}

	// print(card)

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

// NewCardCollectionSearch acts as a constructor for cardCollectionSearch
func (client *HearthstoneAPI) newCardCollectionSearch() cardCollectionSearch {
	return cardCollectionSearch{
		url:            client.apiURL,
		locale:         client.locale,
		optionalString: make(map[string]string),
		optionalInt:    make(map[string]int),
	}
}

// SetGameMode set the optional parameter of game mode for cardCollectionSearch
func (search *cardCollectionSearch) SetGameMode(gameMode string) {
	search.optionalString["gameMode"] = gameMode
}

// SetCardSet set the optional parameter of card set for cardCollectionSearch
func (search *cardCollectionSearch) SetCardSet(set string) {
	search.optionalString["set"] = set
}

// SetClass set the optional parameter of hero class for cardCollectionSearch
func (search *cardCollectionSearch) SetClass(class string) {
	search.optionalString["class"] = class
}

// SetManaCost set the optional parameter of card mana cost for cardCollectionSearch
func (search *cardCollectionSearch) SetManaCost(manaCost int) {
	search.optionalInt["manaCost"] = manaCost
}

// SetTiers set the optional parameter of minion tiers (Battleground Only) for cardCollectionSearch
func (search *cardCollectionSearch) SetTiers(tiers []int) {
	output := string(tiers[0])
	for i := 1; i < len(tiers); i++ {
		output += "%2C" + string(tiers[i])
	}
	search.optionalString["tier"] = output
}

// SetAttack set the optional parameter of minion attack for cardCollectionSearch
func (search *cardCollectionSearch) SetAttack(attack int) {
	search.optionalInt["SetAttack"] = attack
}

// SetHealth set the optional parameter of minion health for cardCollectionSearch
func (search *cardCollectionSearch) SetHealth(health int) {
	search.optionalInt["health"] = health
}

// SetCollectible set the optional parameter of collectible for cardCollectionSearch
func (search *cardCollectionSearch) SetCollectible(collectible int) {
	search.optionalInt["collectible"] = collectible
}

// SetPage set the optional parameter of page number for cardCollectionSearch
// Not all the requle for the request return all at once
func (search *cardCollectionSearch) SetPage(page int) {
	search.optionalInt["page"] = page
}

func (search *cardCollectionSearch) execute(client *http.Client, token string) interface{} {
	url := search.url +
		"hearthstone/cards/?locale=" +
		search.locale + "&"

	for key, element := range search.optionalString {
		url += key + "=" + element + "&"
	}

	for key, element := range search.optionalInt {
		url += key + "=" + strconv.Itoa(element) + "&"
	}

	url += "access_token=" + token

	cardCollection := CardCollection{}
	err := get(client, url, &cardCollection)

	if err != nil {
		panic(err)
	}

	// print(cardCollection)

	return cardCollection
}
