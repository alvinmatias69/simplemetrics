package internal

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Server struct {
}

func New() *Server {
	return &Server{}
}

func (s *Server) Start() {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(middleware.NoCache)

	router.Get("/metrics", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Write([]byte("{\"hello\": \"world\"}"))
	})
	http.ListenAndServe(":3000", router)
}
