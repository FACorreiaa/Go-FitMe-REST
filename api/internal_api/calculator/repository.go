package calculator

import "github.com/jmoiron/sqlx"

type Repository struct {
	db *sqlx.DB
}

func NewCalculatorRepository(db *sqlx.DB) (*Repository, error) {
	return &Repository{db: db}, nil
}

func (r *Repository) InsertDietGoals(data UserMacroDistribution) (UserMacroDistribution, error) {
	query := `INSERT INTO user_macro_distribution (user_id, age, height, weight,
                                     gender, system, activity, activity_description, objective,
									objective_description, calories_distribution, calories_distribution_description,
                                     protein, fats, carbs, bmr, tdee, goal, created_at)
				VALUES (:user_id, :age, :height, :weight, :gender, :system, :activity,
				        :activity_description, :objective, :objective_description, :calories_distribution,
				        :calories_distribution_description, :protein, :fats, :carbs,
				        :bmr, :tdee, :goal, :created_at)
				RETURNING *`
	namedStmt, err := r.db.PrepareNamed(query)
	if err != nil {
		return UserMacroDistribution{}, err
	}

	var insertedData UserMacroDistribution
	err = namedStmt.Get(&insertedData, data)
	if err != nil {
		return UserMacroDistribution{}, err
	}

	return insertedData, nil
}
