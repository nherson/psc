package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"github.com/nherson/psc/api/ent/mixins"
)

// Fighter holds the schema definition for the Fighter entity.
type Fighter struct {
	ent.Schema
}

func (Fighter) Mixins() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
	}
}

// Fields of the Fighter.
func (Fighter) Fields() []ent.Field {
	return nil
}

// Edges of the Fighter.
func (Fighter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("fights", Fight.Type).Through("fighter_results", FighterResults.Type).Ref("fighters"),
		edge.To("fighter_aliases", FighterAlias.Type),
	}
}
