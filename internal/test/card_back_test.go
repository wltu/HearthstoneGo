package hearthstone_go_test

import "testing"

func TestCardBack(t *testing.T) {
	cardBack := client.SearchCardBack("155-pizza-stone")

	if cardBack.Name != "Pizza Stone" {
		t.Errorf("Card back name should be %s, but got %s!", "Pizza Stone", cardBack.Name)
	}

	if cardBack.ID != 155 {
		t.Errorf("Card back ID should be %d, but got %d!", 155, cardBack.ID)
	}
}

func TestFakeCardBack(t *testing.T) {
	cardBack := client.SearchCardBack("-1")

	if cardBack.Name != "" {
		t.Errorf("Card back name should be %s, but got %s!", "", cardBack.Name)
	}

	if cardBack.ID != 0 {
		t.Errorf("Card back ID should be %d, but got %d!", 0, cardBack.ID)
	}
}
