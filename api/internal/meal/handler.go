package meal

import (
	"github.com/getsentry/sentry-go"
	"github.com/gofiber/fiber/v2"
	"net/http"
)

//func NewCustomTime(t time.Time) structs.CustomTime {
//	return structs.CustomTime{Time: t}
//}

type Handler struct {
	service *Service
	c       *fiber.Ctx
}

func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}
func (u *Handler) GetNutrients(c *fiber.Ctx) {
	meals, err := u.service.Nutrient.GetNutrients(c)
	if err != nil {
		sentry.CaptureException(err)
		c.JSON(http.StatusBadRequest)
		return
	}
	c.JSON(meals)
}

//func (u *mealHandler) AddUser(c *gin.Context) {
//	var user domain.User
//	if err := c.ShouldBindJSON(&user); err != nil {
//		sentry.CaptureException(err)
//		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return
//	}
//	if err1 := (u.userService.Validate(&user)); err1 != nil {
//		sentry.CaptureException(err1)
//		c.JSON(http.StatusBadRequest, gin.H{"error": err1.Error()})
//		return
//	}
//	if ageValidation := (u.userService.ValidateAge(&user)); ageValidation != true {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid DOB"})
//		return
//	}
//	uid, err := u.fbService.CreateUser(user.Email, user.Password)
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Couldnot create user in firebase"})
//		return
//	}
//	user.ID = uid
//	u.userService.Create(&user)
//	c.JSON(http.StatusOK, user)
//}
