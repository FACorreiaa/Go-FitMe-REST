package main

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/api"
	"log"

	"github.com/FACorreiaa/Stay-Healthy-Backend/server"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Server initialization
	app := server.Create()

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
