package fightodds

import (
	"context"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/adrg/strutil"
	"github.com/adrg/strutil/metrics"
	"github.com/joho/godotenv"
	"github.com/magefile/mage/sh"
	"github.com/manifoldco/promptui"

	"github.com/nherson/psc/api/ent"
	"github.com/nherson/psc/api/ent/fighter"
	"github.com/nherson/psc/api/internal/clients/db"
	"github.com/nherson/psc/api/internal/clients/fightodds"
)

func Generate() error {
	pathErr := os.Chdir("api/internal/clients/fightodds")
	defer func() {
		if pathErr == nil {
			os.Chdir("../../../..")
		}
	}()

	err := sh.RunV("go", "run", "github.com/Khan/genqlient")
	if err != nil {
		return err
	}

	return nil
}

func AssignUpcomingIDs(dateString string) error {
	loadEnv()
	ctx := context.Background()
	dbClient := db.MustFromEnv()

	t, err := time.Parse("2006-01-02", dateString)
	if err != nil {
		return err
	}

	allFighters := dbClient.Fighter.Query().
		Select(
			fighter.FieldID,
			fighter.FieldFirstName,
			fighter.FieldLastName,
			fighter.FieldFightinsiderID,
		).AllX(ctx)

	oddsClient := fightodds.NewClient()
	oddsData, err := oddsClient.UpcomingFighterOdds(ctx, t)
	if err != nil {
		return err
	}

	for _, fighterOdds := range oddsData {
		_, err := dbClient.Fighter.Query().Where(fighter.FightinsiderID(fighterOdds.ID)).Only(ctx)
		if err == nil {
			fmt.Printf("ID already set for fighter '%s %s'\n", fighterOdds.FirstName, fighterOdds.LastName)
			continue
		} else if !ent.IsNotFound(err) {
			return err
		}

		fmt.Printf("Doing ID assignment for fighter '%s %s'\n", fighterOdds.FirstName, fighterOdds.LastName)
		selection, err := chooseBestMatch(fighterOdds, allFighters)
		if err != nil {
			return err
		}

		if selection != nil {
			fmt.Printf("Assigning %q FightInsider ID %q\n", selection.fullName, fighterOdds.ID)
			dbClient.Fighter.UpdateOne(selection.dbFighter).SetFightinsiderID(fighterOdds.ID).SaveX(ctx)
		} else {
			fmt.Printf("Skipping FightInsider ID assignment for %q\n", fighterOdds.ID)
		}
	}

	return nil
}

type similarityOutput struct {
	dbFighter *ent.Fighter
	fullName  string
	score     float64
}

type similarities []similarityOutput

func (s similarities) Len() int           { return len(s) }
func (s similarities) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s similarities) Less(i, j int) bool { return s[i].score > s[j].score }

// if (nil, nil) is returned, none were chosen
func chooseBestMatch(fighter fightodds.Fighter, candidates []*ent.Fighter) (*similarityOutput, error) {
	sd := metrics.NewSorensenDice()

	fighterFullName := fmt.Sprintf("%s %s", strings.ToLower(fighter.FirstName), strings.ToLower(fighter.LastName))

	var out similarities
	for _, f := range candidates {
		candidateFullName := fmt.Sprintf("%s %s", strings.ToLower(f.FirstName), strings.ToLower(f.LastName))

		score := strutil.Similarity(fighterFullName, candidateFullName, sd)
		out = append(out, similarityOutput{
			fullName:  candidateFullName,
			dbFighter: f,
			score:     score,
		})
	}

	sort.Sort(out)

	var matchOptions []string
	for i := 0; i < 5; i++ {
		match := out[i]
		matchOptions = append(matchOptions, fmt.Sprintf("Name: %s, Score: %.2f", match.fullName, match.score))
	}
	matchOptions = append(matchOptions, "None of the above")

	prompt := promptui.Select{
		Label: "Choose the best DB match",
		Items: matchOptions,
		Size:  6,
	}
	idx, _, err := prompt.Run()
	if err != nil {
		panic(err)
	}

	if idx == 5 {
		return nil, nil
	} else {
		return &out[idx], nil
	}
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("could not pull db creds from .env file")
	}
}
