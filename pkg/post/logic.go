package post

import (
	"context"

	"github.com/google/uuid"
)

type postService struct {
	postRepository PostRepository
}

func NewPostService(u PostRepository) PostService {
	return &postService{
		postRepository: u,
	}
}

func (u *postService) CreatePost(ctx context.Context, post CreatePost) error {
	post.Uuid = uuid.New().String()
	// Extra logic

	return nil
}

func (u *postService) GetPost(ctx context.Context, uuid string) (*GetPost, error) {
	post, err := u.postRepository.GetPost(ctx, uuid)
	if err != nil {
		return nil, err
	}
	// Extra logic

	return post, nil
}

func (u *postService) GetAllPosts(ctx context.Context) ([]Post, error) {

	//Here we should call the repository

	var allPosts []Post

	allPosts = append(allPosts, Post{
		Id:          "123",
		Title:       "My first post",
		Description: "My first description",
	})
	allPosts = append(allPosts, Post{
		Id:          "456",
		Title:       "My second post",
		Description: "My second description",
	})
	allPosts = append(allPosts, Post{
		Id:          "789",
		Title:       "My third post",
		Description: "My third description",
	})
	return allPosts, nil
}
