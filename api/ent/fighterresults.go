// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/nherson/psc/api/ent/fight"
	"github.com/nherson/psc/api/ent/fighter"
	"github.com/nherson/psc/api/ent/fighterresults"
)

// FighterResults is the model entity for the FighterResults schema.
type FighterResults struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Required for M2M relationship between Fights and Fighters. NOT UFC assigned identifier!
	FighterID int `json:"fighter_id,omitempty"`
	// Required for M2M relationship between Fights and Fighters. NOT UFC assigned identifier!
	FightID int `json:"fight_id,omitempty"`
	// Corner holds the value of the "corner" field.
	Corner fighterresults.Corner `json:"corner,omitempty"`
	// SignificantStrikesLanded holds the value of the "significant_strikes_landed" field.
	SignificantStrikesLanded int `json:"significant_strikes_landed,omitempty"`
	// Takedowns holds the value of the "takedowns" field.
	Takedowns int `json:"takedowns,omitempty"`
	// Knockdowns holds the value of the "knockdowns" field.
	Knockdowns int `json:"knockdowns,omitempty"`
	// ControlTimeSeconds holds the value of the "control_time_seconds" field.
	ControlTimeSeconds int `json:"control_time_seconds,omitempty"`
	// Win holds the value of the "win" field.
	Win bool `json:"win,omitempty"`
	// WinByStoppage holds the value of the "win_by_stoppage" field.
	WinByStoppage bool `json:"win_by_stoppage,omitempty"`
	// LossByStoppage holds the value of the "loss_by_stoppage" field.
	LossByStoppage bool `json:"loss_by_stoppage,omitempty"`
	// MissedWeight holds the value of the "missed_weight" field.
	MissedWeight bool `json:"missed_weight,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the FighterResultsQuery when eager-loading is set.
	Edges        FighterResultsEdges `json:"edges"`
	selectValues sql.SelectValues
}

// FighterResultsEdges holds the relations/edges for other nodes in the graph.
type FighterResultsEdges struct {
	// Fighter holds the value of the fighter edge.
	Fighter *Fighter `json:"fighter,omitempty"`
	// Fight holds the value of the fight edge.
	Fight *Fight `json:"fight,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// FighterOrErr returns the Fighter value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FighterResultsEdges) FighterOrErr() (*Fighter, error) {
	if e.loadedTypes[0] {
		if e.Fighter == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: fighter.Label}
		}
		return e.Fighter, nil
	}
	return nil, &NotLoadedError{edge: "fighter"}
}

// FightOrErr returns the Fight value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e FighterResultsEdges) FightOrErr() (*Fight, error) {
	if e.loadedTypes[1] {
		if e.Fight == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: fight.Label}
		}
		return e.Fight, nil
	}
	return nil, &NotLoadedError{edge: "fight"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*FighterResults) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case fighterresults.FieldWin, fighterresults.FieldWinByStoppage, fighterresults.FieldLossByStoppage, fighterresults.FieldMissedWeight:
			values[i] = new(sql.NullBool)
		case fighterresults.FieldID, fighterresults.FieldFighterID, fighterresults.FieldFightID, fighterresults.FieldSignificantStrikesLanded, fighterresults.FieldTakedowns, fighterresults.FieldKnockdowns, fighterresults.FieldControlTimeSeconds:
			values[i] = new(sql.NullInt64)
		case fighterresults.FieldCorner:
			values[i] = new(sql.NullString)
		case fighterresults.FieldCreatedAt, fighterresults.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the FighterResults fields.
func (fr *FighterResults) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case fighterresults.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			fr.ID = int(value.Int64)
		case fighterresults.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				fr.CreatedAt = value.Time
			}
		case fighterresults.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				fr.UpdatedAt = value.Time
			}
		case fighterresults.FieldFighterID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field fighter_id", values[i])
			} else if value.Valid {
				fr.FighterID = int(value.Int64)
			}
		case fighterresults.FieldFightID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field fight_id", values[i])
			} else if value.Valid {
				fr.FightID = int(value.Int64)
			}
		case fighterresults.FieldCorner:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field corner", values[i])
			} else if value.Valid {
				fr.Corner = fighterresults.Corner(value.String)
			}
		case fighterresults.FieldSignificantStrikesLanded:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field significant_strikes_landed", values[i])
			} else if value.Valid {
				fr.SignificantStrikesLanded = int(value.Int64)
			}
		case fighterresults.FieldTakedowns:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field takedowns", values[i])
			} else if value.Valid {
				fr.Takedowns = int(value.Int64)
			}
		case fighterresults.FieldKnockdowns:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field knockdowns", values[i])
			} else if value.Valid {
				fr.Knockdowns = int(value.Int64)
			}
		case fighterresults.FieldControlTimeSeconds:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field control_time_seconds", values[i])
			} else if value.Valid {
				fr.ControlTimeSeconds = int(value.Int64)
			}
		case fighterresults.FieldWin:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field win", values[i])
			} else if value.Valid {
				fr.Win = value.Bool
			}
		case fighterresults.FieldWinByStoppage:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field win_by_stoppage", values[i])
			} else if value.Valid {
				fr.WinByStoppage = value.Bool
			}
		case fighterresults.FieldLossByStoppage:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field loss_by_stoppage", values[i])
			} else if value.Valid {
				fr.LossByStoppage = value.Bool
			}
		case fighterresults.FieldMissedWeight:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field missed_weight", values[i])
			} else if value.Valid {
				fr.MissedWeight = value.Bool
			}
		default:
			fr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the FighterResults.
// This includes values selected through modifiers, order, etc.
func (fr *FighterResults) Value(name string) (ent.Value, error) {
	return fr.selectValues.Get(name)
}

// QueryFighter queries the "fighter" edge of the FighterResults entity.
func (fr *FighterResults) QueryFighter() *FighterQuery {
	return NewFighterResultsClient(fr.config).QueryFighter(fr)
}

// QueryFight queries the "fight" edge of the FighterResults entity.
func (fr *FighterResults) QueryFight() *FightQuery {
	return NewFighterResultsClient(fr.config).QueryFight(fr)
}

// Update returns a builder for updating this FighterResults.
// Note that you need to call FighterResults.Unwrap() before calling this method if this FighterResults
// was returned from a transaction, and the transaction was committed or rolled back.
func (fr *FighterResults) Update() *FighterResultsUpdateOne {
	return NewFighterResultsClient(fr.config).UpdateOne(fr)
}

// Unwrap unwraps the FighterResults entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (fr *FighterResults) Unwrap() *FighterResults {
	_tx, ok := fr.config.driver.(*txDriver)
	if !ok {
		panic("ent: FighterResults is not a transactional entity")
	}
	fr.config.driver = _tx.drv
	return fr
}

// String implements the fmt.Stringer.
func (fr *FighterResults) String() string {
	var builder strings.Builder
	builder.WriteString("FighterResults(")
	builder.WriteString(fmt.Sprintf("id=%v, ", fr.ID))
	builder.WriteString("created_at=")
	builder.WriteString(fr.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_at=")
	builder.WriteString(fr.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("fighter_id=")
	builder.WriteString(fmt.Sprintf("%v", fr.FighterID))
	builder.WriteString(", ")
	builder.WriteString("fight_id=")
	builder.WriteString(fmt.Sprintf("%v", fr.FightID))
	builder.WriteString(", ")
	builder.WriteString("corner=")
	builder.WriteString(fmt.Sprintf("%v", fr.Corner))
	builder.WriteString(", ")
	builder.WriteString("significant_strikes_landed=")
	builder.WriteString(fmt.Sprintf("%v", fr.SignificantStrikesLanded))
	builder.WriteString(", ")
	builder.WriteString("takedowns=")
	builder.WriteString(fmt.Sprintf("%v", fr.Takedowns))
	builder.WriteString(", ")
	builder.WriteString("knockdowns=")
	builder.WriteString(fmt.Sprintf("%v", fr.Knockdowns))
	builder.WriteString(", ")
	builder.WriteString("control_time_seconds=")
	builder.WriteString(fmt.Sprintf("%v", fr.ControlTimeSeconds))
	builder.WriteString(", ")
	builder.WriteString("win=")
	builder.WriteString(fmt.Sprintf("%v", fr.Win))
	builder.WriteString(", ")
	builder.WriteString("win_by_stoppage=")
	builder.WriteString(fmt.Sprintf("%v", fr.WinByStoppage))
	builder.WriteString(", ")
	builder.WriteString("loss_by_stoppage=")
	builder.WriteString(fmt.Sprintf("%v", fr.LossByStoppage))
	builder.WriteString(", ")
	builder.WriteString("missed_weight=")
	builder.WriteString(fmt.Sprintf("%v", fr.MissedWeight))
	builder.WriteByte(')')
	return builder.String()
}

// FighterResultsSlice is a parsable slice of FighterResults.
type FighterResultsSlice []*FighterResults
