// Code generated by entc, DO NOT EDIT.

package ent

import (
	"github.com/ariga/edit-twitter-example-app/ent/changes"
	"github.com/ariga/edit-twitter-example-app/ent/history"
	"github.com/ariga/edit-twitter-example-app/ent/predicate"
	"github.com/ariga/edit-twitter-example-app/ent/tweet"
	"github.com/ariga/edit-twitter-example-app/ent/user"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entql"
	"entgo.io/ent/schema/field"
)

// schemaGraph holds a representation of ent/schema at runtime.
var schemaGraph = func() *sqlgraph.Schema {
	graph := &sqlgraph.Schema{Nodes: make([]*sqlgraph.Node, 4)}
	graph.Nodes[0] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   changes.Table,
			Columns: changes.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: changes.FieldID,
			},
		},
		Type: "Changes",
		Fields: map[string]*sqlgraph.FieldSpec{
			changes.FieldColumn:   {Type: field.TypeString, Column: changes.FieldColumn},
			changes.FieldValue:    {Type: field.TypeString, Column: changes.FieldValue},
			changes.FieldPrevious: {Type: field.TypeString, Column: changes.FieldPrevious},
		},
	}
	graph.Nodes[1] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   history.Table,
			Columns: history.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: history.FieldID,
			},
		},
		Type: "History",
		Fields: map[string]*sqlgraph.FieldSpec{
			history.FieldEntityName: {Type: field.TypeString, Column: history.FieldEntityName},
			history.FieldRecordID:   {Type: field.TypeInt, Column: history.FieldRecordID},
			history.FieldTimestamp:  {Type: field.TypeTime, Column: history.FieldTimestamp},
			history.FieldAction:     {Type: field.TypeEnum, Column: history.FieldAction},
		},
	}
	graph.Nodes[2] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   tweet.Table,
			Columns: tweet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: tweet.FieldID,
			},
		},
		Type: "Tweet",
		Fields: map[string]*sqlgraph.FieldSpec{
			tweet.FieldContent: {Type: field.TypeString, Column: tweet.FieldContent},
			tweet.FieldCreated: {Type: field.TypeTime, Column: tweet.FieldCreated},
		},
	}
	graph.Nodes[3] = &sqlgraph.Node{
		NodeSpec: sqlgraph.NodeSpec{
			Table:   user.Table,
			Columns: user.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: user.FieldID,
			},
		},
		Type: "User",
		Fields: map[string]*sqlgraph.FieldSpec{
			user.FieldName:     {Type: field.TypeString, Column: user.FieldName},
			user.FieldEmail:    {Type: field.TypeString, Column: user.FieldEmail},
			user.FieldRole:     {Type: field.TypeInt8, Column: user.FieldRole},
			user.FieldCreated:  {Type: field.TypeTime, Column: user.FieldCreated},
			user.FieldAge:      {Type: field.TypeInt, Column: user.FieldAge},
			user.FieldImageURL: {Type: field.TypeString, Column: user.FieldImageURL},
		},
	}
	graph.MustAddE(
		"operation",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   changes.OperationTable,
			Columns: []string{changes.OperationColumn},
			Bidi:    false,
		},
		"Changes",
		"History",
	)
	graph.MustAddE(
		"changes",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   history.ChangesTable,
			Columns: []string{history.ChangesColumn},
			Bidi:    false,
		},
		"History",
		"Changes",
	)
	graph.MustAddE(
		"author",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   tweet.AuthorTable,
			Columns: []string{tweet.AuthorColumn},
			Bidi:    false,
		},
		"Tweet",
		"User",
	)
	graph.MustAddE(
		"likes",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   tweet.LikesTable,
			Columns: tweet.LikesPrimaryKey,
			Bidi:    false,
		},
		"Tweet",
		"User",
	)
	graph.MustAddE(
		"tweets",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   user.TweetsTable,
			Columns: []string{user.TweetsColumn},
			Bidi:    false,
		},
		"User",
		"Tweet",
	)
	graph.MustAddE(
		"followers",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.FollowersTable,
			Columns: user.FollowersPrimaryKey,
			Bidi:    false,
		},
		"User",
		"User",
	)
	graph.MustAddE(
		"following",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   user.FollowingTable,
			Columns: user.FollowingPrimaryKey,
			Bidi:    false,
		},
		"User",
		"User",
	)
	graph.MustAddE(
		"liked",
		&sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   user.LikedTable,
			Columns: user.LikedPrimaryKey,
			Bidi:    false,
		},
		"User",
		"Tweet",
	)
	return graph
}()

// predicateAdder wraps the addPredicate method.
// All update, update-one and query builders implement this interface.
type predicateAdder interface {
	addPredicate(func(s *sql.Selector))
}

// addPredicate implements the predicateAdder interface.
func (cq *ChangesQuery) addPredicate(pred func(s *sql.Selector)) {
	cq.predicates = append(cq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the ChangesQuery builder.
func (cq *ChangesQuery) Filter() *ChangesFilter {
	return &ChangesFilter{cq}
}

// addPredicate implements the predicateAdder interface.
func (m *ChangesMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the ChangesMutation builder.
func (m *ChangesMutation) Filter() *ChangesFilter {
	return &ChangesFilter{m}
}

// ChangesFilter provides a generic filtering capability at runtime for ChangesQuery.
type ChangesFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *ChangesFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[0].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *ChangesFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(changes.FieldID))
}

// WhereColumn applies the entql string predicate on the column field.
func (f *ChangesFilter) WhereColumn(p entql.StringP) {
	f.Where(p.Field(changes.FieldColumn))
}

// WhereValue applies the entql string predicate on the value field.
func (f *ChangesFilter) WhereValue(p entql.StringP) {
	f.Where(p.Field(changes.FieldValue))
}

// WherePrevious applies the entql string predicate on the previous field.
func (f *ChangesFilter) WherePrevious(p entql.StringP) {
	f.Where(p.Field(changes.FieldPrevious))
}

// WhereHasOperation applies a predicate to check if query has an edge operation.
func (f *ChangesFilter) WhereHasOperation() {
	f.Where(entql.HasEdge("operation"))
}

// WhereHasOperationWith applies a predicate to check if query has an edge operation with a given conditions (other predicates).
func (f *ChangesFilter) WhereHasOperationWith(preds ...predicate.History) {
	f.Where(entql.HasEdgeWith("operation", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (hq *HistoryQuery) addPredicate(pred func(s *sql.Selector)) {
	hq.predicates = append(hq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the HistoryQuery builder.
func (hq *HistoryQuery) Filter() *HistoryFilter {
	return &HistoryFilter{hq}
}

// addPredicate implements the predicateAdder interface.
func (m *HistoryMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the HistoryMutation builder.
func (m *HistoryMutation) Filter() *HistoryFilter {
	return &HistoryFilter{m}
}

// HistoryFilter provides a generic filtering capability at runtime for HistoryQuery.
type HistoryFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *HistoryFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[1].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *HistoryFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(history.FieldID))
}

// WhereEntityName applies the entql string predicate on the entity_name field.
func (f *HistoryFilter) WhereEntityName(p entql.StringP) {
	f.Where(p.Field(history.FieldEntityName))
}

// WhereRecordID applies the entql int predicate on the record_id field.
func (f *HistoryFilter) WhereRecordID(p entql.IntP) {
	f.Where(p.Field(history.FieldRecordID))
}

// WhereTimestamp applies the entql time.Time predicate on the timestamp field.
func (f *HistoryFilter) WhereTimestamp(p entql.TimeP) {
	f.Where(p.Field(history.FieldTimestamp))
}

// WhereAction applies the entql string predicate on the action field.
func (f *HistoryFilter) WhereAction(p entql.StringP) {
	f.Where(p.Field(history.FieldAction))
}

// WhereHasChanges applies a predicate to check if query has an edge changes.
func (f *HistoryFilter) WhereHasChanges() {
	f.Where(entql.HasEdge("changes"))
}

// WhereHasChangesWith applies a predicate to check if query has an edge changes with a given conditions (other predicates).
func (f *HistoryFilter) WhereHasChangesWith(preds ...predicate.Changes) {
	f.Where(entql.HasEdgeWith("changes", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (tq *TweetQuery) addPredicate(pred func(s *sql.Selector)) {
	tq.predicates = append(tq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the TweetQuery builder.
func (tq *TweetQuery) Filter() *TweetFilter {
	return &TweetFilter{tq}
}

// addPredicate implements the predicateAdder interface.
func (m *TweetMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the TweetMutation builder.
func (m *TweetMutation) Filter() *TweetFilter {
	return &TweetFilter{m}
}

// TweetFilter provides a generic filtering capability at runtime for TweetQuery.
type TweetFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *TweetFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[2].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *TweetFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(tweet.FieldID))
}

// WhereContent applies the entql string predicate on the content field.
func (f *TweetFilter) WhereContent(p entql.StringP) {
	f.Where(p.Field(tweet.FieldContent))
}

// WhereCreated applies the entql time.Time predicate on the created field.
func (f *TweetFilter) WhereCreated(p entql.TimeP) {
	f.Where(p.Field(tweet.FieldCreated))
}

// WhereHasAuthor applies a predicate to check if query has an edge author.
func (f *TweetFilter) WhereHasAuthor() {
	f.Where(entql.HasEdge("author"))
}

// WhereHasAuthorWith applies a predicate to check if query has an edge author with a given conditions (other predicates).
func (f *TweetFilter) WhereHasAuthorWith(preds ...predicate.User) {
	f.Where(entql.HasEdgeWith("author", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasLikes applies a predicate to check if query has an edge likes.
func (f *TweetFilter) WhereHasLikes() {
	f.Where(entql.HasEdge("likes"))
}

// WhereHasLikesWith applies a predicate to check if query has an edge likes with a given conditions (other predicates).
func (f *TweetFilter) WhereHasLikesWith(preds ...predicate.User) {
	f.Where(entql.HasEdgeWith("likes", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// addPredicate implements the predicateAdder interface.
func (uq *UserQuery) addPredicate(pred func(s *sql.Selector)) {
	uq.predicates = append(uq.predicates, pred)
}

// Filter returns a Filter implementation to apply filters on the UserQuery builder.
func (uq *UserQuery) Filter() *UserFilter {
	return &UserFilter{uq}
}

// addPredicate implements the predicateAdder interface.
func (m *UserMutation) addPredicate(pred func(s *sql.Selector)) {
	m.predicates = append(m.predicates, pred)
}

// Filter returns an entql.Where implementation to apply filters on the UserMutation builder.
func (m *UserMutation) Filter() *UserFilter {
	return &UserFilter{m}
}

// UserFilter provides a generic filtering capability at runtime for UserQuery.
type UserFilter struct {
	predicateAdder
}

// Where applies the entql predicate on the query filter.
func (f *UserFilter) Where(p entql.P) {
	f.addPredicate(func(s *sql.Selector) {
		if err := schemaGraph.EvalP(schemaGraph.Nodes[3].Type, p, s); err != nil {
			s.AddError(err)
		}
	})
}

// WhereID applies the entql int predicate on the id field.
func (f *UserFilter) WhereID(p entql.IntP) {
	f.Where(p.Field(user.FieldID))
}

// WhereName applies the entql string predicate on the name field.
func (f *UserFilter) WhereName(p entql.StringP) {
	f.Where(p.Field(user.FieldName))
}

// WhereEmail applies the entql string predicate on the email field.
func (f *UserFilter) WhereEmail(p entql.StringP) {
	f.Where(p.Field(user.FieldEmail))
}

// WhereRole applies the entql int8 predicate on the role field.
func (f *UserFilter) WhereRole(p entql.Int8P) {
	f.Where(p.Field(user.FieldRole))
}

// WhereCreated applies the entql time.Time predicate on the created field.
func (f *UserFilter) WhereCreated(p entql.TimeP) {
	f.Where(p.Field(user.FieldCreated))
}

// WhereAge applies the entql int predicate on the age field.
func (f *UserFilter) WhereAge(p entql.IntP) {
	f.Where(p.Field(user.FieldAge))
}

// WhereImageURL applies the entql string predicate on the image_url field.
func (f *UserFilter) WhereImageURL(p entql.StringP) {
	f.Where(p.Field(user.FieldImageURL))
}

// WhereHasTweets applies a predicate to check if query has an edge tweets.
func (f *UserFilter) WhereHasTweets() {
	f.Where(entql.HasEdge("tweets"))
}

// WhereHasTweetsWith applies a predicate to check if query has an edge tweets with a given conditions (other predicates).
func (f *UserFilter) WhereHasTweetsWith(preds ...predicate.Tweet) {
	f.Where(entql.HasEdgeWith("tweets", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasFollowers applies a predicate to check if query has an edge followers.
func (f *UserFilter) WhereHasFollowers() {
	f.Where(entql.HasEdge("followers"))
}

// WhereHasFollowersWith applies a predicate to check if query has an edge followers with a given conditions (other predicates).
func (f *UserFilter) WhereHasFollowersWith(preds ...predicate.User) {
	f.Where(entql.HasEdgeWith("followers", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasFollowing applies a predicate to check if query has an edge following.
func (f *UserFilter) WhereHasFollowing() {
	f.Where(entql.HasEdge("following"))
}

// WhereHasFollowingWith applies a predicate to check if query has an edge following with a given conditions (other predicates).
func (f *UserFilter) WhereHasFollowingWith(preds ...predicate.User) {
	f.Where(entql.HasEdgeWith("following", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}

// WhereHasLiked applies a predicate to check if query has an edge liked.
func (f *UserFilter) WhereHasLiked() {
	f.Where(entql.HasEdge("liked"))
}

// WhereHasLikedWith applies a predicate to check if query has an edge liked with a given conditions (other predicates).
func (f *UserFilter) WhereHasLikedWith(preds ...predicate.Tweet) {
	f.Where(entql.HasEdgeWith("liked", sqlgraph.WrapFunc(func(s *sql.Selector) {
		for _, p := range preds {
			p(s)
		}
	})))
}
