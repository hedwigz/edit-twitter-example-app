package schema

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/ariga/edit-twitter-example-app/enthistory"
)

// Tweet holds the schema definition for the Tweet entity.
type Tweet struct {
	ent.Schema
}

const (
	MaxTweetFieldsLen = 255
)

// Fields of the Tweet.
func (Tweet) Fields() []ent.Field {
	return []ent.Field{
		field.String("content").Validate(func(c string) error {
			strFields := strings.Fields(c)
			if len(strFields) > MaxTweetFieldsLen {
				return fmt.Errorf("tweet content cannot have more than %d", MaxTweetFieldsLen)
			}
			return nil
		}).Annotations(enthistory.Annotation{}),
		field.Time("created").
			Default(time.Now).
			Annotations(
				entgql.OrderField("CREATED_AT"),
			),
	}
}

// Edges of the Tweet.
func (Tweet) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("author", User.Type).
			Ref("tweets").
			Unique(),
		edge.To("likes", User.Type),
	}
}

// func (Tweet) Policy() ent.Policy {
// 	return privacy.Policy{
// 		Mutation: privacy.MutationPolicy{
// 			// Deny if not set otherwise.
// 			privacy.OnMutationOperation(bootcamp.FilterSelf(), ent.OpDelete|ent.OpDeleteOne),
// 		},
// 		Query: privacy.QueryPolicy{
// 			// Allow any viewer to read anything.
// 			privacy.AlwaysAllowRule(),
// 		},
// 	}
// }
