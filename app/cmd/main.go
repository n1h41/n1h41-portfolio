package main

import (
	"log"
	"net/http"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/n1h41/portfolio/views/home"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.NoCache)
	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("public"))))
	homeView := home.Home()
	homeViewHandler := templ.Handler(homeView)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		homeViewHandler.ServeHTTP(w, r)
	})
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Panic(err)
	}
	log.Println("Server running at localhost:3000")
}
