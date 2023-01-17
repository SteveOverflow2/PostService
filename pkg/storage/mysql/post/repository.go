package post

import (
	"context"
	"database/sql"
	"fmt"
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
	createPostQuery := "INSERT INTO post (_id, title, description, createdAt, updatedAt, views, answers, votes, poster) VALUES (?,?,?,?,?,?,?,?,?);"
	stmt, err := p.db.Prepare(createPostQuery)
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	newId := uuid.New().String()
	_, err = stmt.Exec(newId, post.Title, post.Body, time.Now().Unix(), time.Now().Unix(), 0, 0, 0, post.Subject)
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
		err := result.Scan(&post.Id, &post.Title, &post.Description, &post.CreatedAt, &post.UpdatedAt, &post.Views, &post.Answers, &post.Votes, &post.Poster)
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
		err := result.Scan(&post.Id, &post.Title, &post.Description, &post.CreatedAt, &post.UpdatedAt, &post.Views, &post.Answers, &post.Votes, &post.Poster)
		if err != nil {
			return nil, err
		}
	}
	return &post, nil
}

func (p *postRepository) UpdateTime(ctx context.Context, uuid string) {
	updateTime := "UPDATE post SET updatedAt = ? WHERE _id = ?"
	fmt.Printf("uuid: %v\n", uuid)
	fmt.Printf("\" starting update\": %v\n", " starting update")
	stmt, err := p.db.Prepare(updateTime)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(time.Now().Unix(), uuid)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("\" succes\": %v\n", " succes")
	return
}

func (p *postRepository) DeletePost(ctx context.Context, uuid string) {
	updateTime := "DELETE FROM post WHERE _id = ?"
	stmt, err := p.db.Prepare(updateTime)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	defer stmt.Close()
	_, err = stmt.Exec(uuid)
	if err != nil {
		fmt.Printf("err: %v\n", err)
		return
	}
	fmt.Printf("\" succes\": %v\n", " succes")
	return
}
