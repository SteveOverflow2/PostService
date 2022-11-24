package post

import (
	"context"
	"time"
)

var (
	AllPosts []Post
)

type postService struct {
	postRepository PostRepository
}

func NewPostService(p PostRepository) PostService {
	AllPosts = append(AllPosts, Post{
		Id:          "123",
		Title:       "My first post",
		Description: "My first description",
		Poster:      "aiguagha",
		CreatedAt:   time.Now().Unix(),
		Views:       19,
		Votes:       9,
	})
	AllPosts = append(AllPosts, Post{
		Id:          "456",
		Title:       "My second post",
		Description: "My second description",
		Poster:      "26286",
		CreatedAt:   time.Now().Unix(),
		Views:       4,
		Votes:       2,
	})
	AllPosts = append(AllPosts, Post{
		Id:          "789",
		Title:       "My third post",
		Description: "My third description",
		Poster:      "9999",
		CreatedAt:   time.Now().Unix(),
		Views:       7,
		Votes:       0,
	})
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
