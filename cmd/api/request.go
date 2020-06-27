package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func (client *HearthstoneAPI) authorization(url string, authorization *Authorization) error {
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

	return json.NewDecoder(res.Body).Decode(authorization)
}

func post(client *http.Client, url string, target interface{}) error {
	request, err := http.NewRequest("POST", url, nil)

	if err != nil {
		return err
	}

	res, err := client.Do(request)

	if err != nil {
		return err
	}

	return json.NewDecoder(res.Body).Decode(target)
}

func get(client *http.Client, url string, target interface{}) error {
	request, err := http.NewRequest("GET", url, nil)

	if err != nil {
		panic(err)
	}

	res, err := client.Do(request)

	if err != nil {
		panic(err)
	}

	return json.NewDecoder(res.Body).Decode(target)
}

func readBody(res http.Response) {
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)
}
