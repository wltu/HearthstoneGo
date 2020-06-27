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

	oauthURL string // Hearstone OAuth URL
	apiURL   string // Regional Hearstone API URL

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
func NewAPI(region, clientID, clientSecret string) HearthstoneAPI {
	regionMap := map[string]string{
		"us": "https://us.api.blizzard.com/",
		"eu": "https://eu.api.blizzard.com/",
		"kr": "https://kr.api.blizzard.com/",
		"tw": "https://tw.api.blizzard.com/",
		"ch": "https://gateway.battlenet.com.cn/",
	}

	client := HearthstoneAPI{ClientID: clientID,
		ClientSecret: clientSecret,
		oauthURL:     "https://us.battle.net/oauth/token?grant_type=client_credentials",
		apiURL:       regionMap[region]}

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

	return client
}

func (client *HearthstoneAPI) connect() {
	auth := Authorization{}
	err := client.authorization(client.oauthURL, &auth)
	if err != nil {
		panic(err)
	}

	client.ClientToken = auth.AccessToken
}

// SearchCard make a API call to search for a card with the given id
func (client *HearthstoneAPI) SearchCard(id string) Card {
	cardSearch := client.newCardSearch(id, "en_US")

	if output, ok := client.execute(&cardSearch).(Card); ok {
		return output
	}

	return Card{}
}

func print(body interface{}) {
	var value reflect.Value

	switch v := body.(type) {
	case Card:
		if card, ok := body.(Card); ok {
			value = reflect.ValueOf(card)
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
