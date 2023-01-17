package post

import (
	"context"
	"fmt"
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
	fmt.Printf("\" creating\": %v\n", " creating")
	str, err := p.postRepository.CreatePost(ctx, post)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return "", err
	}
	fmt.Printf("str: %v\n", str)
	return str, nil
}

func (p *postService) GetPost(ctx context.Context, uuid string) (*Post, error) {
	return p.postRepository.GetPost(ctx, uuid)
}

func (p *postService) GetAllPosts(ctx context.Context) ([]Post, error) {
	return p.postRepository.GetAllPosts(ctx)
}

func (p *postService) UpdateTime(ctx context.Context, uuid string) {
	p.postRepository.UpdateTime(ctx, uuid)
}
func (p *postService) DeletePost(ctx context.Context, uuid string) {
	p.postRepository.DeletePost(ctx, uuid)
}
