package meal

import "github.com/gofiber/fiber/v2"

func (m *Service) Routes(route fiber.Router) {
	mealHandler := NewHandler(m)
	route.Get("/books", func(c *fiber.Ctx) error {
		mealHandler.GetNutrients(c)
		return nil
	})
}
