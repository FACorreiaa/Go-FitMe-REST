package activity

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal/domain"
	"github.com/gofiber/fiber/v2"
)

type ActivityService struct {
	repo domain.Repository
}

func NewActivityService(repo *domain.Repository) *ActivityService {
	return &ActivityService{
		repo: *repo,
	}
}

func (a *ActivityService) GetAll(c *fiber.Ctx) ([]domain.Activity, error) {
	return a.repo.Activity.GetAll(c)
}
