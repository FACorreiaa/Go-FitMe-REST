package workouts

import (
	"encoding/json"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/auth"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strings"
	"time"
)

type Handler struct {
	service *StructWorkout
}

func NewExerciseHandler(s *StructWorkout) *Handler {
	return &Handler{
		service: s,
	}
}

// GetExercises godoc
// @Summary      GetExercises
// @Description  Get all exercises
// @Tags         workouts exercises
// @Accept       json
// @Produce      json
// @Success      200  {array}   Exercises
// @Router       /exercises [get]
func (h Handler) GetExercises(w http.ResponseWriter, r *http.Request) {
	exercises, err := h.service.Workout.GetAllExercises(r.Context())

	if err != nil {
		log.Printf("Error fetching exercises data: %v", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(exercises)
}

// GetExerciseByID godoc
// @Summary      GetExerciseByID
// @Description  Get exercise by its id
// @Tags         workouts exercises
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Exercise ID"
// @Success      200  {array}   Exercises
// @Router       /exercises/{id} [get]
func (h Handler) GetExerciseByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))

	exercises, err := h.service.Workout.GetExerciseByID(r.Context(), id)

	if err != nil {
		log.Printf("Error fetching exercises data: %v", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(exercises)
}

// InsertExercise godoc
// @Summary      Insert exercise
// @Description  Insert a new exercise on the list
// @Tags         workouts exercises
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Exercise ID"
// @Success      200  {array}   Exercises
// @Router       /exercises/{id} [post]
func (h Handler) InsertExercise(w http.ResponseWriter, r *http.Request) {
	var newExercise Exercises

	err := json.NewDecoder(r.Body).Decode(&newExercise)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	// Assuming you have a service method called InsertNewUserExercise that handles database insert
	response, err := h.service.Workout.InsertExercise(userSession.Id, Exercises{
		ID:            uuid.New(),
		Name:          newExercise.Name,
		ExerciseType:  newExercise.ExerciseType,
		MuscleGroup:   newExercise.MuscleGroup,
		Equipment:     newExercise.Equipment,
		Difficulty:    newExercise.Difficulty,
		Instructions:  newExercise.Instructions,
		Video:         newExercise.Video,
		CustomCreated: true,
		CreatedAt:     time.Now(),
		UpdatedAt:     nil,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// DeleteExercise godoc
// @Summary      Delete exercise
// @Description  Delete an exercise on the list
// @Tags         workouts exercises
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Exercise ID"
// @Success      200  {array}   Exercises
// @Router       /exercises/{id} [delete]
func (h Handler) DeleteExercise(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	err = h.service.Workout.DeleteExercise(userSession.Id, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// UpdateExercise godoc
// @Summary      Update exercise
// @Description  Update an exercise on the list
// @Tags         workouts exercises
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Exercise ID"
// @Success      200  {array}   Exercises
// @Router       /exercises/{id} [patch]
func (h Handler) UpdateExercise(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*user.UserSession)
	//if !ok {
	//	http.Error(w, "User session not found", http.StatusUnauthorized)
	//	return
	//}

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updates["UpdatedAt"] = time.Now()

	err = h.service.Workout.UpdateExercise(id, updates)
	if err != nil {
		if strings.Contains(err.Error(), "no rows were updated") {
			http.Error(w, "Exercise not found", http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// CreateWorkoutPlan godoc
// @Summary      Create workout plan
// @Description  Create a new workout plan
// @Tags         workouts
// @Accept       json
// @Produce      json
// @Router       /exercises/workout/plan [post]
func (h Handler) CreateWorkoutPlan(w http.ResponseWriter, r *http.Request) {
	var requestBody CreateWorkoutPlanRequest

	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	workoutPlan := requestBody.WorkoutPlan
	workoutPlan.UserID = userSession.Id

	// Create WorkoutDays based on the provided plan
	workoutDays := make([]WorkoutPlanDay, len(requestBody.Plan))
	for i, planDay := range requestBody.Plan {
		exercises := make([]Exercises, len(planDay.ExerciseIDs))
		for j, exerciseID := range planDay.ExerciseIDs {
			// Fetch exercises details
			exerciseDetails, err := h.service.Workout.GetExerciseByID(r.Context(), exerciseID)
			if err != nil {
				// Handle error (exercise not found, etc.)
				http.Error(w, err.Error(), http.StatusBadRequest)
				return
			}
			exercises[j] = exerciseDetails
		}
		workoutDays[i] = WorkoutPlanDay{
			Day:       planDay.Day,
			Exercises: exercises,
		}
	}

	response, err := h.service.Workout.CreateWorkoutPlan(WorkoutPlan{
		ID:          uuid.New(),
		UserID:      userSession.Id,
		Description: requestBody.WorkoutPlan.Description,
		Notes:       requestBody.WorkoutPlan.Notes,
		Rating:      requestBody.WorkoutPlan.Rating,
		WorkoutDays: workoutDays,
		CreatedAt:   time.Now(),
	}, requestBody.Plan)

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	result := struct {
		WorkoutPlan WorkoutPlan `json:"workoutPlan"`
	}{
		WorkoutPlan: response,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(result)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

// GetWorkoutPlans godoc
// @Summary      Get workout plan
// @Description  Retrieve all workout plans
// @Tags         workouts
// @Accept       json
// @Produce      json
// @Success      200  {array}   WorkoutPlan
// @Router       /exercises/workout/plan [get]
func (h Handler) GetWorkoutPlans(w http.ResponseWriter, r *http.Request) {
	workoutPlan, err := h.service.Workout.GetWorkoutPlans(r.Context())

	if err != nil {
		println(err)
		log.Printf("Error fetching workout plan data: %v", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(workoutPlan)
}

// GetWorkoutPlan godoc
// @Summary      Get workout plan
// @Description  Retrieve workout plan by id
// @Tags         workouts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Workout plan ID"
// @Success      200  {array}   WorkoutPlan
// @Router       /exercises/workout/plan/{id} [get]
func (h Handler) GetWorkoutPlan(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))

	workoutPlan, err := h.service.Workout.GetWorkoutPlan(r.Context(), id)

	if err != nil {
		println(err)
		log.Printf("Error fetching workout plan data: %v", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(workoutPlan)
}

// DeleteWorkoutPlan godoc
// @Summary      Delete workout plan
// @Description  Delete workout plan by id
// @Tags         workouts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Workout plan ID"
// @Success      200  {array}   WorkoutPlan
// @Router       /exercises/workout/plan/{id} [delete]
func (h Handler) DeleteWorkoutPlan(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	err = h.service.Workout.DeleteWorkoutPlan(userSession.Id, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// UpdateWorkoutPlan godoc
// @Summary      Update workout plan
// @Description  Update workout plan by id
// @Tags         workouts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Workout plan ID"
// @Success      200  {array}   WorkoutPlan
// @Router       /exercises/workout/plan/{id} [patch]
func (h Handler) UpdateWorkoutPlan(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*user.UserSession)
	//if !ok {
	//	http.Error(w, "User session not found", http.StatusUnauthorized)
	//	return
	//}

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updates["UpdatedAt"] = time.Now()

	err = h.service.Workout.UpdateWorkoutPlan(id, updates)
	if err != nil {
		if strings.Contains(err.Error(), "no rows were updated") {
			http.Error(w, "workout plan not found", http.StatusBadRequest)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// GetWorkoutPlanExercises godoc
// @Summary      Get exercises from workout plan
// @Description  Get exercises from workout plan
// @Tags         workouts
// @Accept       json
// @Produce      json
// @Success      200  {array}   WorkoutPlan
// @Router       /exercises/workout/plan/exercise [get]
func (h Handler) GetWorkoutPlanExercises(w http.ResponseWriter, r *http.Request) {
	workoutPlanExercises, err := h.service.Workout.GetWorkoutPlanExercises(r.Context())

	if err != nil {
		log.Printf("Error fetching workout plan data: %v", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(workoutPlanExercises)
}

// GetWorkoutPlanIdExercises godoc
// @Summary      Get exercises by id from workout plan
// @Description  Get exercises by id from workout plan
// @Tags         workouts
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Workout plan ID"
// @Success      200  {array}   WorkoutPlan
// @Router       /exercises/workout/plan/exercise/{id} [get]
func (h Handler) GetWorkoutPlanIdExercises(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))

	workoutPlanExercises, err := h.service.Workout.GetWorkoutPlanIdExercises(r.Context(), id)

	if err != nil {
		log.Printf("Error fetching workout plan data: %v", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(workoutPlanExercises)
}

// DeleteWorkoutPlanExercise godoc
// @Summary      Delete exercises by id from workout plan
// @Description  Delete exercises by id from workout plan
// @Tags         workouts
// @Accept       json
// @Produce      json
// @Param        workoutPlanID   path      int  true  "workout_plan_id"
// @Param        workoutDay   path      string  true  "Day"
// @Param        exerciseID   path      int  true  "exercise_id"
// @Router       /exercises/workout/plan/{workoutPlanID}/day/{workoutDay}/exercise/{exerciseID} [delete]
func (h Handler) DeleteWorkoutPlanExercise(w http.ResponseWriter, r *http.Request) {
	workoutPlanID, err := uuid.Parse(chi.URLParam(r, "workoutPlanID"))
	workoutDay := chi.URLParam(r, "workoutDay")
	exerciseID, err := uuid.Parse(chi.URLParam(r, "exerciseID"))

	err = h.service.Workout.DeleteWorkoutPlanIdExercises(workoutDay, workoutPlanID, exerciseID)

	if err != nil {
		log.Printf("Error fetching workout plan data: %v", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// CreateWorkoutPlanExercise godoc
// @Summary      Insert new exercise into workout plan
// @Description  Insert new exercise into workout plan
// @Tags         workouts
// @Accept       json
// @Produce      json
// @Param        workoutPlanID   path      int  true  "workout_plan_id"
// @Param        workoutDay   path      string  true  "Day"
// @Param        exerciseID   path      int  true  "exercise_id"
// @Success      200  {array}   WorkoutPlan
// @Router       /exercises/workout/plan/{workoutPlanID}/day/{workoutDay}/exercise/{exerciseID} [post]
func (h Handler) CreateWorkoutPlanExercise(w http.ResponseWriter, r *http.Request) {
	workoutPlanID, err := uuid.Parse(chi.URLParam(r, "workoutPlanID"))
	workoutDay := chi.URLParam(r, "workoutDay")
	exerciseID, err := uuid.Parse(chi.URLParam(r, "exerciseID"))

	err = h.service.Workout.CreateWorkoutPlanExercise(workoutDay, workoutPlanID, exerciseID)

	if err != nil {
		log.Printf("Error fetching workout plan data: %v", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// UpdateWorkoutPlanExercise godoc
// @Summary      Update exercise into workout plan
// @Description  Update exercise into workout plan
// @Tags         workouts
// @Accept       json
// @Produce      json
// @Param        workoutPlanID   path      int  true  "Workout plan ID"
// @Param        workoutDay   path      string  true  "Day"
// @Param        exerciseID   path      int  true  "Exercise ID"
// @Param        prevExerciseID   path      int  true  "Exercise ID"
// @Router       /exercises/workout/plan/{workoutPlanID}/day/{workoutDay}/exercise/{prevExerciseID}/{exerciseID} [patch]
func (h Handler) UpdateWorkoutPlanExercise(w http.ResponseWriter, r *http.Request) {
	workoutPlanID, err := uuid.Parse(chi.URLParam(r, "workoutPlanID"))
	workoutDay := chi.URLParam(r, "workoutDay")
	exerciseID, err := uuid.Parse(chi.URLParam(r, "exerciseID"))
	prevExerciseID, err := uuid.Parse(chi.URLParam(r, "prevExerciseID"))

	err = h.service.Workout.UpdateWorkoutPlanExercise(workoutDay, workoutPlanID, exerciseID, prevExerciseID)

	if err != nil {
		log.Printf("Error fetching workout plan data: %v", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}
