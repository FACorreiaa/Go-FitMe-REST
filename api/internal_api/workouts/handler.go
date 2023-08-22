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

type DependenciesWorkouts interface {
	GetWorkoutsService() *ServiceWorkout
}

type Handler struct {
	dependencies DependenciesWorkouts
}

func NewExerciseHandler(deps DependenciesWorkouts) *Handler {
	return &Handler{
		dependencies: deps,
	}
}

func (h Handler) GetExercises(w http.ResponseWriter, r *http.Request) {
	exercises, err := h.dependencies.GetWorkoutsService().GetAllExercises(r.Context())

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

func (h Handler) GetExerciseByID(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))

	exercises, err := h.dependencies.GetWorkoutsService().GetExerciseByID(r.Context(), id)

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
	response, err := h.dependencies.GetWorkoutsService().InsertExercise(userSession.Id, Exercises{
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

	err = h.dependencies.GetWorkoutsService().DeleteExercise(userSession.Id, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

func (h Handler) UpdateExercise(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
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

	err = h.dependencies.GetWorkoutsService().UpdateExercise(id, updates)
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

	response, err := h.dependencies.GetWorkoutsService().CreateWorkoutPlan(WorkoutPlan{
		ID:          uuid.New(),
		UserID:      userSession.Id,
		Description: requestBody.WorkoutPlan.Description,
		Notes:       requestBody.WorkoutPlan.Notes,
		Rating:      requestBody.WorkoutPlan.Rating,
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
