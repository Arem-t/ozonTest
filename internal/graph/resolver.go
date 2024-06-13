package graph

import (
	"OzonTest/internal/graph/generated"
	"OzonTest/internal/graph/model"
	"context"
	"errors"
)

type Resolver struct {
	Repo Repository
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

func (r *Resolver) Subscription() generated.SubscriptionResolver {
	return &subscriptionResolver{r}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }

func (r *mutationResolver) CreatePost(ctx context.Context, title string, content string) (*model.Post, error) {
	return r.Repo.CreatePost(ctx, title, content)
}

func (r *mutationResolver) CreateComment(ctx context.Context, postID string, parentID *string, content string) (*model.Comment, error) {
	if len(content) > 2000 {
		return nil, errors.New("comment is too long")
	}
	return r.Repo.CreateComment(ctx, postID, parentID, content)
}

func (r *queryResolver) Posts(ctx context.Context) ([]*model.Post, error) {
	return r.Repo.GetPosts(ctx)
}

func (r *queryResolver) Post(ctx context.Context, id string) (*model.Post, error) {
	return r.Repo.GetPost(ctx, id)
}

func (r *subscriptionResolver) CommentAdded(ctx context.Context, postID string) (<-chan *model.Comment, error) {
	return r.Repo.CommentAdded(ctx, postID)
}
