package schema

import (
	"entgo.io/ent"
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
	return nil
}

// Edges of the UpcomingEvent.
func (UpcomingEvent) Edges() []ent.Edge {
	return nil
}
