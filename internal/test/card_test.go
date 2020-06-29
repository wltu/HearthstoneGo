package hearthstone_go_test

import (
	"testing"
)

func TestCard(t *testing.T) {
	card := client.SearchCard("52119-arch-villain-rafaam")

	if card.Name != "Arch-Villain Rafaam" {
		t.Errorf("Card name should be %s, but got %s!", "Arch-Villain Rafaam", card.Name)
	}

	if card.ID != 52119 {
		t.Errorf("Card ID should be %d, but got %d!", 52119, card.ID)
	}
}

func TestFakeCard(t *testing.T) {
	card := client.SearchCard("1111")

	if card.Name != "" {
		t.Errorf("Card name should be %s, but got %s!", "", card.Name)
	}

	if card.ID != 0 {
		t.Errorf("Card ID should be %d, but got %d!", 0, card.ID)
	}
}

func TestCardCollection(t *testing.T) {
	cards := client.SearchCardCollection()

	if cards.CardCount != len(cards.Cards) {
		t.Errorf("Card count should be %d, but got %d!", cards.CardCount, len(cards.Cards))
	}
}
