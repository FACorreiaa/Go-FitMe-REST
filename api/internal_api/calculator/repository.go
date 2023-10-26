package calculator

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
)

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

func (r *Repository) GetUserDietGoals(ctx context.Context, userID int) ([]UserMacroDistribution, error) {
	macroDistribution := make([]UserMacroDistribution, 0)
	query := `SELECT user_id, age, height, weight,
                      gender, system, activity, activity_description, objective,
					  objective_description, calories_distribution, calories_distribution_description,
                      protein, fats, carbs, bmr, tdee, goal, created_at
				FROM user_macro_distribution
				WHERE id = $1
				ORDER BY created_at`

	err := r.db.SelectContext(ctx, &macroDistribution, query, userID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return macroDistribution, fmt.Errorf("user macros not found %w", err)
		}
		return macroDistribution, fmt.Errorf("failed to scan exercises: %w", err)
	}

	return macroDistribution, nil
}

func (r *Repository) GetUserDietGoal(ctx context.Context, planID string) (UserMacroDistribution, error) {
	var macroDistribution UserMacroDistribution
	query := `SELECT id, user_id, age, height, weight,
                      gender, system, activity, activity_description, objective,
					  objective_description, calories_distribution, calories_distribution_description,
                      protein, fats, carbs, bmr, tdee, goal, created_at
				FROM user_macro_distribution
				WHERE id = $1
				ORDER BY created_at`

	err := r.db.GetContext(ctx, &macroDistribution, query, planID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return macroDistribution, fmt.Errorf("user macros not found %w", err)
		}
		return macroDistribution, fmt.Errorf("failed to scan exercises: %w", err)
	}

	return macroDistribution, nil
}
