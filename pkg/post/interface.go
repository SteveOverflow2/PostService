package post

import "context"

//go:generate mockgen --source=post_service.go --destination=./mock/mock_post_service.go
type PostService interface {
	CreatePost(ctx context.Context, post CreatePost) (string, error)
	GetPost(ctx context.Context, uuid string) (*Post, error)
	DeletePost(ctx context.Context, uuid string)
	GetAllPosts(ctx context.Context) ([]Post, error)
	UpdateTime(ctx context.Context, uuid string)
}

//go:generate mockgen --source=post_repository.go --destination=./mock/mock_post_repository.go
type PostRepository interface {
	CreatePost(ctx context.Context, post CreatePost) (string, error)
	GetAllPosts(ctx context.Context) ([]Post, error)
	GetPost(ctx context.Context, uuid string) (*Post, error)
	UpdateTime(ctx context.Context, uuid string)
	DeletePost(ctx context.Context, uuid string)
}
