package main

import (
	"fmt"
	"os"

	"github.com/wltu/HearthstoneGo/cmd/api"
	"github.com/wltu/HearthstoneGo/cmd/hearthstone"
)

func main() {
	fmt.Println("Hello World!")
	hearthstone.Hello()

	if len(os.Args) < 4 {
		fmt.Println("Not enough arguments! Please include client id and secret.")
		return
	}

	if client, ok := api.NewAPI("en_US", os.Args[1], os.Args[2], os.Args[3]); ok {
		fmt.Println(client.ClientToken)

		// client.SearchCard("52119-arch-villain-rafaam")
		// fmt.Println()
		// client.SearchCard("56363")
		// cardCollrection := client.SearchCardCollection()
		// client.SearchCardBack("155-pizza-stone")
		// client.SearchCardBackCollection()
		client.SearchDeck("AAECAQcG+wyd8AKS+AKggAOblAPanQMMS6IE/web8wLR9QKD+wKe+wKz/AL1gAOXlAOalAOSnwMA")

		// fmt.Println(cardCollrection.Cards[0])
	} else {
		fmt.Println("Error in setting up HearthstoneAPI Client!")
	}
}
