package api

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal/activity"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal/domain"
	"github.com/FACorreiaa/Stay-Healthy-Backend/docs"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	repo := domain.Repository{
		Activity: &activity.ActivityService{},
	}

	//activityHandler := activity.NewActivityService(activity.NewActivityService{})
	v1 := app.Group("/api/v1")
	//app.Route("/api/v1/activities", func(route fiber.Router) {
	//	route.Get("/", func(c *fiber.Ctx) error {
	//		activityHandler.GetActivities(c)
	//		return nil
	//	})
	//})
	activity.ActivityRoutes(v1, activity.NewActivityService(&repo))
	docs.SwaggerRoutes(v1)
	//activity.ActivityRoutes(v1)
}
