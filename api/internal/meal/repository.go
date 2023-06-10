package meal

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5/pgxpool"
)

type MealRepository struct {
	db *pgxpool.Pool
}

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
