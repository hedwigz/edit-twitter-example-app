package edit_twitter_example_app

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/ariga/edit-twitter-example-app/ent"
	"github.com/ariga/edit-twitter-example-app/ent/user"
)

func (r *mutationResolver) CreateTweet(ctx context.Context, tweet TweetInput) (*ent.Tweet, error) {
	client := ent.FromContext(ctx)
	t, err := client.Tweet.Create().
		SetContent(tweet.Content).
		Save(ctx)
	return t, err
}

func (r *mutationResolver) LikeTweet(ctx context.Context, id int) (*ent.Tweet, error) {
	client := ent.FromContext(ctx)
	uid := ctx.Value("userid").(int)
	t, err := client.Tweet.UpdateOneID(id).AddLikeIDs(uid).Save(ctx)

	return t, err
}

func (r *mutationResolver) UpdateTweet(ctx context.Context, id int, tweet TweetUpdateInput) (*ent.Tweet, error) {
	client := ent.FromContext(ctx)
	t, err := client.Tweet.UpdateOneID(id).SetContent(*tweet.Content).Save(ctx)
	return t, err
}

func (r *queryResolver) Tweets(ctx context.Context) ([]*ent.Tweet, error) {
	return r.client.Tweet.Query().
		All(ctx)
}

func (r *queryResolver) Tweet(ctx context.Context, id int) (*ent.Tweet, error) {
	return r.client.Tweet.Get(ctx, id)
}

func (r *queryResolver) Me(ctx context.Context) (*ent.User, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Node(ctx context.Context, id int) (ent.Noder, error) {
	client := ent.FromContext(ctx)
	return client.Noder(ctx, id)
}

func (r *queryResolver) Nodes(ctx context.Context, ids []int) ([]ent.Noder, error) {
	client := ent.FromContext(ctx)
	return client.Noders(ctx, ids)
}

func (r *queryResolver) TweetsX(ctx context.Context, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TweetOrder) (*ent.TweetConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) TweetEditHistory(ctx context.Context, id int) ([]*TweetEdit, error) {
	client := ent.FromContext(ctx)
	if client == nil {
		client = r.client
	}
	t, err := client.Tweet.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	// This code can be generated in the future
	hs, err := client.Tweet.QueryHistory(t).WithChanges().All(ctx)
	if err != nil {
		return nil, err
	}
	var tes []*TweetEdit
	for i := 0; i < len(hs); i++ {
		h := hs[i]
		for j := 0; j < len(h.Edges.Changes); j++ {
			c := h.Edges.Changes[j]
			tes = append(tes, &TweetEdit{
				EditedAt: h.Timestamp,
				Editor:   nil,
				Diff:     fmt.Sprintf("%s -> %s", t.Content, c.Value),
			})
		}
	}
	return tes, nil
}

func (r *tweetResolver) History(ctx context.Context, obj *ent.Tweet) ([]*TweetEdit, error) {
	client := ent.FromContext(ctx)
	if client == nil {
		client = r.client
	}
	hs, err := client.Tweet.QueryHistory(obj).WithChanges().All(ctx)
	if err != nil {
		return nil, err
	}
	var tes []*TweetEdit
	for i := 0; i < len(hs); i++ {
		h := hs[i]
		for j := 0; j < len(h.Edges.Changes); j++ {
			c := h.Edges.Changes[j]
			tes = append(tes, &TweetEdit{
				ID:       c.ID,
				EditedAt: h.Timestamp,
				Editor:   nil,
				Diff:     c.Value,
			})
		}
	}
	return tes, nil
}

func (r *tweetResolver) LikesCount(ctx context.Context, obj *ent.Tweet) (int, error) {
	client := ent.FromContext(ctx)
	if client == nil {
		client = r.client
	}
	return client.Tweet.QueryLikes(obj).Count(ctx)
}

func (r *tweetResolver) UserLiked(ctx context.Context, obj *ent.Tweet, id int) (bool, error) {
	return r.client.Tweet.QueryLikes(obj).Where(user.ID(id)).Exist(ctx)
}

func (r *tweetResolver) LikedUsers(ctx context.Context, obj *ent.Tweet, after *ent.Cursor, first *int, before *ent.Cursor, last *int, orderBy *ent.TweetOrder) (*ent.UserConnection, error) {
	return r.client.Tweet.QueryLikes(obj).Paginate(ctx, after, first, before, last)
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Tweet returns TweetResolver implementation.
func (r *Resolver) Tweet() TweetResolver { return &tweetResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type tweetResolver struct{ *Resolver }
