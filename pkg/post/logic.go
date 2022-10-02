package post

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
)

var (
	AllPosts []Post
)

type postService struct {
	postRepository PostRepository
}

func NewPostService(u PostRepository) PostService {
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
		postRepository: u,
	}
}

func (u *postService) CreatePost(ctx context.Context, post CreatePost) (string, error) {
	AllPosts = append(AllPosts, Post{
		Id:          uuid.New().String(),
		Title:       post.Title,
		Description: post.Body,
		Poster:      post.Subject,
		CreatedAt:   time.Now().Unix(),
		Views:       0,
		Votes:       0,
	})
	return AllPosts[len(AllPosts)-1].Id, nil
}

func (u *postService) GetPost(ctx context.Context, uuid string) (Post, error) {

	for i := 0; i < len(AllPosts); i++ {
		if uuid == AllPosts[i].Id {
			return AllPosts[i], nil
		}
	}

	return Post{}, errors.New("error")
}

func (u *postService) GetAllPosts(ctx context.Context) ([]Post, error) {
	return AllPosts, nil
}
