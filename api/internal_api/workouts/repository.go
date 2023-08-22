package workouts

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"time"
)

type RepositoryWorkouts struct {
	db *sqlx.DB
}

func NewWorkoutsRepository(db *sqlx.DB) (*RepositoryWorkouts, error) {
	return &RepositoryWorkouts{db: db}, nil
}

func (r RepositoryWorkouts) GetAllExercises(ctx context.Context) ([]Exercises, error) {
	exercises := make([]Exercises, 0)
	query := `SELECT DISTINCT
    			id, name, type, muscle, equipment, difficulty,
				instructions, video, created_at, updated_at
				FROM exercise_list`

	err := r.db.SelectContext(ctx, &exercises, query)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return exercises, fmt.Errorf("exercises not found %w", err)
		}
		return exercises, fmt.Errorf("failed to scan exercises: %w", err)
	}

	return exercises, nil
}

func (r RepositoryWorkouts) GetExerciseByID(ctx context.Context, id uuid.UUID) (Exercises, error) {
	var exerciseList Exercises
	query := `SELECT 	id, name, type, muscle, equipment, difficulty,
						instructions, video, created_at, updated_at
			   FROM exercise_list
			   WHERE id = $1`
	err := r.db.GetContext(ctx, &exerciseList, query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return exerciseList, fmt.Errorf("exercise id %d not found: %w", id, err)
		}
		return exerciseList, fmt.Errorf("failed to scan activity: %w", err)
	}

	return exerciseList, nil
}

func (r RepositoryWorkouts) InsertNewExercise(userID int, exercise Exercises) (Exercises, error) {
	tx := r.db.MustBegin()
	defer tx.Rollback()

	// Insert into exercise_list
	_, err := tx.NamedExec(`
        INSERT INTO exercise_list (id, name, type, muscle, equipment, difficulty,
                                   instructions, video,
                                   created_at, updated_at)
        VALUES (:id, :name, :type, :muscle, :equipment, :difficulty,
                :instructions, :video, :created_at, :updated_at)`,
		exercise)
	if err != nil {
		return Exercises{}, fmt.Errorf("failed to insert exercise: %w", err)
	}

	// Insert into user_exercises
	_, err = tx.Exec(`
        INSERT INTO user_exercises (user_id, exercise_id)
        VALUES ($1, $2)`,
		userID, exercise.ID)
	if err != nil {
		return Exercises{}, fmt.Errorf("failed to insert exercise: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return Exercises{}, fmt.Errorf("failed to commit exercise: %w", err)
	}

	return exercise, nil
}

func (r RepositoryWorkouts) DeleteUserExercise(userID int, exerciseID uuid.UUID) error {
	tx := r.db.MustBegin()
	defer tx.Rollback()

	// Delete from user_exercises
	result, err := tx.Exec(`
        DELETE FROM user_exercises
        WHERE user_id = $1 AND exercise_id = $2`,
		userID, exerciseID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("exercise id %d not found: %w", exerciseID, err)
		}
		return fmt.Errorf("failed to delete exercise: %w", err)
	}

	// Delete from exercise_list
	_, err = tx.Exec(`
        DELETE FROM exercise_list
        WHERE id = $1 AND custom_created = $2`,
		exerciseID, true)
	if err != nil {
		return fmt.Errorf("failed to delete exercise: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were deleted")
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("error commiting the transaction: %w", err)
	}

	return nil
}

func (r RepositoryWorkouts) UpdateExercise(id uuid.UUID, updates map[string]interface{}) error {
	query := `
		UPDATE exercise_list
		SET name = :name, type = :type, muscle = :muscle,
		    equipment = :equipment, difficulty = :difficulty,
		    instructions = :instructions, video = :video,
		    updated_at = :updated_at
		WHERE id = :id AND custom_created = true
	`

	namedParams := map[string]interface{}{
		"id":           id,
		"name":         updates["name"],
		"type":         updates["type"],
		"muscle":       updates["muscle"],
		"equipment":    updates["equipment"],
		"difficulty":   updates["difficulty"],
		"instructions": updates["instructions"],
		"video":        updates["video"],
		"updated_at":   updates["UpdatedAt"],
	}

	result, err := r.db.NamedExec(query, namedParams)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("exercises not found %w", err)
		}
		return fmt.Errorf("failed to scan exercises: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated")
	}

	return err
}

//// Repository layer
//func (r RepositoryWorkouts) CreateWorkoutPlan(newPlan WorkoutPlan, workoutDays []WorkoutDay) (WorkoutPlan, error) {
//	tx, err := r.db.Beginx()
//	if err != nil {
//		return WorkoutPlan{}, fmt.Errorf("failed to start transaction: %w", err)
//	}
//	defer tx.Rollback()
//
//	// Insert workout plan
//	workoutPlan, err := r.repo.CreateWorkoutPlan(tx, newPlan)
//	if err != nil {
//		return WorkoutPlan{}, fmt.Errorf("failed to insert workout plan: %w", err)
//	}
//
//	// Insert workout days
//	err = r.repo.InsertWorkoutDays(tx, workoutPlan.ID, workoutDays)
//	if err != nil {
//		return WorkoutPlan{}, fmt.Errorf("failed to insert workout days: %w", err)
//	}
//
//	err = tx.Commit()
//	if err != nil {
//		return WorkoutPlan{}, fmt.Errorf("failed to commit transaction: %w", err)
//	}
//
//	return workoutPlan, nil
//}

//func (r RepositoryWorkouts) CreateWorkoutPlan(tx *sqlx.Tx, newPlan WorkoutPlan) (WorkoutPlan, error) {
//	query := `INSERT INTO workout_plan (id, user_id, description, notes, created_at, updated_at, rating)
//				VALUES (:id, :user_id, :description, :notes, :created_at, :updated_at, :rating)
//				RETURNING *`
//
//	workoutPlan := WorkoutPlan{
//		ID:          newPlan.ID,
//		UserID:      newPlan.UserID,
//		Description: newPlan.Description,
//		Notes:       newPlan.Notes,
//		CreatedAt:   newPlan.CreatedAt,
//		UpdatedAt:   newPlan.UpdatedAt,
//		Rating:      newPlan.Rating,
//	}
//
//	err := tx.Get(&workoutPlan, query, workoutPlan)
//	if err != nil {
//		return WorkoutPlan{}, fmt.Errorf("failed to insert workout plan: %w", err)
//	}
//
//	return workoutPlan, nil
//}

//func (r RepositoryWorkouts) InsertWorkoutDays(tx *sqlx.Tx, workoutPlanID uuid.UUID, workoutDays []WorkoutDay) error {
//	query := `INSERT INTO workout_day (workout_plan_id, exercise_id, created_at, updated_at)
//				VALUES (:workout_plan_id, :exercise_id, :created_at, :updated_at)`
//
//	for _, wd := range workoutDays {
//		wd.WorkoutPlanID = workoutPlanID
//		_, err := tx.NamedExec(query, wd)
//		if err != nil {
//			return fmt.Errorf("failed to insert workout day: %w", err)
//		}
//	}
//
//	return nil
//}

// Repository layer

func (r RepositoryWorkouts) CreateWorkoutPlan(newPlan WorkoutPlan, plan []PlanDay) (WorkoutPlan, error) {
	tx, err := r.db.Beginx()
	if err != nil {
		return WorkoutPlan{}, fmt.Errorf("failed to start transaction: %w", err)
	}
	defer tx.Rollback()

	// Insert the workout plan
	query := `
        INSERT INTO workout_plan (id, user_id, description, notes, rating, created_at)
        VALUES (:id, :user_id, :description, :notes, :rating, :created_at)
        RETURNING *;
    `
	_, err = tx.NamedExecContext(context.Background(), query, newPlan)
	var insertedPlan WorkoutPlan

	if err != nil {
		return WorkoutPlan{}, fmt.Errorf("failed to insert workout plan: %w", err)
	}

	// Insert the workout days and associated exercises
	for _, day := range plan {
		workoutDayID := uuid.New()

		workoutDay := WorkoutDay{
			ID:            workoutDayID,
			WorkoutPlanID: newPlan.ID,
			Day:           day.Day,
			CreatedAt:     time.Now(),
		}
		query := `
            INSERT INTO workout_day (id, workout_plan_id, day, created_at)
            VALUES (:id, :workout_plan_id, :day, :created_at);
        `
		result, err := tx.NamedExecContext(context.Background(), query, workoutDay)
		println(result)
		if err != nil {
			return WorkoutPlan{}, fmt.Errorf("failed to insert workout day: %w", err)
		}

		for _, exerciseID := range day.ExerciseIDs {
			workoutDayExercise := WorkoutDayExercise{
				WorkoutDayID: workoutDay.ID,
				ExerciseID:   exerciseID,
			}
			query := `
							INSERT INTO workout_day_exercise (id, workout_day_id, exercise_id)
							VALUES (:id, :workout_day_id, :exercise_id);
						`
			_, err := tx.NamedExecContext(context.Background(), query, workoutDayExercise)
			if err != nil {
				return WorkoutPlan{}, fmt.Errorf("failed to insert workout day exercise: %w", err)
			}

		}
	}

	// Insert the workout plan details
	for _, day := range plan {
		workoutPlanDetail := WorkoutPlanDetail{
			ID:            uuid.New(),
			WorkoutPlanID: insertedPlan.ID,
			Day:           day.Day,
			Exercises:     day.ExerciseIDs,
			CreatedAt:     time.Now(),
		}
		insertWorkoutPlanDetailQuery := `
            INSERT INTO workout_plan_detail (id, workout_plan_id, day, exercises, created_at)
            VALUES (:id, :workout_plan_id, :day, :exercises, :created_at);
        `
		_, err := tx.NamedExecContext(context.Background(), insertWorkoutPlanDetailQuery, workoutPlanDetail)
		if err != nil {
			return WorkoutPlan{}, fmt.Errorf("failed to insert workout plan detail: %w", err)
		}
	}

	err = tx.Commit()
	if err != nil {
		return WorkoutPlan{}, fmt.Errorf("failed to commit transaction: %w", err)
	}

	return newPlan, nil
}
