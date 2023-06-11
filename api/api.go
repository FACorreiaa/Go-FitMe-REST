package api

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal/meal"
	"github.com/FACorreiaa/Stay-Healthy-Backend/docs"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	repo := &meal.MealRepository{}

	service := &meal.Service{
		Nutrient: meal.NutrientService(repo), // Assign the repository instance to the Service
	}

	v1 := app.Group("/api/v1")
	service.Routes(v1)
	docs.SwaggerRoutes(v1)
}
