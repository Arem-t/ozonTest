package postgres

import (
	"OzonTest/internal/graph/model"
	"context"
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() *Repository {
	dsn := "host=" + os.Getenv("DB_HOST") +
		" port=" + os.Getenv("DB_PORT") +
		" user=" + os.Getenv("DB_USER") +
		" dbname=" + os.Getenv("DB_NAME") +
		" password=" + os.Getenv("DB_PASSWORD")

	db, err := gorm.Open("postgres", dsn+" sslmode=disable")
	if err != nil {
		panic("failed to connect database " + err.Error())
	}
	db.AutoMigrate(&model.Post{}, &model.Comment{})
	return &Repository{db: db}
}

func (r *Repository) CreatePost(ctx context.Context, title string, content string) (*model.Post, error) {
	post := &model.Post{Title: title, Content: content}
	if err := r.db.Create(post).Error; err != nil {
		return nil, err
	}
	return post, nil
}

func (r *Repository) CreateComment(ctx context.Context, postID string, parentID *string, content string) (*model.Comment, error) {
	comment := &model.Comment{PostID: postID, Content: content, ParentID: parentID}
	if err := r.db.Create(comment).Error; err != nil {
		return nil, err
	}
	return comment, nil
}

func (r *Repository) GetPosts(ctx context.Context) ([]*model.Post, error) {
	var posts []*model.Post
	if err := r.db.Find(&posts).Error; err != nil {
		return nil, err
	}
	return posts, nil
}

func (r *Repository) GetPost(ctx context.Context, id string) (*model.Post, error) {
	var post model.Post
	if err := r.db.Where("id = ?", id).First(&post).Error; err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *Repository) CommentAdded(ctx context.Context, postID string) (<-chan *model.Comment, error) {
	// Implement subscription logic here
	return nil, errors.New("not implemented")
}
