// Code generated by ent, DO NOT EDIT.

package upcomingfight

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the upcomingfight type in the database.
	Label = "upcoming_fight"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldCardOrder holds the string denoting the card_order field in the database.
	FieldCardOrder = "card_order"
	// EdgeUpcomingEvent holds the string denoting the upcoming_event edge name in mutations.
	EdgeUpcomingEvent = "upcoming_event"
	// EdgeFighters holds the string denoting the fighters edge name in mutations.
	EdgeFighters = "fighters"
	// EdgeUpcomingFighterOdds holds the string denoting the upcoming_fighter_odds edge name in mutations.
	EdgeUpcomingFighterOdds = "upcoming_fighter_odds"
	// Table holds the table name of the upcomingfight in the database.
	Table = "upcoming_fights"
	// UpcomingEventTable is the table that holds the upcoming_event relation/edge.
	UpcomingEventTable = "upcoming_fights"
	// UpcomingEventInverseTable is the table name for the UpcomingEvent entity.
	// It exists in this package in order to avoid circular dependency with the "upcomingevent" package.
	UpcomingEventInverseTable = "upcoming_events"
	// UpcomingEventColumn is the table column denoting the upcoming_event relation/edge.
	UpcomingEventColumn = "upcoming_event_upcoming_fights"
	// FightersTable is the table that holds the fighters relation/edge. The primary key declared below.
	FightersTable = "upcoming_fighter_odds"
	// FightersInverseTable is the table name for the Fighter entity.
	// It exists in this package in order to avoid circular dependency with the "fighter" package.
	FightersInverseTable = "fighters"
	// UpcomingFighterOddsTable is the table that holds the upcoming_fighter_odds relation/edge.
	UpcomingFighterOddsTable = "upcoming_fighter_odds"
	// UpcomingFighterOddsInverseTable is the table name for the UpcomingFighterOdds entity.
	// It exists in this package in order to avoid circular dependency with the "upcomingfighterodds" package.
	UpcomingFighterOddsInverseTable = "upcoming_fighter_odds"
	// UpcomingFighterOddsColumn is the table column denoting the upcoming_fighter_odds relation/edge.
	UpcomingFighterOddsColumn = "upcoming_fight_id"
)

// Columns holds all SQL columns for upcomingfight fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldCardOrder,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "upcoming_fights"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"upcoming_event_upcoming_fights",
}

var (
	// FightersPrimaryKey and FightersColumn2 are the table columns denoting the
	// primary key for the fighters relation (M2M).
	FightersPrimaryKey = []string{"upcoming_fight_id", "fighter_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedAt holds the default value on creation for the "created_at" field.
	DefaultCreatedAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// CardOrderValidator is a validator for the "card_order" field. It is called by the builders before save.
	CardOrderValidator func(int) error
)

// Order defines the ordering method for the UpcomingFight queries.
type Order func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCreatedAt orders the results by the created_at field.
func ByCreatedAt(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldCreatedAt, opts...).ToFunc()
}

// ByUpdatedAt orders the results by the updated_at field.
func ByUpdatedAt(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldUpdatedAt, opts...).ToFunc()
}

// ByCardOrder orders the results by the card_order field.
func ByCardOrder(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldCardOrder, opts...).ToFunc()
}

// ByUpcomingEventField orders the results by upcoming_event field.
func ByUpcomingEventField(field string, opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUpcomingEventStep(), sql.OrderByField(field, opts...))
	}
}

// ByFightersCount orders the results by fighters count.
func ByFightersCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newFightersStep(), opts...)
	}
}

// ByFighters orders the results by fighters terms.
func ByFighters(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFightersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByUpcomingFighterOddsCount orders the results by upcoming_fighter_odds count.
func ByUpcomingFighterOddsCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUpcomingFighterOddsStep(), opts...)
	}
}

// ByUpcomingFighterOdds orders the results by upcoming_fighter_odds terms.
func ByUpcomingFighterOdds(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUpcomingFighterOddsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUpcomingEventStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UpcomingEventInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, UpcomingEventTable, UpcomingEventColumn),
	)
}
func newFightersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FightersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, FightersTable, FightersPrimaryKey...),
	)
}
func newUpcomingFighterOddsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UpcomingFighterOddsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, UpcomingFighterOddsTable, UpcomingFighterOddsColumn),
	)
}
