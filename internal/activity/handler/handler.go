package handler

import (
	"context"
	"encoding/json"
	"github.com/FACorreiaa/Stay-Healthy-Backend/internal/activity/repository"
	service2 "github.com/FACorreiaa/Stay-Healthy-Backend/internal/activity/service"
	"github.com/go-chi/chi/v5"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"log"
	"net/http"
)

type service struct {
	logger          *logrus.Logger
	router          *chi.Router
	activityService service2.Service
	ctx             context.Context
}

func NewHandler(lg *logrus.Logger, db *sqlx.DB) service {
	return service{
		logger:          lg,
		activityService: service2.NewService(repository.NewRepository(db)),
	}
}

// GetActivities func gets all existing activities
// @Description Get all activities
// @Summary Get all activities
// @Tags activities
// @Accept json
// @Produce json
// @Success 200 {array} domain.Activity
// @Router /api/v1/activity [get]

func (s service) GetActivities(w http.ResponseWriter, r *http.Request) {
	activities, err := s.activityService.GetAll(s.ctx)

	if err != nil {
		log.Printf("Error fetching activities data: %v", err)

		// Write an error response to the client
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal server error"))
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(activities)
}
