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

// CardImage download card image and return saved location
func (client *HearthstoneAPI) CardImage(card *Card) string {
	return getImage(client.heartstoneClient, card.Slug, card.Image)
}

// // GoldCardImage download golden card image and return saved location
// func (card *Card) GoldCardImage(id string) string {

// }

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
func (client *HearthstoneAPI) SearchCard(id string) *Card {
	search := client.newCardSearch(id)

	if output, ok := client.execute(&search).(Card); ok {
		if output.Error.Status != 0 {
			return nil
		}

		if output.Collectible == 1 {
			return &output
		}
	}

	return nil
}

// SearchBattlegroundsCard request a specific battlegrounds card by id
func (client *HearthstoneAPI) SearchBattlegroundsCard(id string) *Card {
	search := client.newCardSearch(id)
	search.optional["gameMode"] = "battlegrounds"

	if output, ok := client.execute(&search).(Card); ok {
		if output.Error.Status != 0 {
			return nil
		}

		if output.Battlegrounds != (Battlegrounds{}) {
			return &output
		}
	}

	return nil
}
