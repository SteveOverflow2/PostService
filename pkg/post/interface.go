package post

import "context"

//go:generate mockgen --source=post_service.go --destination=./mock/mock_post_service.go
type PostService interface {
	CreatePost(ctx context.Context, post CreatePost) (string, error)
	GetPost(ctx context.Context, uuid string) (Post, error)
	GetAllPosts(ctx context.Context) ([]Post, error)
}

//go:generate mockgen --source=post_repository.go --destination=./mock/mock_post_repository.go
type PostRepository interface {
	CreatePost(ctx context.Context, post CreatePost) error
	GetPost(ctx context.Context, uuid string) (*GetPost, error)
}
