package api

import (
	"net"
	"net/http"
	"time"
)

// HeartstoneAPI connects to Blizzard API for Hearstone Informations
type HeartstoneAPI struct {
	// Client Log in information
	ClientID, ClientSecret string
	ClientToken            string

	// Defaulr Hearstone API URL
	hearstoneURL string

	heartstoneClient *http.Client
}

// AuthorizationJSON is the json structure returned after OAuth
type AuthorizationJSON struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// New acts as a constructor to initialize the HearstoneAPI
func New(clientID, clientSecret string) HeartstoneAPI {
	client := HeartstoneAPI{ClientID: clientID,
		ClientSecret: clientSecret,
		hearstoneURL: "https://us.battle.net/"}

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
	url := client.hearstoneURL + "oauth/token?grant_type=client_credentials"
	authorization := AuthorizationJSON{}
	err := client.Authorization(url, &authorization)

	if err != nil {
		panic(err)
	}

	client.ClientToken = authorization.AccessToken
}
