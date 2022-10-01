package rest

import (
	"net/http"

	post "post-service/pkg/http/rest/handlers/post"
)

func (s *server) routes() {
	// Example handlers subset
	postHandler := s.Router.PathPrefix("/post").Subrouter()

	//POST
	postHandler.HandleFunc("", post.CreatePostHandler(s.PostService)).Methods(http.MethodPost)

	//GET
	postHandler.HandleFunc("/getAll", post.GetAllPostsHandler(s.PostService)).Methods(http.MethodGet)
	postHandler.HandleFunc("/{uuid}", post.GetPostHandler(s.PostService)).Methods(http.MethodGet)
}
