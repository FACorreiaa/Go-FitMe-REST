package activity

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/auth"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

type activityHandler struct {
	logger          *logrus.Logger
	router          *chi.Router
	ctx             context.Context
	activityService *ActivityService
	sessionManager  *auth.SessionManager
}

func NewActivityHandler(lg *logrus.Logger, db *sqlx.DB, sessionManager *auth.SessionManager) *activityHandler {
	repo, err := NewActivityRepository(db)
	if err != nil {
		errors.New("error injecting activity service")
	}
	service := NewActivityService(repo)
	return &activityHandler{
		logger:          lg,
		activityService: service,
		sessionManager:  sessionManager,
	}
}

// GetActivities gets all existing activities

// @Summary      Get activities
// @Description  get activities
// @Tags         acthttps://www.youtube.com/watch?v=mYWllgYPaWsivities
// @Accept       json
// @Produce      json
// @Param        q    query     string  false  "name search by q"  Format(email)
// @Success      200  {array}   model.Activity
// @Router       /api/v1/activities [get]

func (a activityHandler) GetActivities(w http.ResponseWriter, r *http.Request) {
	activities, err := a.activityService.GetAll(a.ctx)

	if err != nil {
		log.Printf("Error fetching activities data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		_, _ = w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(activities)
}

func (s activityHandler) StartTracker(w http.ResponseWriter, r *http.Request)  {}
func (s activityHandler) StopTracker(w http.ResponseWriter, r *http.Request)   {}
func (s activityHandler) ResumeTracker(w http.ResponseWriter, r *http.Request) {}
