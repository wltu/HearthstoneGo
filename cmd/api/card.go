package api

import "fmt"

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

// CardSearch provides parameters for a search
type CardSearch struct {
	url string
	id  string

	// Optional Parameters
	optional map[string]string
}

// NewCardSearch acts as a constructor for CardSearch
func (client *HeartstoneAPI) NewCardSearch(id string) CardSearch {
	// Required parameters
	cardSearch := CardSearch{url: client.apiURL, id: id, optional: make(map[string]string)}

	return cardSearch
}

// SetID update the current id value for CardSearch
func (cardSearch *CardSearch) SetID(id string) {
	cardSearch.id = id
}

// SetLocale set the optional parameter of locale for CardSearch
func (cardSearch *CardSearch) SetLocale(locale string) {
	cardSearch.optional["locale"] = locale
}

// SetGameMode set the optional parameter of game mode for CardSearch
func (cardSearch *CardSearch) SetGameMode(gameMode string) {
	cardSearch.optional["gameMode"] = gameMode
}

// Execute construct api url for CardSearch
func (cardSearch *CardSearch) Execute() string {
	// card := Card{}
	url := cardSearch.url + "hearthstone/cards/" + cardSearch.id + "?"

	for key, element := range cardSearch.optional {
		fmt.Println("Key:", key, "=>", "Element:", element)
		url += key + "=" + element + "&"
	}

	return url
}
