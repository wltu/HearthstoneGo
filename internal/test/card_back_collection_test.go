package hearthstone_go_test

import "testing"

func TestCardBackCollection(t *testing.T) {
	client.BeginCardBackCollectionSearch()

	if cardBacks := client.EndCardBackCollectionSearch(); cardBacks != nil {
		if cardBacks.CardCount == 0 {
			t.Error("Found zero card backs. Should have found at least one.")
		}

		if cardBacks.CardCount != len(cardBacks.CardBacks) {
			t.Errorf("Card back count should be %d, but got %d!", cardBacks.CardCount, len(cardBacks.CardBacks))
		}
	}
}

func TestCategoryCardBackCollection(t *testing.T) {
	client.BeginCardBackCollectionSearch()
	client.SetCardBackCategory("esports")

	if cardBacks := client.EndCardBackCollectionSearch(); cardBacks != nil {
		if cardBacks.CardCount == 0 {
			t.Error("Found zero card backs. Should have found at least one.")
		}

		if cardBacks.CardCount != len(cardBacks.CardBacks) {
			t.Errorf("Card back count should be %d, but got %d!", cardBacks.CardCount, len(cardBacks.CardBacks))
		}
	}
}

func TestNameCardBackCollection(t *testing.T) {
	client.BeginCardBackCollectionSearch()
	client.SetCardBackTextFilter("Call")

	if cardBacks := client.EndCardBackCollectionSearch(); cardBacks != nil {
		if cardBacks.CardCount == 0 {
			t.Error("Found zero card backs. Should have found at least one.")
		}

		if cardBacks.CardCount != len(cardBacks.CardBacks) {
			t.Errorf("Card back count should be %d, but got %d!", cardBacks.CardCount, len(cardBacks.CardBacks))
		}
	}
}
