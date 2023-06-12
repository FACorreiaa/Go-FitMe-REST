package activity

import "github.com/gofiber/fiber/v2"

func ActivityRoutes(route fiber.Router, service *ActivityService) {
	activityHandler := NewHandler(service)
	route.Get("/activities", func(c *fiber.Ctx) error {
		activityHandler.GetActivities(c)
		return nil
	})
}
