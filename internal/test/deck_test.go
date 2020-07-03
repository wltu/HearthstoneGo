package hearthstone_go_test

import "testing"

func TestDeck(t *testing.T) {
	id := "AAECAQcG+wyd8AKS+AKggAOblAPanQMMS6IE/web8wLR9QKD+wKe+wKz/AL1gAOXlAOalAOSnwMA"

	if deck := client.SearchDeck(id); deck != nil {

		if deck.DeckCode != id {
			t.Errorf("Deck code should be %s, but got %s!", id, deck.DeckCode)
		}

		if deck.CardCount != len(deck.Cards) {
			t.Errorf("Deck should have %d cards, but got %d!", deck.CardCount, len(deck.Cards))
		}
	}
}
