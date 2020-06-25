package hearthstone_test

import (
	"testing"

	"github.com/wltu/HearthstoneGo/cmd/hearthstone"
)

func TestHello(t *testing.T) {
	want := "Hello, world."
	if got := hearthstone.Hello(); got != want {
		t.Errorf("Hello() = %q, want %q", got, want)
	}
}
