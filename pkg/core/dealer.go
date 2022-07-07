// Copyright 2022 decadevvv

package game

import (
	"github.com/decadevvv/blackjack-simulator/pkg/card"
	"github.com/decadevvv/blackjack-simulator/pkg/core"
)

type Dealer struct {
	hand *card.Hand
}

func NewDealer() *Dealer {
	return &Dealer{
		hand: nil,
	}
}

func (d *Dealer) InitRound(ctx *core.Context) {
	d.hand = card.NewHand(false, ctx.Rules)
	openCard := ctx.Shoe.Deal()
	d.hand.Add(openCard)
	d.hand.Add(ctx.Shoe.Deal())
	ctx.DealerOpenCard = openCard
}

func (d *Dealer) PlayRound(ctx *core.Context) {
	for d.hand.Points() < 17 {
		d.hand.Add(ctx.Shoe.Deal())
	}
}

func (d *Dealer) Hand() *card.Hand {
	return d.hand
}
