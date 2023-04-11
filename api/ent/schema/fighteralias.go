package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/nherson/psc/api/ent/mixins"
)

// FighterAlias holds the schema definition for the FighterAlias entity.
type FighterAlias struct {
	ent.Schema
}

func (FighterAlias) Mixins() []ent.Mixin {
	return []ent.Mixin{
		mixins.TimeMixin{},
	}
}

// Fields of the FighterAlias.
func (FighterAlias) Fields() []ent.Field {
	return []ent.Field{
		field.String("alias").NotEmpty().Comment("This is some other name a fighter may be known by. No promises about how entries are cased, so program defensively!"),
	}
}

// Edges of the FighterAlias.
func (FighterAlias) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("fighter", Fighter.Type).
			Ref("fighter_aliases").
			Unique(),
	}
}
