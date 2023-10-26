package workouts

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"strings"
	"time"
)

type Repository struct {
	db *sqlx.DB
}

func NewWorkoutRepository(db *sqlx.DB) (*Repository, error) {
	return &Repository{db: db}, nil
}

func (r Repository) GetAllExercises(ctx context.Context) ([]Exercises, error) {
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

func (r Repository) GetExerciseByID(ctx context.Context, id string) (Exercises, error) {
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

func (r Repository) InsertNewExercise(userID int, exercise Exercises) (Exercises, error) {
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

func (r Repository) DeleteUserExercise(userID int, exerciseID string) error {
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

func (r Repository) UpdateExercise(id string, updates map[string]interface{}) error {
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

func (r Repository) CreateWorkoutPlan(newPlan WorkoutPlan, plan []PlanDay) (WorkoutPlan, error) {
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
	if err != nil {
		return WorkoutPlan{}, fmt.Errorf("failed to insert workout plan: %w", err)
	}

	var insertedPlan WorkoutPlan

	// Fetch the returned plan and assign it to insertedPlan
	err = tx.GetContext(context.Background(), &insertedPlan, "SELECT * FROM workout_plan WHERE id = $1", newPlan.ID)
	if err != nil {
		return WorkoutPlan{}, fmt.Errorf("failed to fetch inserted workout plan: %w", err)
	}

	// Insert the workout days and associated exercises
	for _, day := range plan {
		workoutDayID := uuid.NewString()

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
	}

	// Insert the workout plan details
	for _, day := range plan {
		workoutPlanDetail := WorkoutPlanDetail{
			ID:            uuid.NewString(),
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

func (r Repository) GetWorkoutPlans(ctx context.Context) ([]WorkoutPlanResponse, error) {
	query := `
      SELECT 	wp.id AS workout_plan_id, wp.user_id, wp.description,
			  	wp.notes, wp.rating, wp.created_at, wd.day, wpd.exercises
      			FROM workout_plan AS wp
				LEFT JOIN workout_plan_detail AS wpd ON wp.id = wpd.workout_plan_id
				LEFT JOIN workout_day AS wd ON wp.id = wd.workout_plan_id
				GROUP BY wp.id, wd.day, wpd.exercises
			ORDER BY wd.day;
  `

	rows, err := r.db.QueryxContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	workouts := make(map[string]WorkoutPlanResponse)

	for rows.Next() {
		var row struct {
			WorkoutPlanID string           `json:"workout_plan_id,string" db:"workout_plan_id"`
			UserID        int              `json:"user_id" db:"user_id"`
			Day           string           `db:"day"`
			Description   string           `json:"description" db:"description"`
			Exercises     pq.StringArray   `json:"exercises" db:"exercises"`
			Notes         string           `json:"notes" db:"notes"`
			CreatedAt     time.Time        `json:"created_at" db:"created_at"`
			UpdatedAt     *time.Time       `json:"updated_at" db:"updated_at"`
			Rating        int              `json:"rating" db:"rating"`
			WorkoutDays   []WorkoutPlanDay `json:"workoutDays" db:"-"`
		}
		err := rows.StructScan(&row)
		if err != nil {
			return nil, err
		}

		if _, ok := workouts[row.WorkoutPlanID]; !ok {
			workouts[row.WorkoutPlanID] = WorkoutPlanResponse{
				WorkoutPlanID: row.WorkoutPlanID,
				UserID:        row.UserID,
				Description:   row.Description,
				WorkoutDays:   []WorkoutDayResponse{},
				Notes:         row.Notes,
				CreatedAt:     row.CreatedAt,
				UpdatedAt:     row.UpdatedAt,
				Rating:        row.Rating,
			}
		}

		if plan, ok := workouts[row.WorkoutPlanID]; ok {
			// Create a new WorkoutDayResponse and append it to the WorkoutDays
			day := WorkoutDayResponse{
				Day:       row.Day,
				Exercises: row.Exercises,
			}
			plan.WorkoutDays = append(plan.WorkoutDays, day)

			// Store the updated plan back in the map
			workouts[row.WorkoutPlanID] = plan
		}
	}

	result := make([]WorkoutPlanResponse, 0, len(workouts))
	for _, workout := range workouts {
		result = append(result, workout)
	}

	return result, nil
}

func (r Repository) GetWorkoutPlan(ctx context.Context, id string) (WorkoutPlanResponse, error) {
	var workoutPlan WorkoutPlanResponse
	query := `
      SELECT 	wp.id AS workout_plan_id, wp.user_id, wp.description,
			  	wp.notes, wp.rating, wp.created_at, wd.day, wpd.exercises
      			FROM workout_plan AS wp
				JOIN workout_plan_detail AS wpd ON wp.id = wpd.workout_plan_id
				JOIN workout_day AS wd ON wp.id = wd.workout_plan_id
      			WHERE wp.id = $1
				GROUP BY wp.id, wd.day, wpd.exercises
			ORDER BY wd.day;
  `

	err := r.db.GetContext(ctx, &workoutPlan, query, id)
	if err != nil {
		return workoutPlan, err
	}

	return workoutPlan, nil
}

func (r Repository) DeleteWorkoutPlan(userID int, workoutPlanID string) error {
	tx := r.db.MustBegin()
	defer tx.Rollback()

	// Delete from workout_plan
	result, err := tx.Exec(`
        DELETE FROM workout_day
	   	WHERE workout_plan_id = $1`,
		workoutPlanID)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("workout plan id %d not found: %w", workoutPlanID, err)
		}
		return fmt.Errorf("failed to delete workout plan: %w", err)
	}

	_, err = tx.Exec(`
		DELETE FROM workout_plan_detail
	   	WHERE workout_plan_id = $1`,
		workoutPlanID)
	if err != nil {
		return fmt.Errorf("failed to delete workout_plan_detail: %w", err)
	}

	// Delete from workout_plan_detail
	_, err = tx.Exec(`
		DELETE FROM workout_plan
		WHERE id = $1 AND user_id = $2`,
		workoutPlanID, userID)
	if err != nil {
		return fmt.Errorf("failed to delete workout_plan: %w", err)
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

func (r Repository) UpdateWorkoutPlan(id string, updates map[string]interface{}) error {
	query := `
		UPDATE workout_plan
		SET description = :description, notes = :notes, rating = :rating,
		    updated_at = :updated_at
		WHERE id = :id
	`

	namedParams := map[string]interface{}{
		"id":          id,
		"description": updates["description"],
		"notes":       updates["notes"],
		"rating":      updates["rating"],
		"updated_at":  updates["UpdatedAt"],
	}

	result, err := r.db.NamedExec(query, namedParams)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("workout plan not found %w", err)
		}
		return fmt.Errorf("failed to scan workout plan: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated")
	}

	return err
}

func (r Repository) GetExerciseByIdWorkoutPlan(ctx context.Context, id string) (WorkoutExerciseDay, error) {
	var workoutExerciseDayList WorkoutExerciseDay
	query := `
				SELECT el.*, wpd.day
					FROM workout_plan_detail wpd
					JOIN exercise_list el ON el.id = ANY(wpd.exercises)
					WHERE wpd.workout_plan_id = $1;
				`
	err := r.db.GetContext(ctx, &workoutExerciseDayList, query, id)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return workoutExerciseDayList, fmt.Errorf("workout plan id %d not found: %w", id, err)
		}
		return workoutExerciseDayList, fmt.Errorf("failed to scan activity: %w", err)
	}

	return workoutExerciseDayList, nil
}

func (r Repository) GetWorkoutPlanExercises(ctx context.Context) ([]WorkoutExerciseDay, error) {
	workoutExerciseDayList := make([]WorkoutExerciseDay, 0)
	query := `
				SELECT el.*, wpd.day
					FROM workout_plan_detail wpd
					JOIN exercise_list el ON el.id = ANY(wpd.exercises)
				`
	err := r.db.SelectContext(ctx, &workoutExerciseDayList, query)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return workoutExerciseDayList, fmt.Errorf("workout plan not found: %w", err)
		}
		return workoutExerciseDayList, fmt.Errorf("failed to scan activity: %w", err)
	}

	return workoutExerciseDayList, nil
}

func (r Repository) DeleteWorkoutPlanIdExercises(workoutDay string, workoutPlanID string, exerciseID string) error {
	tx := r.db.MustBegin()
	defer tx.Rollback()

	// Delete from user_exercises
	result, err := tx.Exec(`
        UPDATE workout_plan_detail
		SET exercises = array_remove(exercises, $1)
		WHERE workout_plan_id = $2 AND day = $3`,
		exerciseID, workoutPlanID, workoutDay)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("exercise id %d not found: %w", exerciseID, err)
		}
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

func (r Repository) CreateExerciseWorkoutPlan(workoutDay string, workoutPlanID string, exerciseID string) error {
	tx := r.db.MustBegin()
	defer tx.Rollback()

	// Prepare an SQL statement to insert the exercise into the workout_plan_detail
	query := `
		UPDATE workout_plan_detail
		SET exercises = array_append(exercises, $1)
		WHERE workout_plan_id = $2 AND day = $3
	`

	_, err := tx.Exec(query, exerciseID, workoutPlanID, workoutDay)
	if err != nil {
		return fmt.Errorf("failed to insert exercise: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit exercise: %w", err)
	}

	return nil
}

func (r Repository) UpdateExerciseByIdWorkoutPlan(workoutDay string, workoutPlanID string, exerciseID string, prevExerciseID string) error {
	tx := r.db.MustBegin()
	defer tx.Rollback()

	// Prepare an SQL statement to insert the exercise into the workout_plan_detail
	query := `
		UPDATE workout_plan_detail
		SET exercises = array_replace($1, $2)
		WHERE workout_plan_id = $3 AND day = $4
	`

	_, err := tx.Exec(query, prevExerciseID, exerciseID, workoutPlanID, workoutDay)
	if err != nil {
		return fmt.Errorf("failed to insert exercise: %w", err)
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("failed to commit exercise: %w", err)
	}

	return nil
}

//func (r Repository) GetCompleteWorkoutData(ctx context.Context, userID int, workoutPlanID string) ([]WorkoutPlanExportData, error) {
//	workoutPlan := make([]WorkoutPlanExportData, 0)
//	//exercises := make([]Exercises, 0)
//
//	query := `
//		SELECT
//			wd.day AS workout_day,
//			wp.description AS workout_description,
//			el.id AS exercise_id,
//			el.name AS exercise_name,
//			el.type AS exercise_type,
//			el.muscle AS exercise_muscle,
//			el.equipment AS exercise_equipment,
//			el.difficulty AS exercise_difficulty,
//			el.instructions AS exercise_instructions,
//			el.video AS exercise_video
//		FROM workout_plan_detail wpd
//		JOIN workout_plan wp ON wpd.workout_plan_id = wp.id
//		JOIN exercise_list el ON el.id = ANY(wpd.exercises)
//		JOIN workout_day wd ON wd.workout_plan_id = wp.id
//		WHERE wp.id = $1 AND wp.user_id = $2;
//	`
//	err := r.db.SelectContext(ctx, &workoutPlan, query, workoutPlanID, userID)
//
//	if err != nil {
//		println(err)
//		if errors.Is(err, sql.ErrNoRows) {
//			return workoutPlan, fmt.Errorf("workoutplan id %d not found: %w", workoutPlanID, err)
//		}
//		return workoutPlan, fmt.Errorf("failed to scan exercises: %w", err)
//	}
//
//	return workoutPlan, nil
//}

func (r Repository) GetFullWorkoutPlan(ctx context.Context, workoutPlanID string, userID int) (WorkoutPlan, error) {
	var workoutPlan WorkoutPlan

	// Query the workout plan details from the database.
	query := `
        SELECT wp.id, wp.user_id, wpd.exercises,
               wp.description, wp.notes, wp.rating,
               wd.day, wp.created_at
        FROM workout_plan AS wp
        JOIN workout_day AS wd ON wp.id = wd.workout_plan_id
        JOIN workout_plan_detail AS wpd ON wp.id = wpd.workout_plan_id
        WHERE wp.id = $1 AND wp.user_id = $2
    `

	rows, err := r.db.QueryxContext(ctx, query, workoutPlanID, userID)
	if err != nil {
		return workoutPlan, fmt.Errorf("failed to execute query: %w", err)
	}
	defer rows.Close()

	// Create a map to store workout days and their associated exercises.
	workoutDays := make(map[string][]Exercises)

	for rows.Next() {
		var wpd struct {
			ID          string    `json:"id,string" db:"id" pg:"default:gen_random_uuid()"`
			UserID      int       `db:"user_id"`
			ExerciseIDs string    `db:"exercises" json:"exercises,string"`
			Description string    `db:"description"`
			Notes       string    `db:"notes"`
			Rating      int       `db:"rating"`
			Day         string    `db:"day"`
			CreatedAt   time.Time `db:"created_at"`
		}

		err := rows.StructScan(&wpd)
		if err != nil {
			return workoutPlan, fmt.Errorf("failed to scan rows: %w", err)
		}

		// Parse the ExerciseIDs string into a slice of string
		exerciseIDsStr := strings.Trim(wpd.ExerciseIDs, "{}") // Remove curly braces
		idStrings := strings.Split(exerciseIDsStr, ",")
		exerciseIDs := make([]string, len(idStrings))
		for i, id := range idStrings {
			//id, err := uuid.Parse(idStr)
			//if err != nil {
			//	return workoutPlan, fmt.Errorf("failed to parse exercise IDs: %w", err)
			//}
			exerciseIDs[i] = id
		}

		// Fetch exercise details for each UUID in the exerciseIDs slice.
		exerciseDetails := []Exercises{}
		for _, exerciseID := range exerciseIDs {
			exerciseDetail, err := r.GetExerciseByID(ctx, exerciseID)
			if err != nil {
				return workoutPlan, fmt.Errorf("failed to map exercises: %w", err)
			}
			exerciseDetails = append(exerciseDetails, exerciseDetail)
		}

		// Append exercise details to the appropriate day.
		workoutDays[wpd.Day] = append(
			workoutDays[wpd.Day],
			exerciseDetails...,
		)
		workoutPlan.Description = wpd.Description
		workoutPlan.Notes = wpd.Notes
		workoutPlan.Rating = wpd.Rating
	}

	// Populate the WorkoutPlan structure.
	workoutPlan.ID = workoutPlanID
	workoutPlan.UserID = userID

	workoutPlan.WorkoutDays = make([]WorkoutPlanDay, 0, len(workoutDays))
	for day, exercises := range workoutDays {
		workoutPlan.WorkoutDays = append(workoutPlan.WorkoutDays, WorkoutPlanDay{
			Day:       day,
			Exercises: exercises,
		})
	}

	return workoutPlan, nil
}

func (r Repository) ExportWorkoutToPDF() {}
