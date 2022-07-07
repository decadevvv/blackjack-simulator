package strategy

import (
	"fmt"
	"github.com/decadevvv/blackjack-simulator/pkg/core"
)

var StrategyBasic = NewStrategy(SplitStrategyBasic, DoubleStrategyBasic, HitStrategyBasic)
var StrategyBasicOnlyHit = NewStrategy(SplitStrategyNeverSplit, DoubleStrategyNeverDouble, HitStrategyBasic)
var StrategyNeverExplode = NewStrategy(SplitStrategyNeverSplit, DoubleStrategyNeverDouble, HitStrategyNeverExplode)
var StrategyDealer = NewStrategy(SplitStrategyNeverSplit, DoubleStrategyNeverDouble, HitStrategyDealer)
var StrategyAskUser = NewStrategy(SplitStrategyAskUser, DoubleStrategyAskUser, HitStrategyAskUser)

var HitStrategyNeverExplode = HitStrategySoftHardLimit(17, 12)

var HitStrategyDealer = HitStrategySoftHardLimit(17, 17)

var SplitStrategyBasic SplitStrategy = func(ctx *core.Context, card core.Card) bool {
	return BasicSplitStrategyTable[card][ctx.DealerOpenCard]
}

var DoubleStrategyBasic DoubleStrategy = func(ctx *core.Context, card1 core.Card, card2 core.Card) bool {
	//if card1 == CardA && card2 == CardA {
	//	return true
	//}
	if card1 == core.CardA {
		return BasicSoftDoubleStrategyTable[card2][ctx.DealerOpenCard]
	}
	if card2 == core.CardA {
		return BasicSoftDoubleStrategyTable[card1][ctx.DealerOpenCard]
	}
	hardPoint := core.CardPoint[card1] + core.CardPoint[card2]
	if hardPoint == 8 && (ctx.DealerOpenCard == core.Card5 || ctx.DealerOpenCard == core.Card6) {
		return !(card1 == core.Card6 || card2 == core.Card6)
	}
	return BasicHardDoubleStrategyTable[hardPoint][ctx.DealerOpenCard]
}

var HitStrategyBasic HitStrategy = func(ctx *core.Context, hand *core.Hand) bool {
	if hand.Soft() {
		return BasicSoftHitStrategyTable[hand.Points()][ctx.DealerOpenCard]
	} else {
		if hand.Points() == 16 && ctx.DealerOpenCard.Point() == 10 {
			return hand.Len() == 2
		}
		if hand.Points() == 14 && hand.Cards()[0] == 7 && hand.Cards()[1] == 7 && ctx.DealerOpenCard.Point() == 10 {
			return false
		}
		return BasicHardHitStrategyTable[hand.Points()][ctx.DealerOpenCard]
	}
}

const (
	T bool = true
	F bool = false
	X bool = false
)

var BasicSplitStrategyTable = [14][14]bool{
	/* dealer X, A, 2, 3, 4, 5, 6, 7, 8, 9, X, J, Q, K */
	/* my card                                         */
	/* X  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, /* not exist placeholder card */
	/* A  */ {X, T, T, T, T, T, T, T, T, T, T, T, T, T}, /* always split A */
	/* 2  */ {X, F, T, T, T, T, T, T, F, F, F, F, F, F},
	/* 3  */ {X, F, T, T, T, T, T, T, F, F, F, F, F, F},
	/* 4  */ {X, F, F, F, F, T, F, F, F, F, F, F, F, F}, /* only split 4 on dealer 5 */
	/* 5  */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, /* never split 5  */
	/* 6  */ {X, F, T, T, T, T, T, T, F, F, F, F, F, F},
	/* 7  */ {X, F, T, T, T, T, T, T, T, F, F, F, F, F},
	/* 8  */ {X, T, T, T, T, T, T, T, T, T, T, T, T, T}, /* always split 8 */
	/* 9  */ {X, F, T, T, T, T, T, F, T, T, F, F, F, F},
	/* 10 */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, /* never split 10 */
	/* J  */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, /* never split 10 */
	/* Q  */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, /* never split 10 */
	/* K  */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, /* never split 10 */
}

var BasicSoftDoubleStrategyTable = [14][14]bool{
	/*   never soft double on dealer A, 7, 8, 9, 10 */
	/* dealer X, A, 2, 3, 4, 5, 6, 7, 8, 9, X, J, Q, K */
	/* my card                                      */
	/* X  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, /* not exist placeholder card */
	/* A  */ {X, F, F, F, F, T, T, F, F, F, F, F, F, F},
	/* 2  */ {X, F, F, F, T, T, T, F, F, F, F, F, F, F},
	/* 3  */ {X, F, F, F, T, T, T, F, F, F, F, F, F, F},
	/* 4  */ {X, F, F, F, T, T, T, F, F, F, F, F, F, F},
	/* 5  */ {X, F, F, F, T, T, T, F, F, F, F, F, F, F},
	/* 6  */ {X, F, T, T, T, T, T, F, F, F, F, F, F, F},
	/* 7  */ {X, F, F, T, T, T, T, F, F, F, F, F, F, F},
	/* 8  */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never double, double will highly possible cause point down and lose
	/* 9  */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never double, double will highly possible cause point down and lose
	/* 10 */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // blackjack, cannot double
	/* J  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // blackjack, cannot double
	/* Q  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // blackjack, cannot double
	/* K  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // blackjack, cannot double
}

var BasicHardDoubleStrategyTable = [22][14]bool{
	/* dealer X, A, 2, 3, 4, 5, 6, 7, 8, 9, X, J, Q, K */
	/* my hard point                                */
	/* 0  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // invalid points
	/* 1  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // invalid points
	/* 2  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // invalid points (no hard 2, A + A = 12 and soft)
	/* 3  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // invalid points (no hard 3, A + 2 = 13 and soft)
	/* 4  */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never double on hard point < 9
	/* 5  */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never double on hard point < 9
	/* 6  */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never double on hard point < 9
	/* 7  */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never double on hard point < 9
	/* 8  */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never double on hard point < 9
	/* 9  */ {X, F, T, T, T, T, T, F, F, F, F, F, F, F}, // hard 9 can double on dealer 2~6
	/* 10 */ {X, F, T, T, T, T, T, T, T, T, F, F, F, F}, // hard 10 can double on dealer 2~9
	/* 11 */ {X, T, T, T, T, T, T, T, T, T, T, T, T, T}, // always double on hard 11 (never bust)
	/* 12 */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never double on hard point > 11 (easy to bust)
	/* 13 */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never double on hard point > 11 (easy to bust)
	/* 14 */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never double on hard point > 11 (easy to bust)
	/* 15 */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never double on hard point > 11 (easy to bust)
	/* 16 */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never double on hard point > 11 (easy to bust)
	/* 17 */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never double on hard point > 11 (easy to bust)
	/* 18 */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never double on hard point > 11 (easy to bust)
	/* 19 */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never double on hard point > 11 (easy to bust)
	/* 20 */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never double on hard point > 11 (easy to bust)
	/* 21 */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // invalid point, no 2 card hard blackjack
}

var BasicSoftHitStrategyTable = [22][14]bool{
	/* dealer X, A, 2, 3, 4, 5, 6, 7, 8, 9, X, J, Q, K */
	/* my soft point                                */
	/* 0  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // invalid points
	/* 1  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // invalid points
	/* 2  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // invalid points (soft point >= 12)
	/* 3  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // invalid points (soft point >= 12)
	/* 4  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // invalid points (soft point >= 12)
	/* 5  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // invalid points (soft point >= 12)
	/* 6  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // invalid points (soft point >= 12)
	/* 7  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // invalid points (soft point >= 12)
	/* 8  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // invalid points (soft point >= 12)
	/* 9  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // invalid points (soft point >= 12)
	/* 10 */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // invalid points (soft point >= 12)
	/* 11 */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // invalid points (soft point >= 12)
	/* 12 */ {X, T, T, T, T, T, T, T, T, T, T, T, T, T}, // always hit on soft point <= 17
	/* 13 */ {X, T, T, T, T, T, T, T, T, T, T, T, T, T}, // always hit on soft point <= 17
	/* 14 */ {X, T, T, T, T, T, T, T, T, T, T, T, T, T}, // always hit on soft point <= 17
	/* 15 */ {X, T, T, T, T, T, T, T, T, T, T, T, T, T}, // always hit on soft point <= 17
	/* 16 */ {X, T, T, T, T, T, T, T, T, T, T, T, T, T}, // always hit on soft point <= 17
	/* 17 */ {X, T, T, T, T, T, T, T, T, T, T, T, T, T}, // always hit on soft point <= 17
	/* 18 */ {X, F, F, F, F, F, F, F, F, T, T, T, T, T}, // only hit soft 18 on dealer 9, 10
	/* 19 */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never hit on soft point > 18
	/* 20 */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never hit on soft point > 18
	/* 21 */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // blackjack, cannot hit
}

var BasicHardHitStrategyTable = [22][14]bool{
	/* dealer X, A, 2, 3, 4, 5, 6, 7, 8, 9, X, J, Q, K */
	/* my hard point                                */
	/* 0  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // invalid points
	/* 1  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // invalid points
	/* 2  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // invalid points (no hard 2, A + A = 12 and soft)
	/* 3  */ {X, X, X, X, X, X, X, X, X, X, X, X, X, X}, // invalid points (no hard 3, A + 2 = 13 and soft)
	/* 4  */ {X, T, T, T, T, T, T, T, T, T, T, T, T, T}, // always hit on hard point < 12
	/* 5  */ {X, T, T, T, T, T, T, T, T, T, T, T, T, T}, // always hit on hard point < 12
	/* 6  */ {X, T, T, T, T, T, T, T, T, T, T, T, T, T}, // always hit on hard point < 12
	/* 7  */ {X, T, T, T, T, T, T, T, T, T, T, T, T, T}, // always hit on hard point < 12
	/* 8  */ {X, T, T, T, T, T, T, T, T, T, T, T, T, T}, // always hit on hard point < 12
	/* 9  */ {X, T, T, T, T, T, T, T, T, T, T, T, T, T}, // always hit on hard point < 12
	/* 10 */ {X, T, T, T, T, T, T, T, T, T, T, T, T, T}, // always hit on hard point < 12
	/* 11 */ {X, T, T, T, T, T, T, T, T, T, T, T, T, T}, // always hit on hard point < 12
	/* 12 */ {X, T, T, T, F, F, F, T, T, T, T, T, T, T}, // only hit hard 12 on dealer 2, 3
	/* 13 */ {X, T, F, F, F, F, F, T, T, T, T, T, T, T}, // only hit hard 13~16 on 7 8 9 10 A
	/* 14 */ {X, T, F, F, F, F, F, T, T, T, T, T, T, T}, // only hit hard 13~16 on 7 8 9 10 A
	/* 15 */ {X, T, F, F, F, F, F, T, T, T, T, T, T, T}, // only hit hard 13~16 on 7 8 9 10 A
	/* 16 */ {X, T, F, F, F, F, F, T, T, T, T, T, T, T}, // only hit hard 13~16 on 7 8 9 10 A
	/* 17 */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never hit on hard point > 16
	/* 18 */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never hit on hard point > 16
	/* 19 */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never hit on hard point > 16
	/* 20 */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never hit on hard point > 16
	/* 21 */ {X, F, F, F, F, F, F, F, F, F, F, F, F, F}, // never hit on hard point > 16
}

var SplitStrategyAskUser SplitStrategy = func(ctx *core.Context, card core.Card) bool {
	return AskUser(fmt.Sprintf("[%s %s] V.S. [%s ?], wanna split? ", card.String(), card.String(), ctx.DealerOpenCard.String()))
}

var DoubleStrategyAskUser DoubleStrategy = func(ctx *core.Context, card1 core.Card, card2 core.Card) bool {
	return AskUser(fmt.Sprintf("[%s %s] (%d) V.S. [%s ?], wanna double? ", card1.String(), card2.String(), card1.Point()+card2.Point(), ctx.DealerOpenCard.String()))
}

var HitStrategyAskUser HitStrategy = func(ctx *core.Context, hand *core.Hand) bool {
	return AskUser(fmt.Sprintf("%v V.S. [%s ?], wanna hit? ", hand, ctx.DealerOpenCard.String()))
}

func AskUser(msg string) bool {
	var answer string
	firstTime := true
	fmt.Print(msg)
	for answer != "Y" && answer != "N" {
		if firstTime {
			firstTime = false
		} else {
			fmt.Printf("must be 'Y' or 'N', got %q please input again\n", answer)
		}
		fmt.Scanf("%s\n", &answer)
	}
	return answer == "Y"
}
