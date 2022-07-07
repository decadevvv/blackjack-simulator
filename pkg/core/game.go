// Copyright 2022 decadevvv

package game

import (
	"context"
	"fmt"
	"github.com/decadevvv/blackjack-simulator/pkg/card"
	"github.com/decadevvv/blackjack-simulator/pkg/core"
	"github.com/decadevvv/blackjack-simulator/pkg/strategy"
)

type Game struct {
	ctx        *core.Context
	dealer     *Dealer
	Player     *Player
	statistics Statistics
}

type Statistics struct {
	Hands                int
	CanSplitHands        int
	SplitHands           int
	DoubleHands          int
	BlackJacksHands      int
	BlackJackPushHands   int
	WinHands             int
	DoubleWinHands       int
	PushHands            int
	DoublePushHands      int
	LoseHands            int
	DoubleLoseHands      int
	SplitWinHands        int
	SplitPushHands       int
	SplitLoseHands       int
	SplitDoubleWinHands  int
	SplitDoublePushHands int
	SplitDoubleLoseHands int
	IllegalResults       int
}

func (s Statistics) String() string {
	rounds := s.Hands - s.SplitHands/2
	msg := fmt.Sprintf("%d rounds, %d hands, %d win, %d push, %d lose",
		rounds, s.Hands, s.WinHands, s.PushHands, s.LoseHands)
	blackJackProbability := float64(s.BlackJacksHands) / float64(s.Hands)
	blackJackWinHands := s.BlackJacksHands - s.BlackJackPushHands
	blackJackMargin := float64(blackJackWinHands) * 1.5
	blackJackEdge := blackJackMargin / float64(s.BlackJacksHands)
	blackJackContribution := blackJackMargin / float64(rounds)
	msg += "\n" + fmt.Sprintf("BlackJack analysis: %d blackjacks (%.2f%%), %d win, %d push, margin %.1f, edge %.2f%%, contribution %.2f%%",
		s.BlackJacksHands, blackJackProbability*100, blackJackWinHands, s.BlackJackPushHands, blackJackMargin, blackJackEdge*100, blackJackContribution*100)
	hitStandHands := s.Hands - s.BlackJacksHands - s.DoubleHands
	hitStandProbability := float64(hitStandHands) / float64(s.Hands)
	hitStandWinHands := s.WinHands - s.DoubleWinHands - (s.BlackJacksHands - s.BlackJackPushHands)
	hitStandPushHands := s.PushHands - s.DoublePushHands - s.BlackJackPushHands
	hitStandLoseHands := s.LoseHands - s.DoubleLoseHands
	hitStandWinRate := float64(hitStandWinHands+hitStandPushHands/2) / float64(hitStandHands)
	hitStandMargin := float64(hitStandWinHands - hitStandLoseHands)
	hitStandEdge := hitStandMargin / float64(hitStandHands)
	hitStandContribution := hitStandMargin / float64(rounds)
	msg += "\n" + fmt.Sprintf("Hit stand analysis: %d hit/stand hands (%2.f%%), %d win, %d push, %d lose, %.2f%% win rate, margin %.0f, edge %.2f%%, contribution %.2f%%",
		hitStandHands, hitStandProbability*100, hitStandWinHands, hitStandPushHands, hitStandLoseHands, hitStandWinRate*100, hitStandMargin, hitStandEdge*100, hitStandContribution*100)
	hsbHands := hitStandHands + s.BlackJacksHands
	hsbProbability := hitStandProbability + blackJackProbability
	hsbWinHands := hitStandWinHands + blackJackWinHands
	hsbPushHands := hitStandPushHands + s.BlackJackPushHands
	hsbLoseHands := hitStandLoseHands
	hsbWinRate := float64(hsbWinHands+hsbPushHands/2) / float64(hsbHands)
	hsbMargin := hitStandMargin + blackJackMargin
	hsbEdge := hsbMargin / float64(hsbHands)
	hsbContribution := hsbMargin / float64(rounds)
	msg += "\n" + fmt.Sprintf("Hit/Stand/BlackJack analysis: %d hit/stand/blackjack hands (%2.f%%), %d win, %d push, %d lose, %.2f%% win rate, margin %.1f, edge %.2f%%, contribition %.2f%%",
		hsbHands, hsbProbability*100, hsbWinHands, hsbPushHands, hsbLoseHands,
		hsbWinRate*100, hsbMargin, hsbEdge*100, hsbContribution*100)
	doubleProbability := float64(s.DoubleHands) / float64(s.Hands)
	doubleWinRate := float64(s.DoubleWinHands+s.DoublePushHands/2) / float64(s.DoubleHands)
	doubleMargin := float64(s.DoubleWinHands-s.DoubleLoseHands) * 2
	doubleEdge := doubleMargin / float64(s.DoubleHands)
	doubleEdgeContribution := doubleMargin / float64(rounds)
	msg += "\n" + fmt.Sprintf("Double hand analysis: %d doubles hands (%.2f%%), %d win, %d push, %d lose, %.2f%% win rate, margin %.0f, edge %.2f%%, contribution %.2f%%",
		s.DoubleHands, doubleProbability*100, s.DoubleWinHands, s.DoublePushHands, s.DoubleLoseHands, doubleWinRate*100, doubleMargin, doubleEdge*100, doubleEdgeContribution*100)
	splits := s.SplitHands / 2
	canSplitProbability := float64(s.CanSplitHands) / float64(rounds)
	splitProbability := float64(splits) / float64(s.CanSplitHands)
	splitWinRate := float64(s.SplitWinHands+s.SplitPushHands/2) / float64(s.SplitHands)
	splitMargin := s.SplitWinHands + 2*s.SplitDoubleWinHands - s.SplitLoseHands - 2*s.SplitDoubleLoseHands
	splitEdge := float64(splitMargin) / float64(s.SplitHands)
	splitContribution := float64(splitMargin) / float64(rounds)
	msg += "\n" + fmt.Sprintf("Split analysis: %d hands could split (%.2f%%), %d splits (%.2f%%), %d win, %d push, %d lose, %.2f%% win rate, margin %d, edge %.2f%%, contribution %.2f%%",
		s.CanSplitHands, canSplitProbability*100, splits, splitProbability*100, s.SplitWinHands, s.SplitPushHands, s.SplitLoseHands, splitWinRate*100, splitMargin, splitEdge*100, splitContribution*100)
	return msg
}

func NewGame(setting core.GameSetting) *Game {
	return &Game{
		ctx: &core.Context{
			Context:     context.Background(),
			GameSetting: setting,
		},
		dealer:     NewDealer(),
		Player:     nil,
		statistics: Statistics{},
	}
}

func (g *Game) Statistics() Statistics {
	return g.statistics
}

func (g *Game) AddPlayerWithStrategy(s strategy.Strategy) {
	g.Player = NewPlayer(s)
}

func (g *Game) PlayRounds(round uint64) {
	g.ctx.Shoe = card.NewShoe(g.ctx.Rules.ShoeSize, g.ctx.Rules.ShoePenetrationThreshold)
	g.ctx.Shoe.Reshuffle()
	for i := uint64(1); i <= round; i++ {
		g.ctx.CurrentRound = i
		g.Play()
	}
}

func (g *Game) Play() {
	if g.ctx.PrintSetting.DelimiterBeforeRound {
		fmt.Printf("-----Round %d-----\n", g.ctx.CurrentRound)
	}

	g.dealer.InitRound(g.ctx)
	g.Player.InitRound(g.ctx)

	g.Player.PlayRound(g.ctx)
	g.dealer.PlayRound(g.ctx)

	playerHand := g.Player.Hand()
	playerSplitHand := g.Player.SplitHand()
	dealerHand := g.dealer.Hand()

	var msg string
	var delta float64
	beforeBalance := g.Player.Balance()
	if playerSplitHand == nil {
		roundResult := g.HandResult(playerHand)
		delta = HandRatio[roundResult]
		g.Player.BalanceChange(delta)
		g.UpdateStatistics(playerHand, roundResult)
		msg = fmt.Sprintf("round %d %s: %v V.S. %v", g.ctx.CurrentRound, roundResult, playerHand, dealerHand)
	} else {
		handResult := g.HandResult(playerHand)
		delta = HandRatio[handResult]
		g.UpdateStatistics(playerHand, handResult)
		splitHandResult := g.HandResult(playerSplitHand)
		delta += HandRatio[handResult]
		g.UpdateStatistics(playerSplitHand, splitHandResult)
		g.Player.BalanceChange(delta)
		delta = HandRatio[handResult] + HandRatio[splitHandResult]
		msg = fmt.Sprintf("round %d %s %s: %v %v V.S. %v", g.ctx.CurrentRound, handResult, splitHandResult, playerHand, playerSplitHand, dealerHand)
	}
	if g.ctx.PrintShoeStatus {
		msg = fmt.Sprintf("%s (shoe/stage/garbage=%d/%d/%d)", msg, g.ctx.Shoe.InShoeLen(), g.ctx.Shoe.InStageLen(), g.ctx.Shoe.InGarbageLen())
	}
	if g.ctx.PrintBalanceAfterRound {
		if delta > 0 {
			msg = fmt.Sprintf("%s (player balance %.1f + %.1f => %.1f)", msg, beforeBalance, delta, g.Player.Balance())
		} else if delta == 0 {
			msg = fmt.Sprintf("%s (player balance %.1f unchanged)", msg, beforeBalance)
		} else {
			msg = fmt.Sprintf("%s (player balance %.1f - %.1f => %.1f)", msg, beforeBalance, -delta, g.Player.Balance())
		}

	}
	fmt.Println(msg)

	g.ctx.Shoe.StageToGarbage()
}

func (g *Game) UpdateStatistics(hand *card.Hand, handResult HandResult) {
	g.statistics.Hands++
	switch handResult {
	case HandBlackJack:
		g.statistics.WinHands++
		g.statistics.BlackJacksHands++
	case HandBlackJackPush:
		g.statistics.PushHands++
		g.statistics.BlackJacksHands++
		g.statistics.BlackJackPushHands++
	case HandWin:
		g.statistics.WinHands++
		if hand.IsSplit() {
			g.statistics.SplitWinHands++
		}
	case HandDoubleWin:
		g.statistics.WinHands++
		g.statistics.DoubleWinHands++
		if hand.IsSplit() {
			g.statistics.SplitWinHands++
			g.statistics.SplitDoubleWinHands++
		}
	case HandPush:
		g.statistics.PushHands++
		if hand.IsSplit() {
			g.statistics.SplitPushHands++
		}
	case HandDoublePush:
		g.statistics.PushHands++
		g.statistics.DoublePushHands++
		if hand.IsSplit() {
			g.statistics.SplitPushHands++
			g.statistics.SplitDoublePushHands++
		}
	case HandLose:
		g.statistics.LoseHands++
		if hand.IsSplit() {
			g.statistics.SplitLoseHands++
		}
	case HandDoubleLose:
		g.statistics.LoseHands++
		g.statistics.DoubleLoseHands++
		if hand.IsSplit() {
			g.statistics.SplitLoseHands++
			g.statistics.SplitDoubleLoseHands++
		}
	}
	if hand.IsSplit() {
		g.statistics.SplitHands++
	}
	if hand.Doubled() {
		g.statistics.DoubleHands++
	}
	if hand.CouldSplit() {
		g.statistics.CanSplitHands++
	}
}

func (g *Game) HandResult(hand *card.Hand) (res HandResult) {
	defer func() {
		if hand.Doubled() {
			switch res {
			case HandWin:
				res = HandDoubleWin
			case HandPush:
				res = HandDoublePush
			case HandLose:
				res = HandDoubleLose
			default:
				panic(fmt.Errorf("illegal doubled result: %s", res))
			}
		}

	}()
	if hand.Busted() {
		return HandLose
	} else if hand.BlackJack() {
		if g.dealer.Hand().BlackJack() {
			return HandBlackJackPush
		} else {
			return HandBlackJack
		}
	} else if g.dealer.Hand().Busted() {
		return HandWin
	} else if g.dealer.Hand().BlackJack() {
		return HandLose
	} else if hand.Points() > g.dealer.Hand().Points() {
		return HandWin
	} else if hand.Points() < g.dealer.Hand().Points() {
		return HandLose
	} else {
		return HandPush
	}
}

type HandResult string

const (
	HandBlackJack     HandResult = "blackjack"
	HandBlackJackPush HandResult = "blackjack push"
	HandWin           HandResult = "win"
	HandDoubleWin     HandResult = "double win"
	HandPush          HandResult = "push"
	HandDoublePush    HandResult = "double push"
	HandLose          HandResult = "lose"
	HandDoubleLose    HandResult = "double lose"
)

var HandRatio = map[HandResult]float64{
	HandBlackJack:     1.5,
	HandBlackJackPush: 0,
	HandWin:           1,
	HandDoubleWin:     2,
	HandPush:          0,
	HandDoublePush:    0,
	HandLose:          -1,
	HandDoubleLose:    -2,
}
