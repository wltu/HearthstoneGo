package api

import (
	"fmt"
	"net/http"
)

// Metadata provides information to categorize the cards
type Metadata struct {
	Sets               []Set              `json:"sets"`
	SetGroups          []SetGroup         `json:"setGroups"`
	ArenaIds           []int              `json:"arenaIds"`
	Types              []Type             `json:"types"`
	Rarities           []Rarity           `json:"rarities"`
	Classes            []Class            `json:"classes"`
	MinionTypes        []MinionType       `json:"minionTypes"`
	GameModes          []GameMode         `json:"gameModes"`
	Keywords           []Keyword          `json:"keywords"`
	FilterableFields   []string           `json:"filterableFields"`
	NumericFields      []string           `json:"numericFields"`
	CardBackCategories []CardBackCategory `json:"cardBackCategories"`
}

// Set provides information of Card Sets
type Set struct {
	ID                          int    `json:"id"`
	Name                        string `json:"name"`
	Slug                        string `json:"slug"`
	ReleaseDate                 string `json:"releaseDate"`
	Type                        string `json:"type"`
	CollectibleCount            int    `json:"collectibleCount"`
	CollectibleRevealedCount    int    `json:"collectibleRevealedCount"`
	NonCollectibleCount         int    `json:"nonCollectibleCount"`
	NonCollectibleRevealedCount int    `json:"nonCollectibleRevealedCount"`
}

// SetGroup provides information of Card Set Group
type SetGroup struct {
	Slug     string   `json:"slug"`
	Year     int      `json:"year"`
	CardSets []string `json:"cardSets"`
	Name     string   `json:"name"`
	Standard bool     `json:"standard"`
	Icon     string   `json:"icon"`
}

// Type provides the different type of cards in Hearthstone
type Type struct {
	Slug string `json:"slug"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Rarity provides the different rarities of cards in Hearthstone
type Rarity struct {
	Slug         string `json:"slug"`
	ID           int    `json:"id"`
	CraftingCost [2]int `json:"craftingCost"`
	DustValue    [2]int `json:"dustValue"`
	Name         string `json:"name"`
}

// Class provides the different classes of heros in Hearthstone
type Class struct {
	Slug   string `json:"slug"`
	ID     int    `json:"id"`
	Name   string `json:"name"`
	CardID int    `json:"cardId"`
}

// String function for Hero Class
func (class Class) String() string {
	return fmt.Sprintf(class.Name)
}

// MinionType provides the different types of minions in Hearthstone
type MinionType struct {
	Slug string `json:"slug"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// GameMode provides the different game modes in Hearthstone
type GameMode struct {
	Slug string `json:"slug"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Keyword provides the different key word effect for the cards in Hearthstone
type Keyword struct {
	ID      int    `json:"id"`
	Slug    string `json:"slug"`
	Name    string `json:"name"`
	RefText string `json:"refText"`
	Text    string `json:"text"`
}

// CardBackCategory provides the different kind of card backs in Hearthstone
type CardBackCategory struct {
	Slug string `json:"slug"`
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type metadataSearch struct {
	url    string
	locale string
}

// newMetadataSearch acts as a constructor for metadataSearch
func (client *HearthstoneAPI) newMetadataSearch() metadataSearch {
	// Required parameters
	return metadataSearch{
		url:    client.apiURL,
		locale: client.locale,
	}
}

func (search *metadataSearch) execute(client *http.Client, token string) interface{} {
	url := search.url +
		"hearthstone/metadata/?locale=" +
		search.locale + "&" +
		"access_token=" + token

	metadata := Metadata{}
	err := get(client, url, &metadata)

	if err != nil {
		panic(err)
	}

	// print(metadata)

	return metadata
}

func (client *HearthstoneAPI) setMetadata(metadata *Metadata) {
	size := len(metadata.Sets)
	client.sets = make([]string, size)

	for i := 0; i < size; i++ {
		client.sets[i] = metadata.Sets[i].Slug
	}

	size = len(metadata.SetGroups)
	client.setGroups = make([]string, size)

	for i := 0; i < size; i++ {
		client.setGroups[i] = metadata.SetGroups[i].Slug
	}

	size = len(metadata.Types)
	client.types = make([]string, size)

	for i := 0; i < size; i++ {
		client.types[i] = metadata.Types[i].Slug
	}

	size = len(metadata.Rarities)
	client.rarities = make([]string, size)

	for i := 0; i < size; i++ {
		client.rarities[i] = metadata.Rarities[i].Slug
	}

	size = len(metadata.Classes)
	client.classes = make([]string, size)

	for i := 0; i < size; i++ {
		client.classes[i] = metadata.Classes[i].Slug
	}

	size = len(metadata.MinionTypes)
	client.minionTypes = make([]string, size)

	for i := 0; i < size; i++ {
		client.minionTypes[i] = metadata.MinionTypes[i].Slug
	}

	size = len(metadata.GameModes)
	client.gameModes = make([]string, size)

	for i := 0; i < size; i++ {
		client.gameModes[i] = metadata.GameModes[i].Slug
	}

	size = len(metadata.Keywords)
	client.keywords = make([]string, size)

	for i := 0; i < size; i++ {
		client.keywords[i] = metadata.Keywords[i].Slug
	}

	size = len(metadata.CardBackCategories)
	client.cardBackCategories = make([]string, size)

	for i := 0; i < size; i++ {
		client.cardBackCategories[i] = metadata.CardBackCategories[i].Slug
	}
}

// Sets returns all the available card sets in Hearthstone
func (client *HearthstoneAPI) Sets() []string {
	return client.sets
}

// SetGroups returns all the available card set groups in Hearthstone
func (client *HearthstoneAPI) SetGroups() []string {
	return client.setGroups
}

// Types returns all the available card types in Hearthstone
func (client *HearthstoneAPI) Types() []string {
	return client.types
}

// Rarities returns all the available card rarities in Hearthstone
func (client *HearthstoneAPI) Rarities() []string {
	return client.rarities
}

// Classes returns all the available hero classes in Hearthstone
func (client *HearthstoneAPI) Classes() []string {
	return client.classes
}

// MinionTypes returns all the available card minion types in Hearthstone
func (client *HearthstoneAPI) MinionTypes() []string {
	return client.minionTypes
}

// Keywords returns all the available card keywords in Hearthstone
func (client *HearthstoneAPI) Keywords() []string {
	return client.keywords
}

// CardBackCategories returns all the available categories of card backs in Hearthstone
func (client *HearthstoneAPI) CardBackCategories() []string {
	return client.cardBackCategories
}
