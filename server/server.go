package server

import (
	"fmt"
	"github.com/FACorreiaa/Stay-Healthy-Backend/db"
	"github.com/FACorreiaa/Stay-Healthy-Backend/handlers"
	"github.com/gofiber/fiber/v2/middleware/csrf"
	"github.com/gofiber/fiber/v2/utils"
	"github.com/gofiber/swagger"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/etag"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/helmet/v2"
)

func setupMiddlewares(app *fiber.App) {
	app.Use(helmet.New())
	app.Use(recover.New())
	app.Use(cors.New())
	csrfConfig := csrf.Config{
		KeyLookup:      "header:X-Csrf-Token", // string in the form of '<source>:<key>' that is used to extract token from the request
		CookieName:     "my_csrf_",            // name of the session cookie
		CookieSameSite: "Strict",              // indicates if CSRF cookie is requested by SameSite
		Expiration:     3 * time.Hour,         // expiration is the duration before CSRF token will expire
		KeyGenerator:   utils.UUID,            // creates a new CSRF token
	}
	app.Use(csrf.New(csrfConfig))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed, // 1
	}))
	app.Use(etag.New())

	if os.Getenv("ENABLE_LIMITER") != "" {
		app.Use(limiter.New())
	}

	if os.Getenv("ENABLE_LOGGER") != "" {
		app.Use(logger.New(logger.Config{
			// For more options, see the Config section
			Next:         nil,
			Done:         nil,
			Format:       "[${time}] [${ip}]:${port} ${status} - ${method} ${path}\n\n",
			TimeFormat:   "15:04:05",
			TimeZone:     "Local",
			TimeInterval: 500 * time.Millisecond,
			Output:       os.Stdout,
		}))
	}
}

func Create() *fiber.App {
	db.NewPostgres(db.Config{})

	app := fiber.New(fiber.Config{
		// Override default error handler
		ErrorHandler: func(ctx *fiber.Ctx, err error) error {
			if e, ok := err.(*handlers.Error); ok {
				return ctx.Status(e.Status).JSON(e)
			} else if e, ok := err.(*fiber.Error); ok {
				return ctx.Status(e.Code).JSON(handlers.Error{Status: e.Code, Code: "internal-server", Message: e.Message})
			} else {
				return ctx.Status(500).JSON(handlers.Error{Status: 500, Code: "internal-server", Message: err.Error()})
			}
		},
	})
	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Get("/swagger/*", swagger.New(swagger.Config{ // custom
		URL:         "http://example.com/doc.json",
		DeepLinking: false,
		// Expand ("list") or Collapse ("none") tag groups by default
		DocExpansion: "none",
		// Prefill OAuth ClientId on Authorize popup
		OAuth: &swagger.OAuthConfig{
			AppName:  "OAuth Provider",
			ClientId: "21bb4edc-05a7-4afc-86f1-2e151e4ba6e2",
		},
		// Ability to change OAuth2 redirect uri location
		OAuth2RedirectUrl: "http://localhost:8080/swagger/oauth2-redirect.html",
	}))

	setupMiddlewares(app)

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("OK")
	})

	return app
}

func Listen(app *fiber.App) error {

	// 404 Handler
	app.Use(func(c *fiber.Ctx) error {
		return c.SendStatus(404)
	})

	serverHost := os.Getenv("SERVER_HOST")
	serverPort := os.Getenv("SERVER_PORT")

	return app.Listen(fmt.Sprintf("%s:%s", serverHost, serverPort))
}
