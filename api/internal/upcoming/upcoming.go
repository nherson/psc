package upcoming

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/nherson/psc/api/ent"
	"github.com/nherson/psc/api/ent/fighter"
	"github.com/nherson/psc/api/ent/upcomingfighterodds"
	"github.com/nherson/psc/api/internal/clients/db"
	"github.com/nherson/psc/api/internal/clients/fightodds"
	"github.com/nherson/psc/api/internal/clients/tapology"
	"github.com/nherson/psc/api/internal/fuzzy"
	"github.com/pkg/errors"
)

type matchFunction func(context.Context, *ent.Tx, string) (*ent.Fighter, float64, error)

type matchFinder interface {
	findMatch(context.Context, *ent.Tx, string) (*ent.Fighter, float64, error)
	clearCache()
}

type promptMatcher struct {
	m *fuzzy.Matcher
}

func (pm *promptMatcher) findMatch(ctx context.Context, tx *ent.Tx, fullName string) (*ent.Fighter, float64, error) {
	return pm.m.MatchFighterWithPromptTx(ctx, tx, fullName)
}

func (pm *promptMatcher) clearCache() {
	pm.m.ClearCache()
}

type noPromptMatcher struct {
	m *fuzzy.Matcher
}

func (pm *noPromptMatcher) findMatch(ctx context.Context, tx *ent.Tx, fullName string) (*ent.Fighter, float64, error) {
	return pm.m.MatchFighterTx(ctx, tx, fullName)
}

func (pm *noPromptMatcher) clearCache() {
	pm.m.ClearCache()
}

func Import(
	ctx context.Context,
	dbClient *ent.Client,
	tapologyClient *tapology.Client,
	fightOddsClient *fightodds.Client,
	matcher *fuzzy.Matcher,
	until time.Time,
) error {
	return doImport(ctx, dbClient, tapologyClient, fightOddsClient, &noPromptMatcher{m: matcher}, until)
}

func ImportWithPrompt(
	ctx context.Context,
	dbClient *ent.Client,
	tapologyClient *tapology.Client,
	fightOddsClient *fightodds.Client,
	matcher *fuzzy.Matcher,
	until time.Time,
) error {
	return doImport(ctx, dbClient, tapologyClient, fightOddsClient, &promptMatcher{m: matcher}, until)
}

func doImport(
	ctx context.Context,
	dbClient *ent.Client,
	tapologyClient *tapology.Client,
	fightOddsClient *fightodds.Client,
	matcher matchFinder,
	until time.Time,
) error {
	tx, err := dbClient.Tx(ctx)
	if err != nil {
		return db.Rollback(tx, err)
	}

	err = purgeUpcoming(ctx, tx)
	if err != nil {
		return db.Rollback(tx, errors.Wrap(err, "could not purge old upcoming data"))
	}

	tapologyData, err := tapologyClient.UpcomingEvents(ctx, until)
	if err != nil {
		return errors.Wrap(err, "could not fetch tapology data")
	}

	fightOddsData, err := fightOddsClient.UpcomingFighterOdds(ctx, until)
	if err != nil {
		return errors.Wrap(err, "could not fetch fight odds data")
	}

	// Import events
	for _, tapologyEvent := range tapologyData {
		fmt.Printf("Creating upcoming event %q\n", tapologyEvent.Name)
		event, err := tx.UpcomingEvent.Create().
			SetTapologyID(tapologyEvent.ID).
			SetName(tapologyEvent.Name).
			SetDate(tapologyEvent.Date).
			Save(ctx)
		if err != nil {
			return errors.Wrapf(err, "could not create upcoming event %q", tapologyEvent.Name)
		}

		for i, tapologyFight := range tapologyEvent.Fights {
			fmt.Printf(
				"Creating upcoming fight '%s %s vs %s %s'\n",
				tapologyFight.Fighter1.FirstName,
				tapologyFight.Fighter1.LastName,
				tapologyFight.Fighter2.FirstName,
				tapologyFight.Fighter2.LastName,
			)
			fight, err := tx.UpcomingFight.Create().
				SetUpcomingEvent(event).
				SetCardOrder(tapologyFight.BoutNumber).
				Save(ctx)
			if err != nil {
				return errors.Wrapf(err, "could not create upcoming fight for event %q at index %d", tapologyEvent.Name, i)
			}

			fighterColor := map[int]upcomingfighterodds.Corner{
				0: upcomingfighterodds.CornerRed,
				1: upcomingfighterodds.CornerBlue,
			}
			for i, tapologyFighter := range []tapology.Fighter{tapologyFight.Fighter1, tapologyFight.Fighter2} {
				color := fighterColor[i]

				fighter, err := matchOrCreateTapologyFighter(ctx, tx, tapologyFighter, matcher.findMatch)
				if err != nil {
					return errors.Wrapf(err, "could not associate fighter '%s %s' from tapology data", tapologyFighter.FirstName, tapologyFighter.LastName)
				}

				_, err = tx.UpcomingFighterOdds.Create().
					SetCorner(color).
					SetFighter(fighter).
					SetUpcomingFight(fight).
					Save(ctx)
				if err != nil {
					return errors.Wrapf(err, "could not associate fighter '%s %s' to upcoming fight", fighter.FirstName, fighter.LastName)
				}

			}
		}
	}

	// Clear the matcher cache because the above code may have added new fighters that the below code needs to reference for odds injection
	matcher.clearCache()

	for _, fighterOdds := range fightOddsData {
		fmt.Printf("Attempting odds data association for '%s %s'\n", fighterOdds.FirstName, fighterOdds.LastName)
		fighter, _, err := matcher.findMatch(ctx, tx, fmt.Sprintf("%s %s", fighterOdds.FirstName, fighterOdds.LastName))
		if err == fuzzy.ErrNoMatch {
			fmt.Printf("Skipping odds association for fighter '%s %s'\n", fighterOdds.FirstName, fighterOdds.LastName)
			continue
		}

		oddsEdge, err := tx.UpcomingFighterOdds.Query().Where(upcomingfighterodds.FighterID(fighter.ID)).Only(ctx)
		if ent.IsNotFound(err) {
			fmt.Printf("Could not find upcoming fight to associate odds for fighter '%s %s'\n", fighter.FirstName, fighter.LastName)
			continue
		} else if err != nil {
			fmt.Printf("Unexpected error associating odds data for fighter '%s %s': %v\n", fighter.FirstName, fighter.LastName, err)
			continue
		}

		_, err = oddsEdge.Update().SetBestOdds(fighterOdds.BestOdds).Save(ctx)
		if err != nil {
			fmt.Printf("Could not associate odds data with upcoming fighter '%s %s'\n", fighter.FirstName, fighter.LastName)
		}
	}

	return tx.Commit()
}

func matchOrCreateTapologyFighter(ctx context.Context, tx *ent.Tx, tf tapology.Fighter, matchFn matchFunction) (*ent.Fighter, error) {
	fighter, err := tx.Fighter.Query().Where(fighter.TapologyID(tf.ID)).Only(ctx)
	if err == nil {
		return fighter, nil
	} else if !ent.IsNotFound(err) {
		return nil, errors.Wrap(err, "could not query for existing tapology fighter link in db")
	}

	// Need to establish link or create temporary fighter node
	fighter, _, err = matchFn(ctx, tx, fmt.Sprintf("%s %s", tf.FirstName, tf.LastName))
	if err == fuzzy.ErrNoMatch {
		// create a temporary node
		fighter, err = tx.Fighter.Create().
			SetUfcFighterID("DOES-NOT-EXIST-" + uuid.NewString()).
			SetMmaID(0).
			SetNickName("").
			SetFirstName(tf.FirstName).
			SetLastName(tf.LastName).
			SetTapologyID(tf.ID).
			SetTemporary(true).
			Save(ctx)
		if err != nil {
			return nil, errors.Wrap(err, "could not create temporary fighter node")
		}
	} else if err != nil {
		return nil, errors.Wrap(err, "could not perform fighter fuzzy match")
	} else if fighter != nil {
		fighter, err = fighter.Update().SetTapologyID(tf.ID).Save(ctx)
	} else {
		// this shouldnt happen
		return nil, errors.New("unexpected outcome")
	}

	return fighter, err
}

func purgeUpcoming(ctx context.Context, tx *ent.Tx) error {
	_, err := tx.UpcomingFighterOdds.Delete().Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "could not purge upcoming fighter odds edges")
	}

	_, err = tx.Fighter.Delete().Where(fighter.Temporary(true)).Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "could not purge temporary fighters")
	}

	_, err = tx.UpcomingFight.Delete().Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "could not purge upcoming fights")
	}
	_, err = tx.UpcomingEvent.Delete().Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "could not purge upcoming events")
	}

	return nil
}
