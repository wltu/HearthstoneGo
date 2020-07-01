package hearthstone_go_test

import (
	"os"
	"testing"

	"github.com/wltu/HearthstoneGo/cmd/api"
)

var client *api.HearthstoneAPI
var ok bool

func TestMain(m *testing.M) {
	clientID := os.Getenv("CLIENT_ID")
	clientSecret := os.Getenv("CLIENT_SECRET")
	ok = false
	client, ok = api.NewAPI("USA", clientID, clientSecret)

	os.Exit(m.Run())
}

func TestAuth(t *testing.T) {
	if !ok {
		t.Error("Authorization failed for Hearthston API")
	}
}
