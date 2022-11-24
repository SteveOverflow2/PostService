package post

import (
	"context"
	"database/sql"
	"post-service/pkg/post"
	"time"

	"github.com/google/uuid"
)

type postRepository struct {
	db *sql.DB
}

func NewPostRepository(ctx context.Context, db *sql.DB) post.PostRepository {
	return &postRepository{
		db: db,
	}
}

func (p *postRepository) CreatePost(ctx context.Context, post post.CreatePost) (string, error) {
	createPostQuery := "INSERT INTO post (_id, title, description, createdAt, views, answers, votes, poster) VALUES (?,?,?,?,?,?,?,?);"
	stmt, err := p.db.Prepare(createPostQuery)
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	newId := uuid.New().String()
	_, err = stmt.Exec(newId, post.Title, post.Body, time.Now().Unix(), 0, 0, 0, post.Subject)
	if err != nil {
		return "", err
	}
	return newId, nil
}

func (p *postRepository) GetAllPosts(ctx context.Context) ([]post.Post, error) {
	getAllPosts := "SELECT * FROM post;"

	var posts []post.Post
	result, err := p.db.Query(getAllPosts)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	for result.Next() {
		var post post.Post
		err := result.Scan(&post.Id, &post.Title, &post.Description, &post.CreatedAt, &post.Views, &post.Answers, &post.Votes, &post.Poster)
		if err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}
func (p *postRepository) GetPost(ctx context.Context, uuid string) (*post.Post, error) {
	getAllPosts := "SELECT * FROM post WHERE _id = ?;"

	result, err := p.db.Query(getAllPosts, uuid)
	if err != nil {
		panic(err.Error())
	}
	defer result.Close()
	var post post.Post

	for result.Next() {
		err := result.Scan(&post.Id, &post.Title, &post.Description, &post.CreatedAt, &post.Views, &post.Answers, &post.Votes, &post.Poster)
		if err != nil {
			return nil, err
		}
	}
	return &post, nil
}
