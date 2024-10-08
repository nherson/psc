// Code generated by ent, DO NOT EDIT.

package upcomingevent

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the upcomingevent type in the database.
	Label = "upcoming_event"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedAt holds the string denoting the created_at field in the database.
	FieldCreatedAt = "created_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTapologyID holds the string denoting the tapology_id field in the database.
	FieldTapologyID = "tapology_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDate holds the string denoting the date field in the database.
	FieldDate = "date"
	// EdgeUpcomingFights holds the string denoting the upcoming_fights edge name in mutations.
	EdgeUpcomingFights = "upcoming_fights"
	// Table holds the table name of the upcomingevent in the database.
	Table = "upcoming_events"
	// UpcomingFightsTable is the table that holds the upcoming_fights relation/edge.
	UpcomingFightsTable = "upcoming_fights"
	// UpcomingFightsInverseTable is the table name for the UpcomingFight entity.
	// It exists in this package in order to avoid circular dependency with the "upcomingfight" package.
	UpcomingFightsInverseTable = "upcoming_fights"
	// UpcomingFightsColumn is the table column denoting the upcoming_fights relation/edge.
	UpcomingFightsColumn = "upcoming_event_upcoming_fights"
)

// Columns holds all SQL columns for upcomingevent fields.
var Columns = []string{
	FieldID,
	FieldCreatedAt,
	FieldUpdatedAt,
	FieldTapologyID,
	FieldName,
	FieldDate,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
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
	// TapologyIDValidator is a validator for the "tapology_id" field. It is called by the builders before save.
	TapologyIDValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// Order defines the ordering method for the UpcomingEvent queries.
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

// ByTapologyID orders the results by the tapology_id field.
func ByTapologyID(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldTapologyID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByDate orders the results by the date field.
func ByDate(opts ...sql.OrderTermOption) Order {
	return sql.OrderByField(FieldDate, opts...).ToFunc()
}

// ByUpcomingFightsCount orders the results by upcoming_fights count.
func ByUpcomingFightsCount(opts ...sql.OrderTermOption) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newUpcomingFightsStep(), opts...)
	}
}

// ByUpcomingFights orders the results by upcoming_fights terms.
func ByUpcomingFights(term sql.OrderTerm, terms ...sql.OrderTerm) Order {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newUpcomingFightsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newUpcomingFightsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(UpcomingFightsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, UpcomingFightsTable, UpcomingFightsColumn),
	)
}
