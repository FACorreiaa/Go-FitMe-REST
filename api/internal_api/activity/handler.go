package activity

import (
	"context"
	"encoding/json"
	errors "errors"
	"fmt"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/auth"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
	"strconv"
	"time"
)

type Handler struct {
	logger           *logrus.Logger
	router           *chi.Router
	ctx              context.Context
	activityService  *ActivityService
	sessionManager   *auth.SessionManager
	exerciseSessions map[string]*ExerciseSession // Map to store exercise sessions for each user
	pausedTimers     map[string]time.Time
}

func NewActivityHandler(lg *logrus.Logger, db *sqlx.DB, sessionManager *auth.SessionManager) *Handler {
	repo, err := NewActivityRepository(db)
	if err != nil {
		_ = errors.New("error injecting activity service")
	}
	service := NewActivityService(repo)
	return &Handler{
		logger:           lg,
		activityService:  service,
		sessionManager:   sessionManager,
		ctx:              context.Background(),
		exerciseSessions: make(map[string]*ExerciseSession),
		pausedTimers:     make(map[string]time.Time),
	}
}

//func convertTimeAndCaloriesBurned(exerciseSession *ExerciseSession, activity *Activity) ExerciseSession {
//	// Calculate calories burned based on the duration
//	hours := float32(exerciseSession.Duration) / 60.0
//	if hours < 1.0 {
//		// If the duration is less than an hour, save it in minutes and calculate calories burned accordingly
//		exerciseSession.DurationMinutes = exerciseSession.Duration
//		exerciseSession.Duration = 0
//		exerciseSession.CaloriesBurned = int(activity.CaloriesPerHour * hours)
//	} else {
//		// If the duration is one hour or more, save it in hours and minutes and calculate calories burned accordingly
//		exerciseSession.DurationMinutes = exerciseSession.Duration % 60
//		exerciseSession.Duration /= 60
//		exerciseSession.CaloriesBurned = int(activity.CaloriesPerHour * float32(exerciseSession.Duration))
//	}
//}

// GetActivities gets all existing activities
// @Summary      Get activities
// @Description  get activities
// @Tags         acthttps://www.youtube.com/watch?v=mYWllgYPaWsivities
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200  {array}   model.Activity
// @Router       /api/v1/activities [get]
func (a Handler) GetActivities(w http.ResponseWriter, r *http.Request) {
	activities, err := a.activityService.GetAll(a.ctx)

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

// GetActivitiesByName gets all existing activities by name
// @Summary      GetActivitiesByName
// @Description  gets all existing activities by name
// @Tags
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by name"
// @Success      200  {array}   model.Activity
// @Router       /api/v1/activities/name={name} [get]
func (a Handler) GetActivitiesByName(w http.ResponseWriter, r *http.Request) {
	name := chi.URLParam(r, "name")
	activities, err := a.activityService.GetByName(a.ctx, name)
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

// GetActivitiesById gets all existing activities by name
// @Summary      GetActivitiesById
// @Description  gets all existing activities by name
// @Tags
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by id"
// @Success      200  {array}   model.Activity
// @Router       /api/v1/activities/id={id} [get]
func (a Handler) GetActivitiesById(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		log.Printf("Error parsing id: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)

		return
	}
	activities, err := a.activityService.GetByID(a.ctx, id)
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

// StartActivityTracker start activity tracker

func (a Handler) StartActivityTracker(w http.ResponseWriter, r *http.Request) {
	activityID, err := strconv.Atoi(chi.URLParam(r, "id"))
	currentTime := time.Now()
	if err != nil {
		log.Printf("Error parsing id: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}
	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	sessionId := strconv.Itoa(userSession.Id)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	_, found := a.exerciseSessions[sessionId]
	if found {
		http.Error(w, "Exercise session already in progress", http.StatusNotFound)
		return
	}

	if err != nil {
		log.Printf("Error getting session: %v", err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	activity, err := a.activityService.GetByID(a.ctx, activityID)
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

	a.exerciseSessions[sessionId] = exerciseSession

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(exerciseSession)
}

func (a Handler) PauseActivityTracker(w http.ResponseWriter, r *http.Request) {
	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	sessionId := strconv.Itoa(userSession.Id)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	a.pausedTimers[sessionId] = time.Now()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(userSession)

}
func (a Handler) ResumeActivityTracker(w http.ResponseWriter, r *http.Request) {
	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	sessionId := strconv.Itoa(userSession.Id)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	session, found := a.exerciseSessions[sessionId]
	if !found {
		http.Error(w, "Exercise session not found", http.StatusNotFound)
		return
	}

	// Calculate the duration of the paused state and update the start time
	pausedDuration := time.Since(a.pausedTimers[sessionId])
	session.StartTime = session.StartTime.Add(pausedDuration)

	// Clear the paused timer for the user
	delete(a.pausedTimers, sessionId)

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(session)

}
func (a Handler) StopActivityTracker(w http.ResponseWriter, r *http.Request) {
	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	sessionId := strconv.Itoa(userSession.Id)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}
	// Get the exercise session for the user
	session, found := a.exerciseSessions[sessionId]
	if !found {
		http.Error(w, "Exercise session not found", http.StatusNotFound)
		return
	}

	// Calculate calories burned (assuming activity is fetched from the database)
	activity, err := a.activityService.GetByID(a.ctx, session.ActivityID)
	if err != nil {
		http.Error(w, "Error getting activity", http.StatusInternalServerError)
		return
	}

	//totalDurationSeconds := session.DurationHours*3600 + session.DurationMinutes*60 + session.DurationSeconds
	startUpTime := a.exerciseSessions[sessionId].StartTime
	totalDurationSeconds := int(time.Since(startUpTime).Seconds())

	session.DurationHours = totalDurationSeconds / 3600
	session.DurationMinutes = (totalDurationSeconds % 3600) / 60
	session.DurationSeconds = totalDurationSeconds % 60

	// Calculate calories burned per second
	caloriesPerSecond := activity.CaloriesPerHour / 3600

	// Calculate calories burned for the total duration in seconds
	session.CaloriesBurned = int(caloriesPerSecond * float32(totalDurationSeconds))

	session.EndTime = time.Now()

	err = a.activityService.SaveExerciseSession(a.ctx, session)
	if err != nil {
		http.Error(w, "Error saving exercise session to DB", http.StatusInternalServerError)
		return
	}

	delete(a.exerciseSessions, sessionId)

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(session)
}
