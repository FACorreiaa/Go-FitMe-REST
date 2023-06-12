package activity

import (
	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	service *ActivityService
}

func NewHandler(service *ActivityService) *Handler {
	return &Handler{service: service}
}

// GetActivities func gets all existing activities
// @Description Get all activities
// @Summary Get all activities
// @Tags activities
// @Accept json
// @Produce json
// @Success 200 {array} domain.Activity
// @Router /api/v1/activity [get]
func (h *Handler) GetActivities(c *fiber.Ctx) error {
	activities, err := h.service.repo.Activity.GetAll(c)
	if err != nil {
		// Return, if activities not found.
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": true,
			"msg":   "activities were not found",
			"count": 0,
			"books": nil,
		})
	}

	// Return status 200 OK.
	return c.JSON(fiber.Map{
		"error": false,
		"msg":   nil,
		"count": len(activities),
		"books": activities,
	})
}
