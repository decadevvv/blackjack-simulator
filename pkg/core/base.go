package strategy

import (
	"github.com/decadevvv/blackjack-simulator/pkg/core"
)

type Strategy interface {
	Split(ctx *core.Context, card core.Card) bool
	Double(ctx *core.Context, card1 core.Card, card2 core.Card) bool
	Hit(ctx *core.Context, hand *core.Hand) bool
}

type SplitStrategy func(ctx *core.Context, card core.Card) bool
type DoubleStrategy func(ctx *core.Context, card1 core.Card, card2 core.Card) bool
type HitStrategy func(ctx *core.Context, hand *core.Hand) bool

type strategy struct {
	split  SplitStrategy
	double DoubleStrategy
	hit    HitStrategy
}

func NewStrategy(split SplitStrategy, double DoubleStrategy, hit HitStrategy) Strategy {
	return &strategy{
		split:  split,
		double: double,
		hit:    hit,
	}
}

func (s *strategy) Split(ctx *core.Context, card core.Card) bool {
	return s.split(ctx, card)
}

func (s *strategy) Double(ctx *core.Context, card1 core.Card, card2 core.Card) bool {
	return s.double(ctx, card1, card2)
}

func (s *strategy) Hit(ctx *core.Context, hand *core.Hand) bool {
	return s.hit(ctx, hand)
}

var SplitStrategyNeverSplit SplitStrategy = func(ctx *core.Context, card core.Card) bool {
	return false
}

var DoubleStrategyNeverDouble DoubleStrategy = func(ctx *core.Context, card1 core.Card, card2 core.Card) bool {
	return false
}

var HitStrategyNeverHit HitStrategy = func(ctx *core.Context, hand *core.Hand) bool {
	return false
}

func HitStrategySoftHardLimit(softLimit uint8, hardLimit uint8) HitStrategy {
	return func(ctx *core.Context, hand *core.Hand) bool {
		if hand.Soft() {
			return hand.Points() < softLimit
		} else {
			return hand.Points() < hardLimit
		}
	}
}
