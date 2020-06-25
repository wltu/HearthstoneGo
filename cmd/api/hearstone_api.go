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

// Execute request Heartstone Informations
func (client *HeartstoneAPI) Execute(request interface{}) {

	if endpoint, ok := request.(CardSearch); ok { // type assert on it
		url := endpoint.Execute()
		url += "access_token=" + client.ClientToken
		card := Card{}

		fmt.Println(url)
		err := client.Get(url, &card)
		if err != nil {
			panic(err)
		}
		v := reflect.ValueOf(card)
		typeOfS := v.Type()

		for i := 0; i < v.NumField(); i++ {
			fmt.Printf("Field: %s\tValue: %v\n", typeOfS.Field(i).Name, v.Field(i).Interface())
		}
	}
}
