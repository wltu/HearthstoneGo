package api

import (
	"net/http"
	"strconv"
)

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
func (client *HearthstoneAPI) newCardCollectionSearch() *cardCollectionSearch {
	return &cardCollectionSearch{
		url:            client.apiURL,
		locale:         client.locale,
		optionalString: make(map[string]string),
		optionalInt:    make(map[string]int),
	}
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

func (search *cardCollectionSearch) setPage(page int) {
	search.optionalInt["page"] = page
}

func (search *cardCollectionSearch) removePage() {
	delete(search.optionalInt, "page")
}

// BeginCardCollectionSearch start the process of card collection search
func (client *HearthstoneAPI) BeginCardCollectionSearch() {
	client.cardSearch = client.newCardCollectionSearch()
}

// EndCardCollectionSearch execute the current card collection search
func (client *HearthstoneAPI) EndCardCollectionSearch() *CardCollection {
	search := client.cardSearch

	if output, ok := client.execute(search).(CardCollection); ok {
		page := output.Page
		totalPage := output.PageCount

		for i := page + 1; i <= totalPage; i++ {
			search.setPage(i)
			if cards, ok := client.execute(search).(CardCollection); ok {
				output.Cards = append(output.Cards, cards.Cards...)
			} else {
				return nil
			}
		}

		search.removePage()

		return &output
	}

	return nil
}

// SetCardGameMode select the game mode to search the cards in
func (client *HearthstoneAPI) SetCardGameMode(gameMode string) {
	if client.cardSearch != nil {
		client.cardSearch.optionalString["gameMode"] = gameMode
	}
}

// SetCardSet select the card set to search the cards in
func (client *HearthstoneAPI) SetCardSet(set string) {
	if client.cardSearch != nil {
		client.cardSearch.optionalString["set"] = set
	}
}

// SetCardClass select the hero class for the cards
func (client *HearthstoneAPI) SetCardClass(class string) {
	if client.cardSearch != nil {
		client.cardSearch.optionalString["class"] = class
	}
}

// SetCardManaCost filtered cards search to only ones with same mana cost
func (client *HearthstoneAPI) SetCardManaCost(manaCost int) {
	if client.cardSearch != nil {
		client.cardSearch.optionalInt["manaCost"] = manaCost
	}
}

// SetCardTiers filtered the cards search to the selected tavern tier (BattleGrounds)
func (client *HearthstoneAPI) SetCardTiers(tiers []int) {
	if client.cardSearch != nil {
		output := strconv.Itoa(tiers[0])
		for i := 1; i < len(tiers); i++ {
			output += "%2C" + strconv.Itoa(tiers[i])
		}
		client.cardSearch.optionalString["tier"] = output
	}
}

// SetCardHero allow the card search to get all the hero cards (BattleGrounds)
func (client *HearthstoneAPI) SetCardHero() {
	if client.cardSearch != nil {
		client.cardSearch.optionalString["tier"] = "hero"
	}
}

// SetCardAttack filtered cards search to only ones with attack
func (client *HearthstoneAPI) SetCardAttack(attack int) {
	if client.cardSearch != nil {
		client.cardSearch.optionalInt["attack"] = attack
	}
}

// SetCardHealth filtered cards search to only ones with health
func (client *HearthstoneAPI) SetCardHealth(health int) {
	if client.cardSearch != nil {
		client.cardSearch.optionalInt["health"] = health
	}
}

// SetCardCollectible filtered cards search to collectible or not collectible (0 or 1).SetCardCollectible
// Default is 0 and 1... all
func (client *HearthstoneAPI) SetCardCollectible(collectible int) {
	if client.cardSearch != nil {
		client.cardSearch.optionalInt["collectible"] = collectible
	}
}

// SetCardKeyword filtered cards search to only ones containing the keyword
func (client *HearthstoneAPI) SetCardKeyword(keyword string) {
	if client.cardSearch != nil {
		client.cardSearch.optionalString["keyword"] = keyword
	}
}

// SetCardTextFilter filtered cards search by the name of the card
func (client *HearthstoneAPI) SetCardTextFilter(textFilter string) {
	if client.cardSearch != nil {
		client.cardSearch.optionalString["textFilter"] = textFilter
	}
}

// SetCardRarity filtered cards search by the rarity of the card
func (client *HearthstoneAPI) SetCardRarity(rarity string) {
	if client.cardSearch != nil {
		client.cardSearch.optionalString["rarity"] = rarity
	}
}

// SetCardMinionType filtered cards search by the minion type of the card
func (client *HearthstoneAPI) SetCardMinionType(minionType string) {
	if client.cardSearch != nil {
		client.cardSearch.optionalString["minionType"] = minionType
	}
}
