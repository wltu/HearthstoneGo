package api

import (
	"fmt"
	"net"
	"net/http"
	"reflect"
	"time"
)

// Error is used to handle missing authorization for Hearthstone API
type Error struct {
}

func (e *Error) Error() string {
	return "Required Enviroment Variables Missing! Please include client id and secret."
}

// HearthstoneAPI connects to Blizzard API for Hearstone Informations
type HearthstoneAPI struct {
	// Client Log in information
	ClientID, ClientSecret string
	ClientToken            string

	localeMap map[string]string
	locale    string
	oauthURL  string // Hearstone OAuth URL
	apiURL    string // Regional Hearstone API URL

	heartstoneClient *http.Client

	// Metadata to filter the cards
	sets               []string
	setGroups          []string
	types              []string
	rarities           []string
	classes            []string
	minionTypes        []string
	gameModes          []string
	keywords           []string
	cardBackCategories []string
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

	if len(clientID) == 0 ||
		len(clientSecret) == 0 {
		panic(&Error{})
	}

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
		client.setMetadata(&output)
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
