package post

import (
	"encoding/json"
	"net/http"

	"post-service/pkg/http/rest/handlers"
	"post-service/pkg/post"
	"post-service/pkg/util"

	"github.com/gorilla/mux"
)

func CreatePostHandler(postService post.PostService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var post post.CreatePost

		if err := json.NewDecoder(r.Body).Decode(&post); err != nil {
			handlers.RenderErrorResponse(w, "Invalid request payload", r.URL.Path, util.WrapErrorf(err, util.ErrorCodeInvalid, "json decoder"))
			return
		}

		postService.CreatePost(r.Context(), post)
	}
}

func GetPostHandler(postService post.PostService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		uuid := mux.Vars(r)["uuid"]
		if len(uuid) == 0 {
			err := util.NewErrorf(util.ErrorCodeInternal, "Query parameters are invalid")
			handlers.RenderErrorResponse(w, err.Error(), r.URL.Path, err)
			return
		}

		post, err := postService.GetPost(r.Context(), uuid)
		if err != nil {
			handlers.RenderErrorResponse(w, "internal server error", r.URL.Path, util.WrapErrorf(err, util.ErrorCodeInternal, err.Error()))
		}

		handlers.RenderResponse(w, http.StatusOK, post)
	}
}

func GetAllPostsHandler(postService post.PostService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		enableCors(&w)

		post, err := postService.GetAllPosts(r.Context())
		if err != nil {
			handlers.RenderErrorResponse(w, "internal server error", r.URL.Path, util.WrapErrorf(err, util.ErrorCodeInternal, err.Error()))
		}

		handlers.RenderResponse(w, http.StatusOK, post)
	}
}

func enableCors(w *http.ResponseWriter) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
}
