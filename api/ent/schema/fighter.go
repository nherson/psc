package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/nherson/psc/api/ent/mixins"
)

// Fighter holds the schema definition for the Fighter entity.
type Fighter struct {
	ent.Schema
}

func (Fighter) Mixin() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
	}
}

// Fields of the Fighter.
func (Fighter) Fields() []ent.Field {
	return []ent.Field{
		field.String("ufc_fighter_id").NotEmpty().Unique().Comment("The fighter identifier as assigned by UFC"),
		field.Int("mma_id"),
		field.String("first_name"),
		field.String("last_name"),
		field.String("nick_name"),
		field.String("fightinsider_id").Optional(),
		field.String("tapology_id").Optional(),
		field.Bool("temporary").Default(false).Comment("True if the fighter is found in an external site and not UFC official data (e.g. debut fighters)"),
	}
}

// Edges of the Fighter.
func (Fighter) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("fights", Fight.Type).Through("fighter_results", FighterResults.Type).Ref("fighters"),
		edge.From("upcoming_fights", UpcomingFight.Type).Through("upcoming_fighter_odds", UpcomingFighterOdds.Type).Ref("fighters"),
		edge.To("fighter_aliases", FighterAlias.Type),
	}
}
