package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/nherson/psc/api/ent/mixins"
)

// FighterResults holds the schema definition for the FighterResults entity.
type FighterResults struct {
	ent.Schema
}

func (FighterResults) Mixins() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
	}
}

// Fields of the FighterResults.
func (FighterResults) Fields() []ent.Field {
	return []ent.Field{
		field.Int("fighter_id").Comment("Required for M2M relationship between Fights and Fighters. NOT UFC assigned identifier!"),
		field.Int("fight_id").Comment("Required for M2M relationship between Fights and Fighters. NOT UFC assigned identifier!"),
		field.Int("significant_strikes_landed"),
		field.Int("takedowns"),
		field.Int("knockdowns"),
		field.Int("control_time_seconds"),
		field.Bool("win_by_stoppage"),
		field.Bool("loss_by_stoppage"),
		field.Bool("missed_weight"),
	}
}

// Edges of the FighterResults.
func (FighterResults) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("fighter", Fighter.Type).
			Required().
			Unique().
			Field("fighter_id"),
		edge.To("fight", Fight.Type).
			Required().
			Unique().
			Field("fight_id"),
	}
}
