// Copyright 2022 decadevvv

package main

import "context"

type Context struct {
	context.Context
	CurrentRound   uint64
	DealerOpenCard Card
	Shoe           *Shoe
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
