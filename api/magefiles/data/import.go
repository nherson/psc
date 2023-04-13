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
	_ "github.com/lib/pq"
	"github.com/magefile/mage/mg"
	"github.com/pkg/errors"

	"github.com/nherson/psc/api/ent"
	"github.com/nherson/psc/api/ent/event"
	"github.com/nherson/psc/api/ent/fight"
	"github.com/nherson/psc/api/ent/fighter"
	"github.com/nherson/psc/api/ent/fighterresults"
	"github.com/nherson/psc/api/internal/clients/ufc"
	"github.com/nherson/psc/api/magefiles/internal/db"
)

const timestampFormat = "2006-01-02T15:04Z"

type Import mg.Namespace

func (Import) Event(idString string) error {
	return importEvent(idString)
}

// All imports cached UFC stats found locally on disk. Idempotent.
func (Import) All() error {
	r := regexp.MustCompile(`^data\/final\/events\/event-([0-9]+)\.json`)
	return filepath.Walk("data/final/events",
		func(path string, info os.FileInfo, err error) error {
			if path == "data/final/events" {
				return nil
			}

			matches := r.FindStringSubmatch(path)
			if len(matches) != 2 {
				return errors.New("bad file name " + path)
			}

			return importEvent(matches[1])
		})
}

func importEvent(idString string) error {
	ctx := context.Background()

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

	connString := db.ConnectionString()

	dbClient, err := ent.Open("postgres", connString)
	if err != nil {
		return err
	}

	tx, err := dbClient.Tx(ctx)
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
		return db.Rollback(tx, err)
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
			return db.Rollback(tx, err)
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
				return db.Rollback(tx, err)
			}

			// Determine the fight outcome for this fighter
			var stoppage bool
			var winByStoppage bool
			var lossByStoppage bool
			if strings.ToUpper(eventFightData.Result.Method) == "KO/TKO" ||
				strings.ToUpper(eventFightData.Result.Method) == "SUBMISSION" ||
				strings.ToUpper(eventFightData.Result.Method) == "DQ" {
				stoppage = true
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
				OnConflict(
					sql.ConflictColumns(fighterresults.FieldFightID, fighterresults.FieldFighterID),
				).
				UpdateNewValues().
				Exec(ctx)
			if err != nil {
				return db.Rollback(tx, err)
			}
		}

	}

	return tx.Commit()
}
