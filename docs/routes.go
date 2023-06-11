package docs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
)

func SwaggerRoutes(router fiber.Router) {
	router.Get("/swagger/*", swagger.HandlerDefault) // default

	router.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		//OAuth: &swagger.OAuthConfig{
		//	AppName:  "OAuth Provider",
		//	ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		//},
		// Ability to change OAuth2 redirect uri location
		//OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))
}
