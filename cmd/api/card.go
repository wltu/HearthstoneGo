package api

import (
	"fmt"
	"net/http"
)

// CardError is present if the card is not found
type CardError struct {
	Status        int    `json:"status"`
	StatusMessage string `json:"statusMessage"`
	Message       string `json:"message"`
}

// Battlegrounds provide information of a card in battlegrounds game mode
type Battlegrounds struct {
	Tier      int    `json:"tier"`
	Hero      bool   `json:"hero"`
	UpgradeID int    `json:"upgradeId"`
	Image     string `json:"image"`
	ImageGold string `json:"imageGold"`
}

// Card provide information of a Hearthstone card
type Card struct {
	ID            int           `json:"id"`
	Collectible   int           `json:"collectible"`
	Slug          string        `json:"slug"`
	ClassID       int           `json:"classId"`
	MultiClassIds []int         `json:"multiClassIds"`
	MinionTypeID  int           `json:"minionTypeId"`
	CardTypeID    int           `json:"cardTypeId"`
	CardSetID     int           `json:"cardSetId"`
	RarityID      int           `json:"rarityId"`
	ArtistName    string        `json:"artistName"`
	Health        int           `json:"health"`
	Attack        int           `json:"attack"`
	ManaCost      int           `json:"manaCost"`
	Name          string        `json:"name"`
	Text          string        `json:"text"`
	Image         string        `json:"image"`
	ImageGold     string        `json:"imageGold"`
	FlavorText    string        `json:"flavorText"`
	CropImage     string        `json:"cropImage"`
	ChildIds      []int         `json:"childIds"`
	KeywordIds    []int         `json:"keywordIds"`
	Battlegrounds Battlegrounds `json:"battlegrounds"`
	Error         CardError     `json:"error"`
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
	if card.Battlegrounds != (Battlegrounds{}) {
		return fmt.Sprintf("%s: tier %d %d/%d",
			card.Name, card.Battlegrounds.Tier,
			card.Attack, card.Health)
	}

	return fmt.Sprintf("%s: %d mana %d/%d",
		card.Name, card.ManaCost,
		card.Attack, card.Health)
}

func (search *cardSearch) setID(id string) {
	search.id = id
}

func (search *cardSearch) setLocale(locale string) {
	search.locale = locale
}

func (search *cardSearch) setGameMode(gameMode string) {
	search.optional["gameMode"] = gameMode
}

func (search *cardSearch) execute(client *http.Client, token string) interface{} {
	url := search.url +
		"hearthstone/cards/" +
		search.id + "?locale=" +
		search.locale + "&"

	for key, element := range search.optional {
		url += key + "=" + element + "&"
	}

	url += "access_token=" + token

	card := Card{}
	err := get(client, url, &card)

	if err != nil {
		panic(err)
	}

	// print(card)

	return card
}

// SearchCard request a specific card by id
// Output will be the card and a vailidity check.
// It will return true if the card is found and is a constructed card, otherwise false.
func (client *HearthstoneAPI) SearchCard(id string) (Card, bool) {
	search := client.newCardSearch(id)

	if output, ok := client.execute(&search).(Card); ok {
		if output.Error.Status != 0 {
			return Card{}, false
		}

		if output.Collectible == 1 {
			return output, true
		}
	}

	return Card{}, false
}

// SearchBattlegroundsCard request a specific battlegrounds card by id
// Output will be the card and a vailidity check.
// It will return true if the card is found and is a battlegrounds card, otherwise false.
func (client *HearthstoneAPI) SearchBattlegroundsCard(id string) (Card, bool) {
	search := client.newCardSearch(id)
	search.setGameMode("battlegrounds")

	if output, ok := client.execute(&search).(Card); ok {
		if output.Error.Status != 0 {
			return Card{}, false
		}

		if output.Battlegrounds != (Battlegrounds{}) {
			return output, true
		}
	}

	return Card{}, false
}
