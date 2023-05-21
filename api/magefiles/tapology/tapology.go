package tapology

import (
	"context"
	"fmt"
	"time"

	"github.com/joho/godotenv"
	"github.com/nherson/psc/api/ent"
	"github.com/nherson/psc/api/ent/fighter"
	"github.com/nherson/psc/api/internal/clients/db"
	"github.com/nherson/psc/api/internal/clients/tapology"
	"github.com/nherson/psc/api/internal/fuzzy"
)

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

	tapologyClient := tapology.NewClient()
	data, err := tapologyClient.UpcomingEvents(ctx, t)
	if err != nil {
		return err
	}

	for _, event := range data {
		for _, fight := range event.Fights {
			err := assignFighterID(ctx, dbClient, matcher, fight.Fighter1)
			if err != nil {
				return err
			}
			err = assignFighterID(ctx, dbClient, matcher, fight.Fighter2)
			if err != nil {
				return err
			}
		}

	}

	return nil
}

func assignFighterID(ctx context.Context, dbClient *ent.Client, matcher *fuzzy.Matcher, tf tapology.Fighter) error {
	_, err := dbClient.Fighter.Query().Where(fighter.TapologyID(tf.ID)).Only(ctx)
	if err == nil {
		fmt.Printf("ID already set for fighter '%s %s'\n", tf.FirstName, tf.LastName)
		return nil
	} else if !ent.IsNotFound(err) {
		return err
	}

	fmt.Printf("Doing ID assignment for fighter '%s %s'\n", tf.FirstName, tf.LastName)
	name := fmt.Sprintf("%s %s", tf.FirstName, tf.LastName)
	f, score, err := matcher.MatchFighterWithPrompt(ctx, name)
	if f != nil {
		fmt.Printf("Assigning %q Tapology ID %q\n with score %.2f\n", name, tf.ID, score)
		dbClient.Fighter.UpdateOne(f).SetTapologyID(tf.ID).SaveX(ctx)
	} else if err == fuzzy.ErrNoMatch {
		fmt.Printf("Skipping Tapology ID assignment for %q\n", name)
	} else if err != nil {
		fmt.Println("Unexpected error finding match:", err.Error())
	}

	return nil
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("could not pull db creds from .env file")
	}
}
