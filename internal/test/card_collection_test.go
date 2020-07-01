package hearthstone_go_test

import "testing"

func TestCardCollection(t *testing.T) {
	client.BeginCardCollectionSearch()

	if cards := client.EndCardCollectionSearch(); cards != nil {
		if cards.CardCount == 0 {
			t.Error("Found zero cards. Should have found at least one.")
		}

		if cards.CardCount != len(cards.Cards) {
			t.Errorf("Card count should be %d, but got %d!", cards.CardCount, len(cards.Cards))
		}
	}
}

func TestCardCollectionByName(t *testing.T) {
	client.BeginCardCollectionSearch()
	client.SetCardTextFilter("lookout")

	if cards := client.EndCardCollectionSearch(); cards != nil {
		if cards.CardCount == 0 {
			t.Error("Found zero cards. Should have found at least one.")
		}

		if cards.CardCount != len(cards.Cards) {
			t.Errorf("Card count should be %d, but got %d!", cards.CardCount, len(cards.Cards))
		}

		if cards.Cards[0].Name != "Primalfin Lookout" {
			t.Errorf("First card should be %s, but got %s!", "Primalfin Lookout", cards.Cards[0].Name)
		}
	}
}

func TestBattleGroundsCardCollection(t *testing.T) {
	client.BeginCardCollectionSearch()
	client.SetCardGameMode("battlegrounds")
	client.SetCardTiers([]int{
		4,
		6,
	})

	if cards := client.EndCardCollectionSearch(); cards != nil {
		if cards.CardCount == 0 {
			t.Error("Found zero cards. Should have found at least one.")
		}

		if cards.CardCount != len(cards.Cards) {
			t.Errorf("Card count should be %d, but got %d!", cards.CardCount, len(cards.Cards))
		}
	}
}

func TestBattleGroundsHeroCardCollection(t *testing.T) {
	client.BeginCardCollectionSearch()
	client.SetCardGameMode("battlegrounds")
	client.SetCardHero()

	if cards := client.EndCardCollectionSearch(); cards != nil {
		if cards.CardCount == 0 {
			t.Error("Found zero cards. Should have found at least one.")
		}

		if cards.CardCount != len(cards.Cards) {
			t.Errorf("Card count should be %d, but got %d!", cards.CardCount, len(cards.Cards))
		}
	}
}

func TestAllCardCollection(t *testing.T) {
	client.BeginCardCollectionSearch()
	client.SetCardGameMode("constructed")
	client.SetCardSet("rise-of-shadows")
	client.SetCardClass("mage")
	client.SetCardManaCost(10)
	client.SetCardAttack(4)
	client.SetCardHealth(10)
	client.SetCardCollectible(1)
	client.SetCardRarity("legendary")
	client.SetCardRarity("legendary")
	client.SetCardKeyword("battlecry")
	client.SetCardTextFilter("kalecgos")
	client.SetCardMinionType("dragon")

	if cards := client.EndCardCollectionSearch(); cards != nil {
		if cards.CardCount != 1 {
			t.Error("Found more or less than one card. Should have found only one.")
			return
		}

		if cards.CardCount != len(cards.Cards) {
			t.Errorf("Card count should be %d, but got %d!", cards.CardCount, len(cards.Cards))
		}

		card := cards.Cards[0]
		if card.ID != 53002 {
			t.Errorf("Card ID should be %d, but got %d!", 53002, card.ID)
		}
	}
}
