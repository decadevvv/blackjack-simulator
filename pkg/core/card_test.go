// Copyright 2022 decadevvv

package card

import (
	"fmt"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := NewDeck()
	fmt.Printf("new deck: %v\n", deck)
	fmt.Printf("new deck length: %d\n", deck.Len())
}

func TestCards_Shuffle(t *testing.T) {
	shuffledDeck := NewDeck().Shuffle()
	fmt.Printf("shuffled deck: %v\n", shuffledDeck)
	fmt.Printf("shuffled deck length: %d\n", shuffledDeck.Len())
	reshuffledDeck := shuffledDeck.Shuffle()
	fmt.Printf("re-shuffled deck: %v\n", reshuffledDeck)
	fmt.Printf("re-shuffled deck length: %d\n", reshuffledDeck.Len())
}

func TestNewShoe(t *testing.T) {
	shoe := NewShoe(4, 0)
	fmt.Printf("shoe cards: %v\n", shoe.inShoeCards)
	fmt.Printf("shoe size: %d\n", shoe.Size())
	fmt.Printf("shoe length: %d\n", shoe.InShoeLen())
	shoe.Reshuffle()
	fmt.Printf("shuffled shoe cards: %v\n", shoe.inShoeCards)
	fmt.Printf("shuffled shoe size: %d\n", shoe.Size())
	fmt.Printf("shuffled shoe length: %d\n", shoe.InShoeLen())
}

func TestShoe_Deal(t *testing.T) {
	fmt.Printf("------penetration threshold = 0.8------\n")
	shoe := NewShoe(2, 0.8)
	for i := 1; i <= 1000; i++ {
		card := shoe.Deal()
		fmt.Printf("round %d: deal %v, shoe remaining %d, penetration %f\n", i, card, shoe.InShoeLen(), shoe.Penetration())
		if i%10 == 0 {
			shoe.StageToGarbage()
		}
	}
	fmt.Printf("------penetration threshold = 0------\n")
	shoe = NewShoe(2, 0)
	for i := 1; i <= 1000; i++ {
		card := shoe.Deal()
		fmt.Printf("round %d: deal %v, shoe remaining %d, penetration %f\n", i, card, shoe.InShoeLen(), shoe.Penetration())
		if i%10 == 0 {
			shoe.StageToGarbage()
		}
	}
}

func TestShoe_Cards(t *testing.T) {
	shoe := NewShoe(1, 0.6)
	for i := 1; i <= 100; i++ {
		card := shoe.Deal()
		fmt.Printf("round %d: deal %v, %d shoe %d stage %d garbage (%f), stage %v, garbage %v, shoe %v\n", i, card, shoe.InShoeLen(), shoe.InStageLen(), shoe.InGarbageLen(), shoe.Penetration(), shoe.inStageCards, shoe.inGarbageCards, shoe.inShoeCards)
		if i%10 == 0 {
			shoe.StageToGarbage()
		}
	}
}
