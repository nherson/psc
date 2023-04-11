// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// EventsColumns holds the columns for the "events" table.
	EventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ufc_event_id", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
	}
	// EventsTable holds the schema information for the "events" table.
	EventsTable = &schema.Table{
		Name:       "events",
		Columns:    EventsColumns,
		PrimaryKey: []*schema.Column{EventsColumns[0]},
	}
	// FightsColumns holds the columns for the "fights" table.
	FightsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "ufc_fight_id", Type: field.TypeString, Unique: true},
		{Name: "card_order", Type: field.TypeInt},
		{Name: "card_segment", Type: field.TypeString},
		{Name: "result_method", Type: field.TypeString},
		{Name: "result_ending_round", Type: field.TypeInt},
		{Name: "result_ending_time_seconds", Type: field.TypeInt},
		{Name: "event_fights", Type: field.TypeInt, Nullable: true},
	}
	// FightsTable holds the schema information for the "fights" table.
	FightsTable = &schema.Table{
		Name:       "fights",
		Columns:    FightsColumns,
		PrimaryKey: []*schema.Column{FightsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "fights_events_fights",
				Columns:    []*schema.Column{FightsColumns[7]},
				RefColumns: []*schema.Column{EventsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FightersColumns holds the columns for the "fighters" table.
	FightersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// FightersTable holds the schema information for the "fighters" table.
	FightersTable = &schema.Table{
		Name:       "fighters",
		Columns:    FightersColumns,
		PrimaryKey: []*schema.Column{FightersColumns[0]},
	}
	// FighterAliasColumns holds the columns for the "fighter_alias" table.
	FighterAliasColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "alias", Type: field.TypeString},
		{Name: "fighter_fighter_aliases", Type: field.TypeInt, Nullable: true},
	}
	// FighterAliasTable holds the schema information for the "fighter_alias" table.
	FighterAliasTable = &schema.Table{
		Name:       "fighter_alias",
		Columns:    FighterAliasColumns,
		PrimaryKey: []*schema.Column{FighterAliasColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "fighter_alias_fighters_fighter_aliases",
				Columns:    []*schema.Column{FighterAliasColumns[2]},
				RefColumns: []*schema.Column{FightersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// FighterResultsColumns holds the columns for the "fighter_results" table.
	FighterResultsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "significant_strikes_landed", Type: field.TypeInt},
		{Name: "takedowns", Type: field.TypeInt},
		{Name: "knockdowns", Type: field.TypeInt},
		{Name: "control_time_seconds", Type: field.TypeInt},
		{Name: "win_by_stoppage", Type: field.TypeBool},
		{Name: "loss_by_stoppage", Type: field.TypeBool},
		{Name: "missed_weight", Type: field.TypeBool},
		{Name: "fighter_id", Type: field.TypeInt},
		{Name: "fight_id", Type: field.TypeInt},
	}
	// FighterResultsTable holds the schema information for the "fighter_results" table.
	FighterResultsTable = &schema.Table{
		Name:       "fighter_results",
		Columns:    FighterResultsColumns,
		PrimaryKey: []*schema.Column{FighterResultsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "fighter_results_fighters_fighter",
				Columns:    []*schema.Column{FighterResultsColumns[8]},
				RefColumns: []*schema.Column{FightersColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "fighter_results_fights_fight",
				Columns:    []*schema.Column{FighterResultsColumns[9]},
				RefColumns: []*schema.Column{FightsColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
		Indexes: []*schema.Index{
			{
				Name:    "fighterresults_fight_id_fighter_id",
				Unique:  true,
				Columns: []*schema.Column{FighterResultsColumns[9], FighterResultsColumns[8]},
			},
		},
	}
	// UpcomingEventsColumns holds the columns for the "upcoming_events" table.
	UpcomingEventsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// UpcomingEventsTable holds the schema information for the "upcoming_events" table.
	UpcomingEventsTable = &schema.Table{
		Name:       "upcoming_events",
		Columns:    UpcomingEventsColumns,
		PrimaryKey: []*schema.Column{UpcomingEventsColumns[0]},
	}
	// UpcomingFightsColumns holds the columns for the "upcoming_fights" table.
	UpcomingFightsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
	}
	// UpcomingFightsTable holds the schema information for the "upcoming_fights" table.
	UpcomingFightsTable = &schema.Table{
		Name:       "upcoming_fights",
		Columns:    UpcomingFightsColumns,
		PrimaryKey: []*schema.Column{UpcomingFightsColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		EventsTable,
		FightsTable,
		FightersTable,
		FighterAliasTable,
		FighterResultsTable,
		UpcomingEventsTable,
		UpcomingFightsTable,
	}
)

func init() {
	FightsTable.ForeignKeys[0].RefTable = EventsTable
	FighterAliasTable.ForeignKeys[0].RefTable = FightersTable
	FighterResultsTable.ForeignKeys[0].RefTable = FightersTable
	FighterResultsTable.ForeignKeys[1].RefTable = FightsTable
}
