package data

import (
	"context"
	"fmt"
	"math"
	"net/http"
	"sort"
	"strings"

	"github.com/gocarina/gocsv"
	"github.com/pkg/errors"

	"github.com/nherson/psc/api/ent"
	"github.com/nherson/psc/api/ent/fighter"
	"github.com/nherson/psc/api/internal/score"
)

type Service struct {
	DB *ent.Client
}

// Name, sig strikes, takedowns, knockdowns, control time, odds, score, date
type Entry struct {
	Name               string  `csv:"name"`
	SigStrikes         int     `csv:"sig_strikes"`
	Takedowns          int     `csv:"takedowns"`
	Knockdowns         int     `csv:"knockdowns"`
	ControlTimeSeconds int     `csv:"control_time_seconds"`
	Odds               int     `csv:"odds"`
	Score              float32 `csv:"score"`
	debut              bool    `csv:"-"`
}

func (s *Service) PublicCSV(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	entries, err := s.fetchEntries(ctx)
	if err != nil {
		writeError(w, err)
		return
	}

	data, err := gocsv.MarshalString(&entries)
	if err != nil {
		writeError(w, err)
		return
	}

	_, err = w.Write([]byte(data))
	if err != nil {
		writeError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/csv")

}

type NickEntry struct {
	Name       string  `csv:"last_name"`
	Odds       int     `csv:"odds"`
	Score      float32 `csv:"score"`
	FightCount int     `csv:"fight_count"`
}

type ByFightCount []*NickEntry

func (a ByFightCount) Len() int           { return len(a) }
func (a ByFightCount) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByFightCount) Less(i, j int) bool { return a[i].FightCount < a[j].FightCount }

func (s *Service) NickCSV(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	entries, err := s.fetchEntries(ctx)
	if err != nil {
		writeError(w, err)
		return
	}

	m := make(map[string][]*Entry)
	for _, entry := range entries {
		m[entry.Name] = append(m[entry.Name], entry)
	}

	var nickEntries []*NickEntry
	for name, entries := range m {
		ne := &NickEntry{
			Name:       name,
			Odds:       entries[0].Odds,
			Score:      avgScore(entries),
			FightCount: len(entries),
		}
		if entries[0].debut {
			ne.FightCount = 0
		}
		nickEntries = append(nickEntries, ne)
	}

	sort.Sort(ByFightCount(nickEntries))

	data, err := gocsv.MarshalString(&nickEntries)
	if err != nil {
		writeError(w, err)
		return
	}

	_, err = w.Write([]byte(data))
	if err != nil {
		writeError(w, err)
		return
	}

	w.Header().Set("Content-Type", "text/csv")
}

func (s *Service) fetchEntries(ctx context.Context) ([]*Entry, error) {
	fighters, err := s.DB.Fighter.Query().
		Where(
			fighter.HasUpcomingFights(),
		).
		WithUpcomingFights().
		WithUpcomingFighterOdds().
		WithFights(
			func(q *ent.FightQuery) {
				q.WithFighterResults()
			},
		).
		All(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "could not fetch upcoming fighters")
	}

	var entries []*Entry
	for _, f := range fighters {
		fullName := strings.TrimSpace(fmt.Sprintf("%s %s", f.FirstName, f.LastName))
		if len(f.Edges.UpcomingFighterOdds) == 0 {
			return nil, errors.New("no m2m record for upcoming fighter")
		}

		oddsPtr := f.Edges.UpcomingFighterOdds[0].BestOdds
		var odds int
		if oddsPtr != nil {
			odds = *oddsPtr
		}

		for _, fight := range f.Edges.Fights {
			endingRound := fight.ResultEndingRound
			for _, fr := range fight.Edges.FighterResults {
				if fr.FighterID != f.ID {
					continue
				}

				entries = append(entries, &Entry{
					Name:               fullName,
					SigStrikes:         fr.SignificantStrikesLanded,
					Takedowns:          fr.Takedowns,
					Knockdowns:         fr.Knockdowns,
					ControlTimeSeconds: fr.ControlTimeSeconds,
					Odds:               odds,
					Score:              score.Compute(fr, endingRound),
				})
			}
		}
		// Make a dummy entry if this is a new fighter
		if len(f.Edges.Fights) == 0 {
			entries = append(entries, &Entry{
				Name:               fullName,
				SigStrikes:         0,
				Takedowns:          0,
				Knockdowns:         0,
				ControlTimeSeconds: 0,
				Odds:               odds,
				Score:              0,
				debut:              true,
			})
		}

	}

	return entries, nil
}

func writeError(w http.ResponseWriter, err error) {
	w.WriteHeader(500)
	w.Write([]byte(err.Error()))
}

func avgScore(entries []*Entry) float32 {
	s := float32(0)
	for _, e := range entries {
		s += e.Score
	}
	avg := s / float32(len(entries))
	return float32(math.Round(float64(avg)*100) / 100)
}
