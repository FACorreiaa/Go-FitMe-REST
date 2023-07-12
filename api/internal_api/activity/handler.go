package activity

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

type service struct {
	logger          *logrus.Logger
	router          *chi.Router
	ctx             context.Context
	activityService Service
}

func NewActivityHandler(lg *logrus.Logger, db *sqlx.DB) service {
	repo, err := NewRepository(db)
	if err != nil {
		err := errors.New("error creating activity handler")
		if err != nil {
			return service{}
		}
	}
	return service{
		logger:          lg,
		activityService: NewService(*repo),
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

func (s service) GetActivities(w http.ResponseWriter, r *http.Request) {
	activities, err := s.activityService.GetAll(s.ctx)

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

func (s service) StartTracker(w http.ResponseWriter, r *http.Request)  {}
func (s service) StopTracker(w http.ResponseWriter, r *http.Request)   {}
func (s service) ResumeTracker(w http.ResponseWriter, r *http.Request) {}
