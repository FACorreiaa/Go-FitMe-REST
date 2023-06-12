package server

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/internal/activity/handler"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)

func Register(r chi.Router, lg *logrus.Logger, db *sqlx.DB) {
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	h := handler.NewHandler(lg, db)
	//app.Use(handler.MiddlewareLogger())

	r.Get("/activities", h.GetActivities)
}

//
//func Create() *fiber.App {
//	db, err := db.SetupDatabase()
//	if err != nil {
//		log.Fatal(err)
//		panic("Something happened on creation app")
//	}
//	app := fiber.New(fiber.Config{
//		// Override default error handler
//		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
//			if e, ok := err.(*handlers.Error); ok {
//				return ctx.Status(e.Status).JSON(e)
//			} else if e, ok := err.(*fiber.Error); ok {
//				return ctx.Status(e.Code).JSON(handlers.Error{Status: e.Code, Code: "internal-server", Message: e.Message})
//			} else {
//				return ctx.Status(500).JSON(handlers.Error{Status: 500, Code: "internal-server", Message: err.Error()})
//			}
//		},
//	})
//	app.Get("/swagger/*", swagger.HandlerDefault) // default
//
//	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
//		URL:         "http://example.com/doc.json",
//		DeepLinking: false,
//		// Expand ("list") or Collapse ("none") tag groups by default
//		DocExpansion: "none",
//		// Prefill OAuth ClientId on Authorize popup
//		OAuth: &swagger.OAuthConfig{
//			AppName:  "OAuth Provider",
//			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
//		},
//		// Ability to change OAuth2 redirect uri location
//		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
//	}))
//
//	Register(app, db)
//
//	app.Get("/", func(c *fiber.Ctx) error {
//		return c.SendString("OK")
//	})
//
//	return app
//}
//
//func Listen(app *fiber.App) error {
//
//	// 404 Handler
//	app.Use(func(c *fiber.Ctx) error {
//		return c.SendStatus(404)
//	})
//
//	serverHost := os.Getenv("SERVER_HOST")
//	serverPort := os.Getenv("SERVER_PORT")
//
//	return app.Listen(fmt.Sprintf("%s:%s", serverHost, serverPort))
//}
