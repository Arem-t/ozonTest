package graph

import (
	"OzonTest/internal/graph/model"
	"context"
)

type Repository interface {
	CreatePost(ctx context.Context, title string, content string) (*model.Post, error)
	CreateComment(ctx context.Context, postID string, parentID *string, content string) (*model.Comment, error)
	GetPosts(ctx context.Context) ([]*model.Post, error)
	GetPost(ctx context.Context, id string) (*model.Post, error)
	CommentAdded(ctx context.Context, postID string) (<-chan *model.Comment, error)
}
