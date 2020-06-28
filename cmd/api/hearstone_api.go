package api

import (
	"fmt"
	"net"
	"net/http"
	"reflect"
	"time"
)

// HearthstoneAPI connects to Blizzard API for Hearstone Informations
type HearthstoneAPI struct {
	// Client Log in information
	ClientID, ClientSecret string
	ClientToken            string

	localeMap map[string]string
	locale    string
	oauthURL  string // Hearstone OAuth URL
	apiURL    string // Regional Hearstone API URL

	metadata Metadata

	heartstoneClient *http.Client
}

// Authorization is the json structure returned after OAuth
type Authorization struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

type endpoint interface {
	execute(client *http.Client, url string) interface{}
}

// NewAPI acts as a constructor to initialize the HearstoneAPI
func NewAPI(locale, region, clientID, clientSecret string) (*HearthstoneAPI, bool) {

	client := HearthstoneAPI{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		oauthURL:     "https://us.battle.net/oauth/token?grant_type=client_credentials",
		apiURL:       "https://" + region + ".api.blizzard.com/",
		locale:       locale,
		localeMap: map[string]string{
			"USA":           "en_US",
			"Mexico":        "es_MX",
			"Brazil":        "pt_BR",
			"Great Britain": "en_GB",
			"Spain":         "es_ES",
			"France":        "fr_FR",
			"Russia":        "ru_RU",
			"Germany":       "de_DE",
			"Portugal":      "pt_PT",
			"Italy":         "it_IT",
			"Korea":         "ko_KR",
			"Taiwan":        "zh_TW",
			"China":         "zh_CH",
		},
	}

	netTransport := &http.Transport{
		Dial: (&net.Dialer{
			Timeout: 5 * time.Second,
		}).Dial,
		TLSHandshakeTimeout: 5 * time.Second,
	}

	client.heartstoneClient = &http.Client{
		Timeout:   time.Second * 10,
		Transport: netTransport,
	}

	client.connect()

	search := client.newMetadataSearch()

	if output, ok := client.execute(&search).(Metadata); ok {
		client.metadata = output
	} else {
		return nil, false
	}

	return &client, true
}

func (client *HearthstoneAPI) connect() {
	auth := Authorization{}
	err := client.authorization(client.oauthURL, &auth)
	if err != nil {
		panic(err)
	}

	client.ClientToken = auth.AccessToken
}

// SearchDeck make a API call to search for a deck with the given id
func (client *HearthstoneAPI) SearchDeck(id string) Deck {
	search := client.newDeckSearch(id)

	if output, ok := client.execute(&search).(Deck); ok {
		return output
	}

	return Deck{}
}

// SearchCard make a API call to search for a card with the given id
func (client *HearthstoneAPI) SearchCard(id string) Card {
	search := client.newCardSearch(id)

	if output, ok := client.execute(&search).(Card); ok {
		return output
	}

	return Card{}
}

// SearchCardCollection make a API call to search for a set of cards
func (client *HearthstoneAPI) SearchCardCollection() CardCollection {
	search := client.newCardCollectionSearch()

	if output, ok := client.execute(&search).(CardCollection); ok {
		return output
	}

	return CardCollection{}
}

// SearchCardBack make a API call to search for a card back with the given id
func (client *HearthstoneAPI) SearchCardBack(id string) CardBack {
	search := client.newCardBackSearch(id)

	if output, ok := client.execute(&search).(CardBack); ok {
		return output
	}

	return CardBack{}
}

// SearchCardBackCollection make a API call to search for a collection of card backs
func (client *HearthstoneAPI) SearchCardBackCollection() CardBackCollection {
	search := client.newCardBackCollectionSearch()

	if output, ok := client.execute(&search).(CardBackCollection); ok {
		return output
	}

	return CardBackCollection{}
}

func print(body interface{}) {
	var value reflect.Value

	switch v := body.(type) {
	case Card:
		if data, ok := body.(Card); ok {
			value = reflect.ValueOf(data)
		}
	case CardCollection:
		if data, ok := body.(CardCollection); ok {
			value = reflect.ValueOf(data)
		}
	case CardBack:
		if data, ok := body.(CardBack); ok {
			value = reflect.ValueOf(data)
		}
	case CardBackCollection:
		if data, ok := body.(CardBackCollection); ok {
			value = reflect.ValueOf(data)
		}
	case Deck:
		if data, ok := body.(Deck); ok {
			value = reflect.ValueOf(data)
		}
	case Metadata:
		if data, ok := body.(Metadata); ok {
			value = reflect.ValueOf(data)
		}
	default:
		fmt.Printf("I don't know about type %T!\n", v)
		return
	}

	typeOfS := value.Type()

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, value.Field(i).Interface())
	}
}

func (client *HearthstoneAPI) execute(request endpoint) interface{} {
	return request.execute(
		client.heartstoneClient,
		client.ClientToken,
	)
}
