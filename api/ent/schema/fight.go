package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/nherson/psc/api/ent/mixins"
)

// Fight holds the schema definition for the Fight entity.
type Fight struct {
	ent.Schema
}

func (Fight) Mixins() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
	}
}

// Fields of the Fight.
func (Fight) Fields() []ent.Field {
	return []ent.Field{
		field.String("ufc_fight_id").NotEmpty().Unique().Comment("The fight identifier as assigned by UFC"),
		field.Int("card_order").Min(1).Comment("The order in which the fight occurred in the event it took place in. 1 represents a main event, etc."),
		field.String("card_segment").NotEmpty().Comment("Generally, when the fight took place within the card, e.g. 'Main' or 'Prelim', etc"),
		field.String("result_method").NotEmpty().Comment("How the fight ended, e.g. 'Submission' or 'Decision - Unanimous'"),
		field.Int("result_ending_round").Comment("The round when the fight ended; if a decision, will equal the number of rounds in the fight"),
		field.Int("result_ending_time_seconds").Comment("How much time elapsed in the final round of fighting, in seconds; will be the round time if it goes to decision (i.e. 300 seconds)"),
	}
}

// Edges of the Fight.
func (Fight) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("event", Event.Type).
			Ref("fights").
			Unique(),
		edge.To("fighters", Fighter.Type).Through("fighter_results", FighterResults.Type),
	}
}
