// Copyright 2022 decadevvv

package main

type Dealer struct {
	hand *Hand
}

func NewDealer() *Dealer {
	return &Dealer{
		hand: nil,
	}
}

func (d *Dealer) InitRound(ctx *Context) {
	d.hand = NewHand(false, ctx.Rules)
	openCard := ctx.Shoe.Deal()
	d.hand.Add(openCard)
	d.hand.Add(ctx.Shoe.Deal())
	ctx.DealerOpenCard = openCard
}

func (d *Dealer) PlayRound(ctx *Context) {
	for d.hand.Points() < 17 {
		d.hand.Add(ctx.Shoe.Deal())
	}
}

func (d *Dealer) Hand() *Hand {
	return d.hand
}
