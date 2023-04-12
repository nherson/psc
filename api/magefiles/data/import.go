package data

import (
	"context"
	"errors"
	"log"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql"
	_ "github.com/lib/pq"
	"github.com/magefile/mage/mg"

	"github.com/nherson/psc/api/ent"
	"github.com/nherson/psc/api/ent/event"
	"github.com/nherson/psc/api/ent/fight"
	"github.com/nherson/psc/api/ent/fighter"
	"github.com/nherson/psc/api/ent/fighterresults"
	"github.com/nherson/psc/api/internal/clients/ufc"
	"github.com/nherson/psc/api/magefiles/internal/db"
)

type Import mg.Namespace

func (Import) Event(idString string) error {
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

	var fightIDs []int

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
			OnConflictColumns(fight.FieldUfcFightID).
			UpdateNewValues().
			ID(ctx)
		if err != nil {
			return db.Rollback(tx, err)
		}
		fightIDs = append(fightIDs, fightID)

		// var fighterIDs []int

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
			// fighterIDs = append(fighterIDs, fighterID)

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

	err = tx.Event.Create().
		SetUfcEventID(strconv.Itoa(eventData.EventID)).
		SetName(eventData.Name).
		AddFightIDs(fightIDs...).
		OnConflictColumns(event.FieldUfcEventID).
		UpdateNewValues().
		Exec(ctx)
	if err != nil {
		return db.Rollback(tx, err)
	}

	return tx.Commit()
}

func (Import) All() error {
	return errors.New("not implemented")
}
