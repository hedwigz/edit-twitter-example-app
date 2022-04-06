package enthistory

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Changes describe a single change (field that changed) during a CREATE or UPDATE operations.
type Changes struct {
	ent.Schema
}

// Fields of the Changes.
func (Changes) Fields() []ent.Field {
	return []ent.Field{
		field.String("column"),
		field.Text("value"),
		field.Text("previous").Optional(),
	}
}

// Edges of the Changes.
func (Changes) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("operation", History.Type).
			Ref("changes").
			Unique(),
	}
}
