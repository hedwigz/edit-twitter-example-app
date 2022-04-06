package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/field"
	"github.com/ariga/edit-twitter-example-app/enthistory"
)

// Entity holds the schema definition for the Entity entity.
type Entity struct {
	ent.Schema
}

// Fields of the Entity.
func (Entity) Fields() []ent.Field {
	return []ent.Field{
		field.String("data").
			Annotations(enthistory.TrackField()),
		field.Bool("isFun").
			Optional().
			Annotations(enthistory.TrackField()),
		field.Int("counter").
			Optional().
			Annotations(enthistory.TrackField()),
		field.Time("timestamp").
			Optional().
			Annotations(enthistory.TrackField()),
		field.Strings("strings").
			Optional().
			Annotations(enthistory.TrackField()),
	}
}

// Edges of the Entity.
func (Entity) Edges() []ent.Edge {
	return nil
}
