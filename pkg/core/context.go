// Copyright 2022 decadevvv

package game

import (
	"context"
	"github.com/decadevvv/blackjack-simulator/pkg/card"
)

type Context struct {
	context.Context
	CurrentRound   uint64
	DealerOpenCard card.Card
	Shoe           *card.Shoe
	GameSetting
}

type GameSetting struct {
	Rules
	PrintSetting
}

type Rules struct {
	ShoeSize                  uint
	ShoePenetrationThreshold  float64
	AllowTriple7BlackJack     bool
	AllowAceAceSplitAndDouble bool
}

type PrintSetting struct {
	DelimiterBeforeRound   bool
	PrintShoeStatus        bool
	PrintBalanceAfterRound bool
}
