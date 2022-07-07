// Copyright 2022 decadevvv

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
	for supportStrategyName, _ := range StrategyMap {
		SupportStrategyNames = append(SupportStrategyNames, supportStrategyName)
	}
}

var StrategyMap = map[string]Strategy{
	"never-explode": StrategyNeverExplode,
	"basic":         StrategyBasic,
	"dealer":        StrategyDealer,
	"ask-user":      StrategyAskUser,
}
var SupportStrategyNames []string

func main() {
	strategyName := flag.String("s", "basic", fmt.Sprintf("what strategy to use (support %v)", SupportStrategyNames))
	round := flag.Uint64("r", 10000, "number of round to simulate")
	flag.Parse()

	game := NewGame(GameSetting{
		Rules: Rules{
			ShoeSize:                  2,
			ShoePenetrationThreshold:  0.5,
			AllowTriple7BlackJack:     false,
			AllowAceAceSplitAndDouble: true,
		},
		PrintSetting: PrintSetting{
			DelimiterBeforeRound:   false,
			PrintShoeStatus:        false,
			PrintBalanceAfterRound: true,
		},
	})
	strat, ok := StrategyMap[*strategyName]
	if !ok {
		fmt.Printf("Unrecognized strategy name %q: only support %v\n", *strategyName, SupportStrategyNames)
		os.Exit(1)
	}
	game.AddPlayerWithStrategy(strat)
	game.PlayRounds(*round)
	fmt.Println("-----statistics-----")
	fmt.Println(game.Statistics().String())
	fmt.Printf("player edge = %.3f%%\n", game.player.Balance()/float64(*round)*100.0)
}
