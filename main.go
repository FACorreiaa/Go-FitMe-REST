package main

import (
	"context"
	"github.com/FACorreiaa/Stay-Healthy-Backend/server/internals"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"log"
)

// @title StayHealthy Swagger Documentation
// @version 2.0
// @description Alpha server built with Go and Chi
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /
// @schemes http
func main() {
	if err := run(context.Background()); err != nil {
		println(err)
		log.Fatalf("%+v", err)
	}
}

func run(ctx context.Context) error {
	err := godotenv.Load()
	if err != nil {
		println(err)
		log.Fatal("Error loading .env file")
	}

	server, err := internals.NewServer()

	if err != nil {
		println(err)
		return err
	}

	err = server.Run(ctx)

	return err
}

// HealthCheck godoc
// @Summary Show the status of server.
// @Description get the status of server.
// @Tags root
// @Accept */*
// @Produce json
// @Success 200 {object} map[string]interface{}
// @Router / [get]
//func HealthCheck(c *fiber.Ctx) error {
//	res := map[string]interface{}{
//		"data": "Server is up and running",
//	}
//
//	if err := c.JSON(res); err != nil {
//		return err
//	}
//
//	return nil
//}
