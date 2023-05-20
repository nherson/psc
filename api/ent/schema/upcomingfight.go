package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
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
	return []ent.Field{
		field.Int("card_order").Min(1),
	}
}

// Edges of the UpcomingFight.
func (UpcomingFight) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("upcoming_event", UpcomingEvent.Type).
			Ref("upcoming_fights").
			Unique(),
		edge.To("fighters", Fighter.Type).Through("upcoming_fighter_odds", UpcomingFighterOdds.Type),
	}
}
