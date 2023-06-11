package main

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/api"
	"github.com/gofiber/fiber/v2"
	"log"

	_ "github.com/FACorreiaa/Stay-Healthy-Backend/docs"
	"github.com/FACorreiaa/Stay-Healthy-Backend/server"
	"github.com/joho/godotenv"
)

// @title Fiber Swagger Example API
// @version 2.0
// @description This is a sample server server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:3000
// @BasePath /
// @schemes http

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Server initialization
	app := server.Create()
	app.Get("/", HealthCheck)

	//// Migrations
	//db.DB.AutoMigrate(&books.Book{})
	//if err != nil {
	//	log.Fatal("Error doing migration")
	//}
	// Api routes
	api.Setup(app)

	if err := server.Listen(app); err != nil {
		log.Panic(err)
	}
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
func HealthCheck(c *fiber.Ctx) error {
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	if err := c.JSON(res); err != nil {
		return err
	}

	return nil
}
