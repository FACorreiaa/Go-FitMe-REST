package workouts

import "github.com/jmoiron/sqlx"

type RepositoryWorkouts struct {
	db *sqlx.DB
}

func NewWorkoutsRepository(db *sqlx.DB) (*RepositoryWorkouts, error) {
	return &RepositoryWorkouts{db: db}, nil
}
