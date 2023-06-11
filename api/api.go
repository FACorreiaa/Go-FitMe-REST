package api

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal/activity"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal/domain"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal/meal"
	"github.com/FACorreiaa/Stay-Healthy-Backend/docs"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	mealRepo := &meal.MealRepository{}
	activityRepo := &activity.ActivityQueries{}

	mealService := &meal.Service{
		Nutrient: meal.NutrientService(mealRepo), // Assign the repository instance to the Service
	}

	activityService := &domain.Service{
		Activity: activityRepo, // Assign the repository instance to the Service
	}

	v1 := app.Group("/api/v1")
	mealService.Routes(v1)
	activityService.Routes(v1)
	docs.SwaggerRoutes(v1)
}
