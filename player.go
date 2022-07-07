// Copyright 2022 decadevvv

package main

type Player struct {
	hand        *Hand
	splitHand   *Hand
	strategy    Strategy
	splitPlayed bool
	balance     float64
}

func NewPlayer(strategy Strategy) *Player {
	return &Player{
		hand:        nil,
		splitHand:   nil,
		strategy:    strategy,
		splitPlayed: false,
		balance:     0,
	}
}

func (p *Player) BalanceChange(delta float64) {
	p.balance += delta
}

func (p *Player) Balance() float64 {
	return p.balance
}

func (p *Player) Hand() *Hand {
	return p.hand
}

func (p *Player) SplitHand() *Hand {
	return p.splitHand
}

func (p *Player) InitRound(ctx *Context) {
	p.hand = NewHand(false, ctx.Rules)
	p.splitHand = nil
	p.hand.Add(ctx.Shoe.Deal())
	p.hand.Add(ctx.Shoe.Deal())
	p.splitPlayed = false
}

func (p *Player) PlayRound(ctx *Context) {
	defer func() {
		if p.splitHand != nil && !p.splitPlayed {
			p.PlayRound(ctx)
		}
	}()
	var hand *Hand
	if p.splitHand != nil {
		hand = p.splitHand
		defer func() {
			p.splitPlayed = true
		}()
	} else {
		hand = p.hand
	}
	for (!hand.Busted()) && (!hand.BlackJack()) && (hand.Points() <= 21) {
		if hand.CanSplit() && p.strategy.Split(ctx, hand.Cards()[0]) {
			p.hand, p.splitHand = hand.Split()
			hand = p.hand
			hand.Add(ctx.Shoe.Deal())
			p.splitHand.Add(ctx.Shoe.Deal())
		}
		if hand.CanDouble() && p.strategy.Double(ctx, hand.Cards()[0], hand.Cards()[1]) {
			hand.Double()
			hand.Add(ctx.Shoe.Deal())
			break
		}
		if p.strategy.Hit(ctx, hand) {
			hand.Add(ctx.Shoe.Deal())
		} else {
			break
		}
	}
}
