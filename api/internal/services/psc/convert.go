package psc

import (
	"math"
	"strconv"

	"github.com/nherson/psc/api/ent"
	apiv1 "github.com/nherson/psc/api/proto/api/v1"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func dbEventToApi(event *ent.Event) *apiv1.Event {
	return &apiv1.Event{
		Id:   strconv.Itoa(event.ID),
		Name: event.Name,
		Date: timestamppb.New(event.Date),
	}
}

// Assumes eager loading of related FighterResults and FighterResults.Fighter edges
// and also there are exactly two Fight.Edges.FighterResults
func dbFightResultsToApi(fightResults *ent.Fight, event *ent.Event) *apiv1.FightResult {
	f0 := fightResults.Edges.FighterResults[0]
	f1 := fightResults.Edges.FighterResults[1]
	return &apiv1.FightResult{
		Event: dbEventToApi(event),
		FighterResults: []*apiv1.FighterResult{
			{
				Fighter: &apiv1.Fighter{
					FirstName: f0.Edges.Fighter.FirstName,
					LastName:  f0.Edges.Fighter.LastName,
					NickName:  f0.Edges.Fighter.NickName,
				},
				SignificantStrikes: int32(f0.SignificantStrikesLanded),
				Takedowns:          int32(f0.Takedowns),
				Knockdowns:         int32(f0.Knockdowns),
				ControlTimeSeconds: int32(f0.ControlTimeSeconds),
				Score:              computeScore(f0, fightResults.ResultEndingRound),
			},
			{
				Fighter: &apiv1.Fighter{
					FirstName: f1.Edges.Fighter.FirstName,
					LastName:  f1.Edges.Fighter.LastName,
					NickName:  f1.Edges.Fighter.NickName,
				},
				SignificantStrikes: int32(f1.SignificantStrikesLanded),
				Takedowns:          int32(f1.Takedowns),
				Knockdowns:         int32(f1.Knockdowns),
				ControlTimeSeconds: int32(f1.ControlTimeSeconds),
				Score:              computeScore(f1, fightResults.ResultEndingRound),
			},
		},
	}
}

func computeScore(r *ent.FighterResults, endingRound int) float32 {
	score := (float32(r.SignificantStrikesLanded) * float32(0.1)) + float32(r.Takedowns) + (float32(r.Knockdowns) * float32(2)) + (float32(r.ControlTimeSeconds) * 0.01)

	if r.LossByStoppage {
		score = score - 3
	} else if r.WinByStoppage {
		// 12, 10, 8, 8, 8 for round 1, 2, 3, 4, 5 stoppages, respectively
		score = score + float32(math.Max(float64(8), float64(12-2*(endingRound-1))))
	}

	return score
}
