package domain

import "github.com/gofiber/fiber/v2"

type Activity struct {
	ID              int     `json:"id" pg:"default:gen_random_uuid()"`
	Name            string  `json:"name"`
	CaloriesPerHour float32 `json:"calories_per_hour"`
	DurationMinutes float32 `json:"duration_minutes"`
	TotalCalories   float32 `json:"total_calories"`
}

type Service struct {
	Activity ActivityService
}

type ActivityService interface {
	GetAll(c *fiber.Ctx) ([]Activity, error)
}

type Repository struct {
	Activity ActivityRepository
}

type ActivityRepository interface {
	GetAll(c *fiber.Ctx) ([]Activity, error)
}
