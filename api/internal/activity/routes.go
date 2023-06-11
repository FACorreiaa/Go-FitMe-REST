package activity

import "github.com/gofiber/fiber/v2"

func (s *Service) Routes(route fiber.Router) {
	activityHandler := NewHandler(s)
	route.Get("/activities", func(c *fiber.Ctx) error {
		activityHandler.GetActivities(c)
		return nil
	})
}
