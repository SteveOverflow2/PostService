package rest

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"post-service/pkg/config"
	"post-service/pkg/http/rest/handlers"
	"post-service/pkg/post"

	"github.com/gorilla/mux"
)

type server struct {
	environment string

	Server *http.Server
	Router *mux.Router

	PostService post.PostService
}

const serverLog string = "[Server]: "

func NewServer(version string, environment string, cfg config.HTTPConfig) *server {
	baseUrl := fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)

	s := &server{
		environment: environment,
		Server: &http.Server{
			Addr: baseUrl,

			WriteTimeout: cfg.WriteTimeOut,
			ReadTimeout:  cfg.ReadTimeOut,
			IdleTimeout:  cfg.IdleTimeOut,
		},
		Router: mux.NewRouter(),
	}

	s.Router = s.Router.PathPrefix("/api/").Subrouter()
	log.Println(serverLog+"started api on base url: ", baseUrl+"/api/")

	// Generic routes
	s.Router.NotFoundHandler = http.HandlerFunc(handleNotFound)

	s.Server.Handler = s.Router

	return s
}

func (s *server) Init() {
	s.PostService = post.NewPostService(nil)

	s.routes()
}

func (s *server) Run(name string) {
	var wait time.Duration

	s.Server.Handler = s.Router

	go func() {
		if err := s.Server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	log.Println(serverLog+name, "is running..")

	c := make(chan os.Signal, 1)

	signal.Notify(c, os.Interrupt)
	<-c

	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()

	s.Server.Shutdown(ctx)

	log.Println(serverLog+name, "is shutting down..")

	os.Exit(0)
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Router.ServeHTTP(w, r)
}

func handleNotFound(w http.ResponseWriter, r *http.Request) {
	err := fmt.Errorf("404 - Endpoint was not found")
	handlers.RenderErrorResponse(w, err.Error(), r.URL.Path, err)
}
