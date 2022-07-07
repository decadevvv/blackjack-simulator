// Copyright 2022 decadevvv

package main

import (
	"math/rand"
)

type Card byte

const (
	CardA  = 1
	Card2  = 2
	Card3  = 3
	Card4  = 4
	Card5  = 5
	Card6  = 6
	Card7  = 7
	Card8  = 8
	Card9  = 9
	Card10 = 10
	CardJ  = 11
	CardQ  = 12
	CardK  = 13
)

var CardName = map[Card]string{
	CardA:  "A",
	Card2:  "2",
	Card3:  "3",
	Card4:  "4",
	Card5:  "5",
	Card6:  "6",
	Card7:  "7",
	Card8:  "8",
	Card9:  "9",
	Card10: "10",
	CardJ:  "J",
	CardQ:  "Q",
	CardK:  "K",
}

var CardPoint = map[Card]uint8{
	CardA:  11,
	Card2:  2,
	Card3:  3,
	Card4:  4,
	Card5:  5,
	Card6:  6,
	Card7:  7,
	Card8:  8,
	Card9:  9,
	Card10: 10,
	CardJ:  10,
	CardQ:  10,
	CardK:  10,
}

func (c Card) String() string {
	return CardName[c]
}

func (c Card) Point() uint8 {
	return CardPoint[c]
}

var CardEnums = []Card{
	CardA,
	Card2,
	Card3,
	Card4,
	Card5,
	Card6,
	Card7,
	Card8,
	Card9,
	Card10,
	CardJ,
	CardQ,
	CardK,
}

const DeckSize = 52

func NewDeck() Cards {
	var newDeck []Card
	for _, card := range CardEnums {
		for i := 1; i <= 4; i++ {
			newDeck = append(newDeck, card)
		}
	}
	return newDeck
}

type Cards []Card

func (c Cards) Len() int {
	return len(c)
}

func (c Cards) Less(i, j int) bool {
	return c[i] < c[j]
}

func (c Cards) Swap(i, j int) {
	c[i], c[j] = c[j], c[i]
}

func (c Cards) Append(o ...Card) Cards {
	return append(c, o...)
}

func (c Cards) Shuffle() Cards {
	rand.Shuffle(len(c), func(i, j int) {
		c[i], c[j] = c[j], c[i]
	})
	return c
}

type Shoe struct {
	inShoeCards          Cards
	inStageCards         Cards
	inGarbageCards       Cards
	size                 uint
	originalLen          uint
	penetrationThreshold float64
}

func NewShoe(size uint, penetrationThreshold float64) *Shoe {
	shoe := Shoe{
		inShoeCards:          []Card{},
		inStageCards:         []Card{},
		inGarbageCards:       []Card{},
		size:                 size,
		originalLen:          DeckSize * size,
		penetrationThreshold: penetrationThreshold,
	}
	for i := uint(1); i <= size; i++ {
		shoe.inShoeCards = shoe.inShoeCards.Append(NewDeck()...)
	}
	return &shoe
}

func (s *Shoe) Size() uint {
	return s.size
}

func (s *Shoe) InShoeLen() int {
	return s.inShoeCards.Len()
}

func (s *Shoe) InStageLen() int {
	return s.inStageCards.Len()
}

func (s *Shoe) InGarbageLen() int {
	return s.inGarbageCards.Len()
}

// Deal returns the next card off the top on the shoe and whether the shoe is reshuffled
// (if the shoe has 0 cards or penetration threshold is reached, the shoe gets reshuffled)
func (s *Shoe) Deal() Card {
	if s.inShoeCards.Len() == 0 || s.Penetration() <= s.penetrationThreshold {
		s.Reshuffle()
	}
	card := s.inShoeCards[0]
	s.inShoeCards = s.inShoeCards[1:]
	s.inStageCards = append(s.inStageCards, card)
	return card
}

// Penetration returns the ratio of remainingCards that are still in the shoe to all initial remainingCards
func (s *Shoe) Penetration() float64 {
	return float64(s.inShoeCards.Len()) / float64(s.originalLen)
}

func (s *Shoe) Reshuffle() {
	s.inShoeCards = s.inShoeCards.Append(s.inGarbageCards...).Shuffle()
	s.inGarbageCards = []Card{}
}

func (s *Shoe) StageToGarbage() {
	s.inGarbageCards = append(s.inGarbageCards, s.inStageCards...)
	s.inStageCards = []Card{}
}
