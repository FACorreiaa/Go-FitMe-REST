package docs

import (
	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger"
)

func SwaggerRoutes(router *chi.Mux) {

	router.Get("/api/v1/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("./doc.json"), // Use a relative URL for the Swagger documentation file
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("#swagger-ui"),
	))
}
