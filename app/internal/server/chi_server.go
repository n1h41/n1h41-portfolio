package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/n1h41/portfolio/internal/delivery/http/route"
)

type ChiServer struct{}

func (s ChiServer) Run() {
	r := chi.NewRouter()

	// INFO: Register middlewares
	r.Use(middleware.Logger)

	// INFO: Initialize routes
	route.RegisterRoutes(r)

	http.ListenAndServe(":3000", r)
}
