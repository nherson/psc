package schema

import (
	"entgo.io/ent"
	"github.com/nherson/psc/api/ent/mixins"
)

// UpcomingFight holds the schema definition for the UpcomingFight entity.
type UpcomingFight struct {
	ent.Schema
}

func (UpcomingFight) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
	}
}

// Fields of the UpcomingFight.
func (UpcomingFight) Fields() []ent.Field {
	return nil
}

// Edges of the UpcomingFight.
func (UpcomingFight) Edges() []ent.Edge {
	return nil
}
