package post

import (
	"context"
)

type postService struct {
	postRepository PostRepository
}

func NewPostService(p PostRepository) PostService {
	return &postService{
		postRepository: p,
	}
}

func (p *postService) CreatePost(ctx context.Context, post CreatePost) (string, error) {

	return p.postRepository.CreatePost(ctx, post)
}

func (p *postService) GetPost(ctx context.Context, uuid string) (*Post, error) {
	return p.postRepository.GetPost(ctx, uuid)
}

func (p *postService) GetAllPosts(ctx context.Context) ([]Post, error) {
	return p.postRepository.GetAllPosts(ctx)
}
