package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// UpcomingFighterOdds holds the schema definition for the UpcomingFighterOdds entity.
type UpcomingFighterOdds struct {
	ent.Schema
}

// Fields of the UpcomingFighterOdds.
func (UpcomingFighterOdds) Fields() []ent.Field {
	return []ent.Field{
		field.Int("fighter_id").Comment("Required for M2M relationship between UpcomingFights and Fighters. NOT tapology assigned identifier!"),
		field.Int("upcoming_fight_id").Comment("Required for M2M relationship between UpcomingFights and Fighters. NOT tapology assigned identifier!"),
		field.Int("best_odds").Optional().Nillable(),
		field.Enum("corner").Values("red", "blue").Default("red"),
	}
}

// Edges of the UpcomingFighterOdds.
func (UpcomingFighterOdds) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("fighter", Fighter.Type).
			Required().
			Unique().
			Field("fighter_id"),
		edge.To("upcoming_fight", UpcomingFight.Type).
			Required().
			Unique().
			Field("upcoming_fight_id"),
	}
}
