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

	client := api.NewAPI(os.Args[1], os.Args[2], os.Args[3])
	fmt.Println(client.ClientToken)

	client.SearchCard("52119-arch-villain-rafaam")
}
