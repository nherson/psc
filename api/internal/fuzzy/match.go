package fuzzy

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/adrg/strutil"
	"github.com/manifoldco/promptui"
	"github.com/nherson/psc/api/ent"
)

// MatchFigher returns the best fighter name match so long as the confidence threshold is met.
// Otherwise, returns an error
func (m *Matcher) MatchFighter(ctx context.Context, fullName string) (*ent.Fighter, float64, error) {
	ss, err := m.fighterSimilarities(ctx, fullName)
	if err != nil {
		return nil, 0, err
	}

	if ss[0].score >= float64(m.confidenceThreshold) {
		return ss[0].dbFighter, ss[0].score, nil
	} else {
		return nil, 0, ErrNoMatch
	}
}

func (m *Matcher) MatchWithPrompt(ctx context.Context, fullName string) (*ent.Fighter, float64, error) {
	ss, err := m.fighterSimilarities(ctx, fullName)
	if err != nil {
		return nil, 0, err
	}

	if ss[0].score >= float64(m.confidenceThreshold) {
		return ss[0].dbFighter, ss[0].score, nil
	}

	var matchOptions []string
	for i := 0; i < m.promptChoices; i++ {
		match := ss[i]
		matchOptions = append(matchOptions, fmt.Sprintf("Name: %s, Score: %.2f", match.fullName, match.score))
	}
	matchOptions = append(matchOptions, "None of the above")

	prompt := promptui.Select{
		Label: "Choose the best DB match",
		Items: matchOptions,
		Size:  m.promptChoices + 1,
	}
	idx, _, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	if idx == 5 {
		return nil, 0, ErrNoMatch
	} else {
		return ss[idx].dbFighter, ss[idx].score, nil
	}
}

func (m *Matcher) fighterSimilarities(ctx context.Context, fullName string) (similarities, error) {
	fullName = strings.ToLower(fullName)

	fighters, err := m.getFighters(ctx)
	if err != nil {
		return nil, err
	}
	if len(fighters) == 0 {
		// this shouldnt happen, but lets not risk panicking later
		return nil, errors.New("no fighters")
	}

	var out similarities
	for _, f := range fighters {
		candidateFullName := fmt.Sprintf("%s %s", strings.ToLower(f.FirstName), strings.ToLower(f.LastName))

		score := strutil.Similarity(fullName, candidateFullName, m.stringMetric)
		out = append(out, similarityOutput{
			fullName:  candidateFullName,
			dbFighter: f,
			score:     score,
		})
	}

	sort.Sort(out)

	return out, nil
}
