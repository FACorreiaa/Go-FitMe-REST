package activity

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal/domain"
	"github.com/gofiber/fiber/v2"
)

//type ActivityService struct {
//	repo *ActivityRepository
//}

/*****************
** AIRPORT  **
******************/

type Service struct {
	repo *domain.Repository
}

func (s *Service) GetActivities(c *fiber.Ctx) ([]domain.Activity, error) {
	return s.repo.Activity.GetAll(c)
}
