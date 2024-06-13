package graph

import (
	"OzonTest/internal/graph/model"
	"context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

type MockRepository struct {
	mock.Mock
}

func (m *MockRepository) CreatePost(ctx context.Context, title string, content string) (*model.Post, error) {
	args := m.Called(ctx, title, content)
	return args.Get(0).(*model.Post), args.Error(1)
}

func (m *MockRepository) CreateComment(ctx context.Context, postID string, parentID *string, content string) (*model.Comment, error) {
	args := m.Called(ctx, postID, parentID, content)
	return args.Get(0).(*model.Comment), args.Error(1)
}

func (m *MockRepository) GetPosts(ctx context.Context) ([]*model.Post, error) {
	args := m.Called(ctx)
	return args.Get(0).([]*model.Post), args.Error(1)
}

func (m *MockRepository) GetPost(ctx context.Context, id string) (*model.Post, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*model.Post), args.Error(1)
}

func (m *MockRepository) CommentAdded(ctx context.Context, postID string) (<-chan *model.Comment, error) {
	args := m.Called(ctx, postID)
	return args.Get(0).(<-chan *model.Comment), args.Error(1)
}

func TestCreatePost(t *testing.T) {
	mockRepo := new(MockRepository)
	resolver := &Resolver{Repo: mockRepo}

	expectedPost := &model.Post{ID: "1", Title: "Test Title", Content: "Test Content"}
	mockRepo.On("CreatePost", mock.Anything, "Test Title", "Test Content").Return(expectedPost, nil)

	post, err := resolver.Mutation().CreatePost(context.Background(), "Test Title", "Test Content")
	assert.NoError(t, err)
	assert.Equal(t, expectedPost, post)
	mockRepo.AssertExpectations(t)
}

func TestCreateComment(t *testing.T) {
	mockRepo := new(MockRepository)
	resolver := &Resolver{Repo: mockRepo}

	expectedComment := &model.Comment{ID: "1", Content: "Test Comment"}
	mockRepo.On("CreateComment", mock.Anything, "1", (*string)(nil), "Test Comment").Return(expectedComment, nil)

	comment, err := resolver.Mutation().CreateComment(context.Background(), "1", nil, "Test Comment")
	assert.NoError(t, err)
	assert.Equal(t, expectedComment, comment)
	mockRepo.AssertExpectations(t)
}

func TestPosts(t *testing.T) {
	mockRepo := new(MockRepository)
	resolver := &Resolver{Repo: mockRepo}

	expectedPosts := []*model.Post{{ID: "1", Title: "Test Title", Content: "Test Content"}}
	mockRepo.On("GetPosts", mock.Anything).Return(expectedPosts, nil)

	posts, err := resolver.Query().Posts(context.Background())
	assert.NoError(t, err)
	assert.Equal(t, expectedPosts, posts)
	mockRepo.AssertExpectations(t)
}

func TestPost(t *testing.T) {
	mockRepo := new(MockRepository)
	resolver := &Resolver{Repo: mockRepo}

	expectedPost := &model.Post{ID: "1", Title: "Test Title", Content: "Test Content"}
	mockRepo.On("GetPost", mock.Anything, "1").Return(expectedPost, nil)

	post, err := resolver.Query().Post(context.Background(), "1")
	assert.NoError(t, err)
	assert.Equal(t, expectedPost, post)
	mockRepo.AssertExpectations(t)
}
