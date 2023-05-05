package data

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/magefile/mage/mg"
	"github.com/pkg/errors"

	"github.com/nherson/psc/api/ent"
	"github.com/nherson/psc/api/ent/event"
	"github.com/nherson/psc/api/ent/fight"
	"github.com/nherson/psc/api/ent/fighter"
	"github.com/nherson/psc/api/ent/fighterresults"
	"github.com/nherson/psc/api/internal/clients/db"
	"github.com/nherson/psc/api/internal/clients/ufc"
)

const timestampFormat = "2006-01-02T15:04Z"

type Import mg.Namespace

func (Import) Event(idString string) error {
	ctx := context.Background()
	loadEnv()
	dbClient := db.MustFromEnv()

	return importEvent(ctx, dbClient, idString)
}

// All imports cached UFC stats found locally on disk. Idempotent.
func (Import) All() error {
	ctx := context.Background()
	loadEnv()
	dbClient := db.MustFromEnv()

	var eventIDs []string

	r := regexp.MustCompile(`^data\/final\/events\/event-([0-9]+)\.json`)
	err := filepath.Walk("data/final/events",
		func(path string, info os.FileInfo, err error) error {
			if path == "data/final/events" {
				return nil
			}

			matches := r.FindStringSubmatch(path)
			if len(matches) != 2 {
				return errors.New("bad file name " + path)
			}

			eventIDs = append(eventIDs, matches[1])

			return nil
		},
	)
	if err != nil {
		return err
	}

	for _, e := range eventIDs {
		err := importEvent(ctx, dbClient, e)
		if err != nil {
			return err
		}
	}

	return nil
}

func importEvent(ctx context.Context, dbClient *ent.Client, idString string) error {
	tx, err := dbClient.Tx(ctx)
	if err != nil {
		return err
	}

	err = importEventTx(ctx, tx, dbClient, idString)
	if err != nil {
		db.Rollback(tx, err)
		return err
	}

	return tx.Commit()
}

func importEventTx(ctx context.Context, tx *ent.Tx, dbClient *ent.Client, idString string) error {
	id64, err := strconv.ParseInt(idString, 10, 32)
	if err != nil {
		return err
	}
	id := int(id64)

	log.Printf("doing import of event with id %d\n", id)

	client := ufc.NewCacheClient()
	eventData, err := client.EventByID(id)
	if err != nil {
		return err
	}

	date, err := time.Parse(timestampFormat, eventData.StartTime)
	if err != nil {
		return errors.Wrap(err, "could not parse event start time")
	}
	date = time.Date(date.Year(), date.Month(), date.Day(), 0, 0, 0, 0, date.Location())

	eventID, err := tx.Event.Create().
		SetUfcEventID(strconv.Itoa(eventData.EventID)).
		SetName(eventData.Name).
		SetDate(date).
		OnConflictColumns(event.FieldUfcEventID).
		UpdateNewValues().
		ID(ctx)
	if err != nil {
		return err
	}

	// For each fight, upsert both fighters, upsert the fight, link them together
	for _, eventFightData := range eventData.FightCard {
		log.Printf("doing import of fight with id %d\n", eventFightData.FightID)

		// Get detailed fight stats
		fightData, err := client.FightByID(eventFightData.FightID)
		if err != nil {
			return db.Rollback(tx, err)
		}
		if fightData.FightStats == nil || len(fightData.FightStats) != 2 {
			return db.Rollback(tx, errors.New("bad FightStats data"))
		}

		// Upsert fight
		fightID, err := tx.Fight.Create().
			SetUfcFightID(strconv.Itoa(eventFightData.FightID)).
			SetCardOrder(eventFightData.FightOrder).
			SetCardSegment(eventFightData.CardSegment).
			SetResultEndingRound(eventFightData.Result.EndingRound).
			SetResultEndingTimeSeconds(eventFightData.Result.EndingTimeSeconds()).
			SetResultMethod(eventFightData.Result.Method).
			SetEventID(eventID).
			OnConflictColumns(fight.FieldUfcFightID).
			UpdateNewValues().
			ID(ctx)
		if err != nil {
			return err
		}

		if len(eventFightData.Fighters) != 2 {
			// Should never happen
			return errors.New("unexpected number of fighters in fight")
		}

		// Upsert fighters
		for _, f := range eventFightData.Fighters {
			log.Printf("doing import of fighter with id %d\n", f.FighterID)

			fighterID, err := tx.Fighter.Create().
				SetUfcFighterID(strconv.Itoa(f.FighterID)).
				SetMmaID(f.MMAID).
				SetFirstName(f.Name.FirstName).
				SetLastName(f.Name.LastName).
				SetNickName(f.Name.NickName).
				OnConflictColumns(fighter.FieldUfcFighterID).
				UpdateNewValues().
				ID(ctx)
			if err != nil {
				return err
			}

			// Determine the fight outcome for this fighter
			var stoppage bool
			var win bool
			var winByStoppage bool
			var lossByStoppage bool
			if strings.ToUpper(eventFightData.Result.Method) == "KO/TKO" ||
				strings.ToUpper(eventFightData.Result.Method) == "SUBMISSION" ||
				strings.ToUpper(eventFightData.Result.Method) == "DQ" {
				stoppage = true
			}

			if f.Outcome.OutcomeID == 1 {
				win = true
			}

			if stoppage && f.Outcome.OutcomeID == 2 {
				lossByStoppage = true
			} else if stoppage && f.Outcome.OutcomeID == 1 {
				winByStoppage = true
			}

			// Find detailed fight results for figther
			var fighterStats ufc.FighterStats
			if f.FighterID == fightData.FightStats[0].FighterID {
				fighterStats = fightData.FightStats[0]
			} else if f.FighterID == fightData.FightStats[1].FighterID {
				fighterStats = fightData.FightStats[1]
			} else {
				return errors.New("fighter not found in FightStats")
			}

			// Get the fighter's corner to get consistent rendering in the UI
			var corner fighterresults.Corner
			if strings.ToLower(f.Corner) == "red" {
				corner = fighterresults.CornerRed
			} else if strings.ToLower(f.Corner) == "blue" {
				corner = fighterresults.CornerBlue
			} else {
				return errors.New("unrecognized corner: '" + f.Corner + "'")
			}

			// link fighter to fight via M2M relationship
			err = tx.FighterResults.Create().
				SetFightID(fightID).
				SetFighterID(fighterID).
				SetSignificantStrikesLanded(fighterStats.SigStrikesLanded).
				SetControlTimeSeconds(fighterStats.ControlTimeSeconds()).
				SetKnockdowns(fighterStats.Knockdowns).
				SetTakedowns(fighterStats.TakedownsLanded).
				SetWinByStoppage(winByStoppage).
				SetLossByStoppage(lossByStoppage).
				SetWin(win).
				SetCorner(corner).
				OnConflict(
					sql.ConflictColumns(fighterresults.FieldFightID, fighterresults.FieldFighterID),
				).
				UpdateNewValues().
				Exec(ctx)
			if err != nil {
				return err
			}
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
