package fightodds

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/magefile/mage/sh"

	"github.com/nherson/psc/api/ent"
	"github.com/nherson/psc/api/ent/fighter"
	"github.com/nherson/psc/api/internal/clients/db"
	"github.com/nherson/psc/api/internal/clients/fightodds"
	"github.com/nherson/psc/api/internal/fuzzy"
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

	matcher, err := fuzzy.NewMatcher(
		fuzzy.WithDB(dbClient),
	)
	if err != nil {
		return err
	}

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
		name := fmt.Sprintf("%s %s", fighterOdds.FirstName, fighterOdds.LastName)
		f, score, err := matcher.MatchFighterWithPrompt(ctx, name)
		if f != nil {
			fmt.Printf("Assigning %q FightInsider ID %q\n with score %.2f\n", name, fighterOdds.ID, score)
			dbClient.Fighter.UpdateOne(f).SetFightinsiderID(fighterOdds.ID).SaveX(ctx)
		} else if err == fuzzy.ErrNoMatch {
			fmt.Printf("Skipping FightInsider ID assignment for %q\n", fighterOdds.ID)
		} else if err != nil {
			fmt.Println("Unexpected error finding match:", err.Error())
		}
	}

	return nil
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("could not pull db creds from .env file")
	}
}
