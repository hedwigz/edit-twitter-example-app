package enthistory

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// History is a schema for logging operations of history tracked schemas.
type History struct {
	ent.Schema
}

func (History) Fields() []ent.Field {
	return []ent.Field{
		// start default fields
		field.String("entity_name"),
		field.Int("record_id"),
		field.Time("timestamp"),
		field.Enum("action").
			NamedValues(
				"Create", "CREATE",
				"Update", "UPDATE",
				"Delete", "DELETE",
			),
	}
}

// Edges of the History.
func (History) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("changes", Changes.Type),
	}
}
