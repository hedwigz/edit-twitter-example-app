package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("email"),
		field.Int8("role"),
		field.Time("created").
			Default(time.Now),
		field.Int("age").
			Range(0, 1000).
			Optional().
			Nillable(),
		field.String("image_url"),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("tweets", Tweet.Type),
		edge.To("following", User.Type).
			From("followers"),
		edge.From("liked", Tweet.Type).
			Ref("likes"),
	}
}
