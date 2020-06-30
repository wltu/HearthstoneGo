package hearthstone_go_test

import (
	"testing"
)

func TestCard(t *testing.T) {

	if card, ok := client.SearchCard("52119-arch-villain-rafaam"); ok {
		if card.Name != "Arch-Villain Rafaam" {
			t.Errorf("Card name should be %s, but got %s!", "Arch-Villain Rafaam", card.Name)
		}

		if card.ID != 52119 {
			t.Errorf("Card ID should be %d, but got %d!", 52119, card.ID)
		}
	} else {
		t.Error("Card not found or is battlegrounds only!")
	}
}

func TestBattleGroundOnlyCard(t *testing.T) {
	if card, ok := client.SearchCard("60040-zapp-slywick"); ok {
		t.Errorf("Card should be battleground only. Got %s!", card)
	}
}
func TestFakeCard(t *testing.T) {
	if card, ok := client.SearchCard("1111"); ok {
		t.Errorf("Card should be missing, but found %s!", card.Name)
	}
}

func TestBattlegroundsCard(t *testing.T) {

	if card, ok := client.SearchBattlegroundsCard("60040-zapp-slywick"); ok {
		if card.Name != "Zapp Slywick" {
			t.Errorf("Card name should be %s, but got %s!", "Zapp Slywick", card.Name)
		}

		if card.ID != 60040 {
			t.Errorf("Card ID should be %d, but got %d!", 60040, card.ID)
		}

		if card.Battlegrounds.Tier != 6 {
			t.Errorf("Card tier should be %d, but got %d!", 6, card.Battlegrounds.Tier)
		}
	} else {
		t.Error("Card not found or is constructed only!")
	}
}

func TestConstructedOnlyCard(t *testing.T) {
	if card, ok := client.SearchBattlegroundsCard("52119-arch-villain-rafaam"); ok {
		t.Errorf("Card should be battleground only. Got %s!", card)
	}
}
func TestFakeBattleGroundsCard(t *testing.T) {
	if card, ok := client.SearchBattlegroundsCard("1111"); ok {
		t.Errorf("Card should be missing, but found %s!", card.Name)
	}
}

func TestCardCollection(t *testing.T) {
	cards := client.SearchCardCollection()

	if cards.CardCount != len(cards.Cards) {
		t.Errorf("Card count should be %d, but got %d!", cards.CardCount, len(cards.Cards))
	}
}
