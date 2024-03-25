package handlers

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

type Response struct {
	Msg  string
	Code int16
}

func CreateRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTION"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CRSF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Route("/api", func(router chi.Router) {
		router.Get("/healthcheck", HealthCheck)
	})

	return router
}
