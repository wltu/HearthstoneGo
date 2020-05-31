package discordpp_test

import (
	"testing"

	"github.com/wltu/HearthstoneGo/cmd/discordpp"
)

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := discordpp.Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
