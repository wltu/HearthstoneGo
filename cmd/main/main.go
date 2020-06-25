package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/wltu/HearthstoneGo/cmd/api"
	"github.com/wltu/HearthstoneGo/cmd/hearthstone"
)

// curl -u 2f6c48af4ad74926a7fcd8eac98bf16e:rnJh1STqD5jhxC4KRYUfs8Ba1G1g6ObT -d grant_type=client_credentials https://us.battle.net/oauth/token
// USmp8Z5ffUk3AWseyJO1Zvx2IpIjQYQUTk
// {"access_token":"USmp8Z5ffUk3AWseyJO1Zvx2IpIjQYQUTk","token_type":"bearer","expires_in":86399}

func basicAuth(username, password string) string {
	auth := username + ":" + password
	return base64.StdEncoding.EncodeToString([]byte(auth))
}

func main() {
	fmt.Println("Hello World!")
	hearthstone.Hello()

	if len(os.Args) < 4 {
		fmt.Println("Not enough arguments! Please include client id and secret.")
		return
	}

	client := api.NewAPI(os.Args[1], os.Args[2], os.Args[3])
	fmt.Println(client.ClientToken)

	cardSearch := client.NewCardSearch("52119-arch-villain-rafaam")
	cardSearch.SetLocale("en_US")
	client.Execute(cardSearch)
}
