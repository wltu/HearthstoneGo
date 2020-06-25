package api

import (
	"fmt"
	"net"
	"net/http"
	"reflect"
	"time"
)

// HeartstoneAPI connects to Blizzard API for Hearstone Informations
type HeartstoneAPI struct {
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

// NewAPI acts as a constructor to initialize the HearstoneAPI
func NewAPI(region, clientID, clientSecret string) HeartstoneAPI {
	regionMap := map[string]string{
		"us": "https://us.api.blizzard.com/",
		"eu": "https://eu.api.blizzard.com/",
		"kr": "https://kr.api.blizzard.com/",
		"tw": "https://tw.api.blizzard.com/",
		"ch": "https://gateway.battlenet.com.cn/",
	}

	client := HeartstoneAPI{ClientID: clientID,
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

func (client *HeartstoneAPI) connect() {
	authorization := Authorization{}
	err := client.Authorization(client.oauthURL, &authorization)
	if err != nil {
		panic(err)
	}

	client.ClientToken = authorization.AccessToken
}

func (client *HeartstoneAPI) print(body interface{}) {
	var value reflect.Value

	switch v := body.(type) {
	case Card:
		if card, ok := body.(Card); ok {
			value = reflect.ValueOf(card)
		}

	case CardAll:
		if card, ok := body.(CardAll); ok {
			value = reflect.ValueOf(card)
		}
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}

	typeOfS := value.Type()

	for i := 0; i < value.NumField(); i++ {
		fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, value.Field(i).Interface())
	}
}

func (client *HeartstoneAPI) searchCards(url string, card interface{}) interface{} {

	if cardType, ok := card.(Card); ok {
		fmt.Println(url)
		err := client.Get(url, &cardType)
		if err != nil {
			panic(err)
		}

		client.print(cardType)
		return cardType
	}

	if cardType, ok := card.(CardAll); ok {
		fmt.Println(url)
		err := client.Get(url, &cardType)
		if err != nil {
			panic(err)
		}

		client.print(cardType)

		return cardType
	}

	return nil
}

// Execute request Heartstone Informations
func (client *HeartstoneAPI) Execute(request interface{}) {

	if endpoint, ok := request.(CardSearch); ok { // type assert on it
		url, card := endpoint.Execute()
		url += "access_token=" + client.ClientToken
		client.searchCards(url, card)
	}
}
