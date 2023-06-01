package score

import (
	"math"

	"github.com/nherson/psc/api/ent"
)

func Compute(r *ent.FighterResults, endingRound int) float32 {
	score := (float32(r.SignificantStrikesLanded) * float32(0.1)) + float32(r.Takedowns) + (float32(r.Knockdowns) * float32(2)) + (float32(r.ControlTimeSeconds) * 0.01)

	if r.LossByStoppage {
		score = score - 3
	} else if r.WinByStoppage {
		// 12, 10, 8, 8, 8 for round 1, 2, 3, 4, 5 stoppages, respectively
		score = score + float32(math.Max(float64(8), float64(12-2*(endingRound-1))))
	}

	return score
}
