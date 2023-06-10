package meal

import (
	"github.com/FACorreiaa/Stay-Healthy-Backend/db"
	"github.com/gofiber/fiber/v2"
)

type QueryExecMode uint

const (
	CacheStatement = iota
)

func (m QueryExecMode) value() string {
	switch m {
	case CacheStatement:
		return "cache_statement"
	default:
		return ""
	}
}

type Config struct {
	postgresConfig db.Config
}

type Nutrients struct {
	ID                 int     `json:"id" pg:"default:gen_random_uuid()"`
	Name               string  `json:"name"`
	Calories           float64 `json:"calories"`
	ServingSize        float64 `json:"serving_size_g"`
	FatTotal           float64 `json:"fat_total_g"`
	FatSaturated       float64 `json:"fat_saturated_g"`
	Protein            float64 `json:"protein_g"`
	Sodium             int     `json:"sodium_mg"`
	Potassium          int     `json:"potassium_mg"`
	Cholesterol        int     `json:"cholesterol_mg"`
	CarbohydratesTotal float64 `json:"carbohydrates_total_g"`
	Fiber              float64 `json:"fiber_g"`
	Sugar              float64 `json:"sugar_g"`
}

//func NewConfig(config Config) Config {
//	return Config{}
//}

type NutrientService interface {
	GetNutrients(c *fiber.Ctx) ([]Nutrients, error)
}

type Service struct {
	Nutrient NutrientService
}

type NutrientRepository interface {
	GetNutrients(c *fiber.Ctx) ([]Nutrients, error)
}

type Repository struct {
	Nutrient NutrientRepository
}

//func NewService(repo *repository.Repository) *Service {
//	return &Service{
//		User:     user.NewService(repo),
//
//	}
//}
