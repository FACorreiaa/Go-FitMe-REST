package api

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/docs"
	"github.com/go-chi/chi/v5"
)

func Setup() {
	r := chi.NewRouter()
	docs.SwaggerRoutes(r)
}
