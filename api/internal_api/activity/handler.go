package activity

import (
	"encoding/json"
	"fmt"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/auth"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	service          *StructActivity
	exerciseSessions map[string]*ExerciseSession // Map to store exercise sessions for each user
	pausedTimers     map[string]time.Time
}

func NewActivityHandler(s *StructActivity) *Handler {
	return &Handler{
		service: s,
	}
}

// GetActivities godoc
// @Summary      Show all activities
// @Description  get activities
// @Tags         activities
// @Accept       json
// @Produce      json
// @Success      200  {array}   Activity
// @Router       /activities [get]
func (h Handler) GetActivities(w http.ResponseWriter, r *http.Request) {
	activities, err := h.service.Activity.GetAll(r.Context())

	if err != nil {
		log.Printf("Error fetching activities data: %v", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(activities)
}

// GetActivitiesByName godoc
// @Summary      Show all activities by name
// @Description  Get activities by name
// @Tags         activities
// @Accept       json
// @Produce      json
// @Param        name   path      string  true  "Activity Name"
// @Success      200  {array}   Activity
// @Router       /activities/name/{name} [get]
func (h Handler) GetActivitiesByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	activities, err := h.service.Activity.GetByName(r.Context(), name)
	if err != nil {
		log.Printf("Error fetching activities data: %v", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(activities)
	if err != nil {
		_ = fmt.Errorf("failed to encode activities: %w", err)
		return
	}
}

// GetActivitiesById godoc
// @Summary      Show all activity by id
// @Description  Get activity by id
// @Tags         activities
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Activity ID"
// @Success      200  {array}   Activity
// @Router       /activities/id/{id} [get]
func (h Handler) GetActivitiesById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error parsing id: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}
	activities, err := h.service.Activity.GetByID(r.Context(), id)
	if err != nil {
		log.Printf("Error fetching activities data: %v", err)

		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(activities)
	if err != nil {
		_ = fmt.Errorf("failed to encode activities: %w", err)
		return
	}
}

// StartActivityTracker godoc
// @Summary      Start activity timer
// @Description  Start Activity
// @Tags         activities
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Activity ID"
// @Success      200  {array}   ExerciseSession
// @Router       /activities/start/session/id/{id} [post]
func (h Handler) StartActivityTracker(w http.ResponseWriter, r *http.Request) {
	activityID, err := strconv.Atoi(chi.URLParam(r, "id"))
	currentTime := time.Now()
	if err != nil {
		log.Printf("Error parsing id: %v", err)
		http.Error(w, "Internal server error", http.StatusNotFound)
		return
	}
	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	sessionId := strconv.Itoa(userSession.Id)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	_, found := h.exerciseSessions[sessionId]
	if found {
		http.Error(w, "Exercise session already in progress", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Printf("Error getting session: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	activity, err := h.service.Activity.GetByID(r.Context(), activityID)
	if err != nil {
		log.Printf("Error getting id activity: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	exerciseSession := &ExerciseSession{
		ID:          uuid.New(),
		UserID:      userSession.Id,
		ActivityID:  activity.ID,
		SessionName: activity.Name,
		StartTime:   currentTime,
		CreatedAt:   currentTime,
	}

	h.exerciseSessions[sessionId] = exerciseSession

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(exerciseSession)
}

// PauseActivityTracker godoc
// @Summary      Pause activity timer
// @Description  Pause Activity
// @Tags         activities
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Activity ID"
// @Success      200  {array}   ExerciseSession
// @Router       /activities/start/session/id/{id} [post]
func (h Handler) PauseActivityTracker(w http.ResponseWriter, r *http.Request) {
	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	sessionId := strconv.Itoa(userSession.Id)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	h.pausedTimers[sessionId] = time.Now()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(userSession)

}

// ResumeActivityTracker godoc
// @Summary      Resume activity timer
// @Description  Resume Activity
// @Tags         activities
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Activity ID"
// @Success      200  {array}   ExerciseSession
// @Router       /activities/start/session/id/{id} [post]
func (h Handler) ResumeActivityTracker(w http.ResponseWriter, r *http.Request) {
	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	sessionId := strconv.Itoa(userSession.Id)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	session, found := h.exerciseSessions[sessionId]
	if !found {
		http.Error(w, "Exercise session not found", http.StatusNotFound)
		return
	}

	// Calculate the duration of the paused state and update the start time
	pausedDuration := time.Since(h.pausedTimers[sessionId])
	session.StartTime = session.StartTime.Add(pausedDuration)

	// Clear the paused timer for the user
	delete(h.pausedTimers, sessionId)

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(session)

}

// StopActivityTracker godoc
// @Summary      Stop activity timer
// @Description  Stop Activity
// @Tags         activities
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Activity ID"
// @Success      200  {array}   ExerciseSession
// @Router       /activities/start/session/id/{id} [post]
func (h Handler) StopActivityTracker(w http.ResponseWriter, r *http.Request) {
	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	sessionId := strconv.Itoa(userSession.Id)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}
	// Get the exercise session for the user
	session, found := h.exerciseSessions[sessionId]
	if !found {
		http.Error(w, "Exercise session not found", http.StatusNotFound)
		return
	}

	// Calculate calories burned (assuming activity is fetched from the database)
	activity, err := h.service.Activity.GetByID(r.Context(), session.ActivityID)
	if err != nil {
		http.Error(w, "Error getting activity", http.StatusInternalServerError)
		return
	}

	//totalDurationSeconds := session.DurationHours*3600 + session.DurationMinutes*60 + session.DurationSeconds
	startUpTime := h.exerciseSessions[sessionId].StartTime
	totalDurationSeconds := int(time.Since(startUpTime).Seconds())

	session.DurationHours = totalDurationSeconds / 3600
	session.DurationMinutes = (totalDurationSeconds % 3600) / 60
	session.DurationSeconds = totalDurationSeconds % 60

	// Calculate calories burned per second
	caloriesPerSecond := activity.CaloriesPerHour / 3600

	// Calculate calories burned for the total duration in seconds
	session.CaloriesBurned = int(caloriesPerSecond * float32(totalDurationSeconds))

	session.EndTime = time.Now()

	err = h.service.Activity.SaveExerciseSession(r.Context(), session)
	if err != nil {
		http.Error(w, "Error saving exercise session to DB", http.StatusInternalServerError)
		return
	}

	delete(h.exerciseSessions, sessionId)

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(session)
}

// GetUserExerciseSession godoc
// @Summary      Get user exercise session
// @Description  Get exercise session
// @Tags         activities
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Activity ID"
// @Success      200  {array}   ExerciseSession
// @Router       /activities/user/exercises/user/{user_id} [post]
func (h Handler) GetUserExerciseSession(w http.ResponseWriter, r *http.Request) {

	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	exerciseSession, err := h.service.Activity.GetExerciseSession(r.Context(), userSession.Id)
	if err != nil {
		http.Error(w, "Error finding exercise session", http.StatusInternalServerError)
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(exerciseSession)
}

// GetUserExerciseTotalData godoc
// @Summary      Get user exercise data
// @Description  Get user exercise total data for durations and calories
// @Tags         activities
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Activity ID"
// @Success      200  {array}   ExerciseSession
// @Router       /activities/user/session/total/user/{user_id} [post]
func (h Handler) GetUserExerciseTotalData(w http.ResponseWriter, r *http.Request) {
	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	// Calculate and save the total exercise session data
	totalExerciseSession, err := h.service.Activity.GetExerciseTotalSession(r.Context(), userSession.Id)
	if err != nil {
		http.Error(w, "Error calculating total exercise session", http.StatusInternalServerError)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(totalExerciseSession)
}

func (h Handler) GetExerciseSessionStats(w http.ResponseWriter, r *http.Request) {
	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	// Calculate and save the total exercise session data
	sessionStats, err := h.service.Activity.GetExerciseSessionStats(r.Context(), userSession.Id)
	if err != nil {
		http.Error(w, "Error calculating session status", http.StatusInternalServerError)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(sessionStats)
}

// GetUserExerciseSessionStats godoc
// @Summary      Get user exercise data
// @Description  Get user exercise total data for durations and calories
// @Tags         activities
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Activity ID"
// @Success      200  {array}   ExerciseSession
// @Router       /activities/user/session/total/stats/{user_id} [post]
func (h Handler) GetUserExerciseSessionStats(w http.ResponseWriter, r *http.Request) {
	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	// Calculate and save the total exercise session data
	sessionStats, err := h.service.Activity.GetUserExerciseSessionStats(r.Context(), userSession.Id)
	if err != nil {
		http.Error(w, "Error calculating session status", http.StatusInternalServerError)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(sessionStats)
}
