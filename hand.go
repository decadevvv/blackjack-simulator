// Copyright 2022 decadevvv

package main

import (
	"fmt"
)

type Hand struct {
	rules    Rules
	cards    Cards
	split    bool
	doubled  bool
	aces     uint8
	softAces uint8
	points   uint8
}

func NewHand(split bool, rules Rules) *Hand {
	return &Hand{
		rules:   rules,
		cards:   []Card{},
		split:   split,
		doubled: false,
		aces:    0,
	}
}

func (h *Hand) Add(c Card) *Hand {
	// add the card to remainingCards
	h.cards = append(h.cards, c)
	// update the number of aces
	if c == CardA {
		h.aces++
	}
	// re-calculate the points and number of soft aces (aces valued at 11 when calculating points)
	points := uint8(0)
	for _, card := range h.cards {
		points += card.Point()
	}
	softAces := h.aces
	if points > 21 && h.Aces() > 0 {
		for ace := uint8(1); ace <= h.Aces(); ace++ {
			softAces--
			points -= 10
			if points <= 21 {
				break
			}
		}
	}
	h.points = points
	h.softAces = softAces
	return h
}

func (h *Hand) Cards() Cards {
	return h.cards
}

func (h *Hand) Len() int {
	return h.cards.Len()
}

// Points returns the total points of the hand (Ace counted as either 1 or 11)
func (h *Hand) Points() uint8 {
	return h.points
}

// SoftAces returns the number of soft aces (Aces valued at 11 when calculating points) in the hand
func (h *Hand) SoftAces() uint8 {
	return h.softAces
}

// Aces returns the number of Aces in the hand
func (h *Hand) Aces() uint8 {
	return h.aces
}

// Soft returns whether the current hand is soft (meaning that it consists of aces valued at 11)
func (h *Hand) Soft() bool {
	return h.softAces > 0
}

// Busted returns whether the current hand is busted (points > 21)
func (h *Hand) Busted() bool {
	return h.points > 21
}

// BlackJack returns whether the current hand is a blackjack (Ace + 10-point card)
func (h *Hand) BlackJack() bool {
	return (!h.split) && (h.points == 21) && (h.cards.Len() == 2 || (h.rules.AllowTriple7BlackJack && h.cards.Len() == 3 && h.cards[0] == Card7 && h.cards[1] == Card7 && h.cards[2] == Card7))
}

// CanSplit returns whether the current hand can canSplit
func (h *Hand) CanSplit() bool {
	return (!h.split) && h.cards.Len() == 2 && h.cards[0] == h.cards[1]
}

func (h *Hand) CouldSplit() bool {
	return (!h.split) && h.cards[0] == h.cards[1]
}

// Split splits the current hand
func (h *Hand) Split() (*Hand, *Hand) {
	hand1 := NewHand(true, h.rules)
	hand2 := NewHand(true, h.rules)
	hand1.Add(h.cards[0])
	hand2.Add(h.cards[1])
	return hand1, hand2
}

func (h *Hand) IsSplit() bool {
	return h.split
}

func (h *Hand) CanDouble() bool {
	return h.cards.Len() == 2 && (h.rules.AllowAceAceSplitAndDouble || !h.IsSplit() || h.cards[0] != CardA)
}

func (h *Hand) Double() {
	h.doubled = true
}

func (h *Hand) Doubled() bool {
	return h.doubled
}

func (h *Hand) String() string {
	return fmt.Sprintf("%v (%d)", h.cards, h.points)
}
