package routes

import (
	"fmt"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/grishagavrin/link-shortener/internal/handlers"
	"github.com/grishagavrin/link-shortener/internal/handlers/middlewares"
)

func ServiceRouter() chi.Router {
	r := chi.NewRouter()
	h, err := handlers.New()
	if err != nil {
		fmt.Println("get instance db error")
	}

	r.Use(middleware.Recoverer)
	r.Use(middlewares.GzipMiddleware)
	r.Use(middlewares.CooksMiddleware)
	r.Get("/{id}", h.GetLink)
	r.Post("/", h.SaveTXT)
	r.Post("/api/shorten", h.SaveJSON)
	r.Get("/user/urls", h.GetLinks)
	return r
}
