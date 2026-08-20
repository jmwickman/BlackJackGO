package main

import "fmt"

type HandEval int

const (
	NoResult HandEval = iota
	DealerWin
	PlayerWin
)

func EvaluateDrawResults(p Hand, d Hand, b float32) (HandEval, float32) {

	pSum := GetHandValue(p)
	dSum := GetHandValue(d)

	fmt.Println(pSum, dSum)

	switch {
	case pSum == 21, dSum > 21, dSum >= 17 && pSum > dSum:
		{
			return PlayerWin, GetWinAmount(b, pSum == 21)
		}
	case dSum == 21, pSum > 21, dSum >= 17 && pSum < dSum:
		{
			return DealerWin, 0
		}
	}

	return NoResult, 0
}

func GetWinAmount(b float32, isBJ bool) float32 {

	// A blackjack pays 3:2
	if isBJ {
		return b + b*1.5
	}

	// Base pay is 1:1
	return b * 2
}
