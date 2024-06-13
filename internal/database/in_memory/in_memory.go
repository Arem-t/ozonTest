package in_memory

import (
	"OzonTest/internal/graph/model"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"
)

type Repository struct {
	posts    map[string]*model.Post
	comments map[string]*model.Comment
	mu       sync.Mutex
}

func NewRepository() *Repository {
	return &Repository{
		posts:    make(map[string]*model.Post),
		comments: make(map[string]*model.Comment),
	}
}

func (r *Repository) CreatePost(ctx context.Context, title string, content string) (*model.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	post := &model.Post{
		ID:       generateID(),
		Title:    title,
		Content:  content,
		Comments: []*model.Comment{},
	}
	r.posts[post.ID] = post
	return post, nil
}

func (r *Repository) CreateComment(ctx context.Context, postID string, parentID *string, content string) (*model.Comment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	comment := &model.Comment{
		ID:       generateID(),
		PostID:   postID,
		Content:  content,
		Children: []*model.Comment{},
	}
	r.comments[comment.ID] = comment

	if parentID != nil {
		parentComment, ok := r.comments[*parentID]
		if ok {
			parentComment.Children = append(parentComment.Children, comment)
		}
	} else {
		post, ok := r.posts[postID]
		if ok {
			post.Comments = append(post.Comments, comment)
		}
	}
	return comment, nil
}

func (r *Repository) GetPosts(ctx context.Context) ([]*model.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	posts := make([]*model.Post, 0, len(r.posts))
	for _, post := range r.posts {
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *Repository) GetPost(ctx context.Context, id string) (*model.Post, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	post, ok := r.posts[id]
	if !ok {
		return nil, errors.New("post not found")
	}
	return post, nil
}

func (r *Repository) CommentAdded(ctx context.Context, postID string) (<-chan *model.Comment, error) {
	// Implement subscription logic here
	return nil, errors.New("not implemented")
}

// generateID generates a unique ID for posts and comments.
func generateID() string {
	return fmt.Sprintf("%d", time.Now().UnixNano())
}
