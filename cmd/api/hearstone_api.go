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

	oauthURL string // Hearstone OAuth URL
	apiURL   string // Regional Hearstone API URL

	heartstoneClient *http.Client
}

// AuthorizationJSON is the json structure returned after OAuth
type AuthorizationJSON struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int    `json:"expires_in"`
}

// New acts as a constructor to initialize the HearstoneAPI
func New(region, clientID, clientSecret string) HeartstoneAPI {
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
	authorization := AuthorizationJSON{}
	err := client.Authorization(client.oauthURL, &authorization)
	if err != nil {
		panic(err)
	}

	client.ClientToken = authorization.AccessToken
}
