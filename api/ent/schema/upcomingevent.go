package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/nherson/psc/api/ent/mixins"
)

// UpcomingEvent holds the schema definition for the UpcomingEvent entity.
type UpcomingEvent struct {
	ent.Schema
}

func (UpcomingEvent) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
	}
}

// Fields of the UpcomingEvent.
func (UpcomingEvent) Fields() []ent.Field {
	return []ent.Field{
		field.String("tapology_id").NotEmpty(),
		field.String("name").NotEmpty(),
		field.Time("date"),
	}
}

// Edges of the UpcomingEvent.
func (UpcomingEvent) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("upcoming_fights", UpcomingFight.Type),
	}
}
