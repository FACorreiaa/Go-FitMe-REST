package meal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MealRepository struct {
	db *pgxpool.Pool
}

// ListAccounts lists all existing accounts
//
//  @Summary      List nutrients
//  @Description  get nutrients
//  @Tags         nutrients
//  @Accept       json
//  @Produce      json
//  @Param        q    query     string  false  "name search by q"  Format(email)
//  @Success      200  {array}   meal.Nutrient
//  @Failure      400  {object}  httputil.HTTPError
//  @Failure      404  {object}  httputil.HTTPError
//  @Failure      500  {object}  httputil.HTTPError
//  @Router       /api/v1/nutrients [get]

func (r *MealRepository) GetNutrients(c *fiber.Ctx) ([]Nutrients, error) {
	//var Nutrients []internal.Nutrients
	nutrients, err := r.GetNutrients(c)
	return nutrients, err
}

//func (u *userRepository) Delete(user *User) error {
//	return u.DB.Delete(&user).Error
//}
//func (u *userRepository) Migrate() error {
//	return u.DB.AutoMigrate(&domain.User{}).Error
//}
