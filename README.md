# HearthstoneGo
Unofficial library of accessing Hearthstone game data in Go

![Windows Go](https://github.com/wltu/HearthstoneGo/workflows/Windows%20Go/badge.svg?branch=master)
![Mac Go](https://github.com/wltu/HearthstoneGo/workflows/Mac%20Go/badge.svg?branch=master)
![Linux Go](https://github.com/wltu/HearthstoneGo/workflows/Linux%20Go/badge.svg?branch=master)

The HearthstoneGo library provides the developer easy access to the Blizzard Hearthstone API. It handels all of the official API calls and supply the developer all the information they need.

## Get Getting Started
### Installing
Given that you have working Golang environment. If not refer to [here](https://golang.org/doc/install).
```
go get github.com/wltu/HearthstoneGo
```

### Simple Example
The `CLIENT_ID` and `CLIENT_SECRET` is created for the Blizzard API. Please follow the instruction [here](https://develop.battle.net/documentation/guides/getting-started) to create your own to use the library.
```
package main

import (
	"fmt"
	"os"

	"github.com/wltu/HearthstoneGo/cmd/api"
)

func main() {
	fmt.Println("Hello World!")

	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")

	if client, ok := api.NewAPI("USA", clientID, clientSecret); ok {
		fmt.Println(client.ClientToken)

		// Search for single card.
		client.SearchCard("52119-arch-villain-rafaam")

		// Search for single card back.
		client.SearchCardBack("155-pizza-stone")

		// Search for a set of cards
		client.BeginCardCollectionSearch()

		// Set optional parameters.
		// Visit card_collection.go for more info.
		client.SetCardTextFilter("lookout")

		client.EndCardCollectionSearch()

		// Search for a set of card backs
		client.BeginCardBackCollectionSearch()

		// Set optional parameters.
		// Visit card_back_collection.go for more info.
		client.SetCardBackCategory("esports")

		client.SetCardTextFilter("lookout")

		client.EndCardBackCollectionSearch()

		// Search for deck
		id := "AAECAQcG+wyd8AKS+AKggAOblAPanQMMS6IE/web8wLR9QKD+wKe+wKz/AL1gAOXlAOalAOSnwMA"
		client.SearchDeck(id)

	} else {
		fmt.Println("Error in setting up HearthstoneAPI Client!")
	}
}

```


## Project Structure
This project follows loosely the proejct structure [here](https://github.com/golang-standards/project-layout).

## Design Document
The rough design document for the project can be found [here](https://docs.google.com/document/d/1hwWPqrOF7vG7u6qqmdCPqRR4Js99LyKEcchpjR17Z3E/edit?usp=sharing).