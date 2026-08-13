package main

type HandEval int

const (
	NoResult HandEval = iota
	DealerWin
	PlayerWin
)

func EvaluateDrawResults(p Hand, d Hand) HandEval {

	pSum := GetHandValue(p)
	dSum := GetHandValue(d)

	switch {
	case pSum == 21:
	case dSum > 21:
	case dSum >= 17 && pSum > dSum:
		{
			return PlayerWin
		}
	case dSum == 21:
	case pSum > 21:
	case dSum >= 17 && pSum < dSum:
		{
			return DealerWin
		}
	}

	return NoResult
}
