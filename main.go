package main

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/FACorreiaa/Stay-Healthy-Backend/server/internals"
	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
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
		fmt.Println(err)
		log.Fatalf("%+v", err)
	}
}

func run(ctx context.Context) error {
	err := godotenv.Load(".env")

	if os.Getenv("APP_ENV") == "dev" {
		if err != nil {
			fmt.Println(err)
			log.Fatal("Error loading .env file")
		}
	}

	server, err := internals.NewServer()
	//println("err %+v", err)
	if err != nil {
		server.Close()
		fmt.Println(err)
		//log.Fatalf("%+v", err)
		return err
	}

	err = server.Run(ctx)
	defer server.Close()
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
func HealthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	res := map[string]interface{}{
		"data": "Server is up and running",
	}

	err := json.NewEncoder(w).Encode(res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
