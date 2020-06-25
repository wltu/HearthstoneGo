package api

import (
	"encoding/json"
	"net/http"
)

// Authorization request to Hearstone API
func (client *HeartstoneAPI) Authorization(url string, target interface{}) error {
	request, err := http.NewRequest("POST", url, nil)
	request.SetBasicAuth(client.ClientID, client.ClientSecret)

	if err != nil {
		return err
	}

	res, err := client.heartstoneClient.Do(request)

	if err != nil {
		return err
	}

	defer res.Body.Close()

	return json.NewDecoder(res.Body).Decode(target)
}

// Post request to Hearstone API
func (client *HeartstoneAPI) Post(url string, target interface{}) error {
	request, err := http.NewRequest("POST", url, nil)

	if err != nil {
		return err
	}

	res, err := client.heartstoneClient.Do(request)

	if err != nil {
		return err
	}

	return json.NewDecoder(res.Body).Decode(target)
}

// Get request to Hearstone API
func (client *HeartstoneAPI) Get(url string, target interface{}) error {
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	res, err := client.heartstoneClient.Do(request)

	if err != nil {
		panic(err)
	}

	return json.NewDecoder(res.Body).Decode(target)
}
