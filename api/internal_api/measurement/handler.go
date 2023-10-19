package measurement

import (
	"encoding/json"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/auth"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"log"
	"net/http"
	"time"
)

type Handler struct {
	service *StructMeasurement
}

func NewMeasurementHandler(s *StructMeasurement) *Handler {
	return &Handler{
		service: s,
	}
}

// weight

// InsertWeight godoc
// @Summary      Insert user weight
// @Description  Insert user weight
// @Tags         measurements weight
// @Accept       json
// @Produce      json
// @Param        w   path      int  true  "Weight"
// @Success      200  {object}   Weight
// @Router       /weight [post]
func (h Handler) InsertWeight(w http.ResponseWriter, r *http.Request) {
	var input struct {
		WeightValue float32 `json:"weight_value"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	response, err := h.service.Measurement.InsertWeight(Weight{
		ID:          uuid.New(),
		UserID:      userSession.Id,
		WeightValue: input.WeightValue,
		CreatedAt:   time.Now(),
		UpdatedAt:   nil,
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

// UpdateWeight godoc
// @Summary      Update user weight
// @Description  Update user weight
// @Tags         measurements weight
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Weight ID"
// @Success      200  {array}   Weight
// @Router       /weight/{id} [patch]
func (h Handler) UpdateWeight(w http.ResponseWriter, r *http.Request) {
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

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updates["UpdatedAt"] = time.Now()

	if err := h.service.Measurement.UpdateWeight(id, userSession.Id, updates); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// DeleteWeight godoc
// @Summary      Delete user weight
// @Description  Delete user weight
// @Tags         measurements weight
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Weight ID"
// @Success      200  {array}   Weight
// @Router       /weight/{id} [delete]
func (h Handler) DeleteWeight(w http.ResponseWriter, r *http.Request) {
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

	err = h.service.Measurement.DeleteWeight(id, userSession.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// GetWeight godoc
// @Summary      Get user weight
// @Description  Get user weight
// @Tags         measurements weight
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Weight"
// @Success      200  {array}   Weight
// @Router       /weight/{id} [get]
func (h Handler) GetWeight(w http.ResponseWriter, r *http.Request) {
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

	weight, err := h.service.Measurement.GetWeight(id, userSession.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weight)
}

// GetWeights godoc
// @Summary      Get all user weight list
// @Description  Get all user weight list
// @Tags         measurements weight
// @Accept       json
// @Produce      json
// @Success      200  {array}   Weight
// @Router       /weights [get]
func (h Handler) GetWeights(w http.ResponseWriter, r *http.Request) {
	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	weight, err := h.service.Measurement.GetWeights(userSession.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weight)
}

//water intake

// InsertWaterIntake godoc
// @Summary      Insert user water intake
// @Description  Insert user water intake
// @Tags         measurements water
// @Accept       json
// @Produce      json
// @Param        w   path      int  true  "Water"
// @Success      200  {object}   WaterIntake
// @Router       /water [post]
func (h Handler) InsertWaterIntake(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Quantity float32 `json:"quantity"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	response, err := h.service.Measurement.InsertWaterIntake(WaterIntake{
		ID:        uuid.New(),
		UserID:    userSession.Id,
		Quantity:  input.Quantity,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
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

// UpdateWaterIntake godoc
// @Summary      Update user water intake
// @Description  Update user water intake
// @Tags         measurements water
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Water"
// @Success      200  {object}   WaterIntake
// @Router       /water/{id} [patch]
func (h Handler) UpdateWaterIntake(w http.ResponseWriter, r *http.Request) {
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

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updates["UpdatedAt"] = time.Now()

	if err := h.service.Measurement.UpdateWaterIntake(id, userSession.Id, updates); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// DeleteWaterIntake godoc
// @Summary      Delete user water intake
// @Description  Delete user water intake
// @Tags         measurements water
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Water"
// @Success      200  {object}   WaterIntake
// @Router       /water/{id} [delete]
func (h Handler) DeleteWaterIntake(w http.ResponseWriter, r *http.Request) {
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

	err = h.service.Measurement.DeleteWaterIntake(id, userSession.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// GetWaterIntake godoc
// @Summary      Get user water intake
// @Description  Get user water intake
// @Tags         measurements water
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "Water"
// @Success      200  {object}   WaterIntake
// @Router       /water/{id} [get]
func (h Handler) GetWaterIntake(w http.ResponseWriter, r *http.Request) {
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

	weight, err := h.service.Measurement.GetWaterIntake(id, userSession.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weight)
}

// GetWaterIntakes godoc
// @Summary      Get user water intake list
// @Description  Get user water intake list
// @Tags         measurements water
// @Accept       json
// @Produce      json
// @Success      200  {object}   WaterIntake
// @Router       /water [get]
func (h Handler) GetWaterIntakes(w http.ResponseWriter, r *http.Request) {
	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	weight, err := h.service.Measurement.GetWaterIntakes(userSession.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(weight)
}

//waist line

// InsertWaistLine godoc
// @Summary      Get user waist line
// @Description  Get user waist line
// @Tags         measurements waistline
// @Accept       json
// @Produce      json
// @Success      200  {object}   WaistLine
// @Router       /waistline [post]
func (h Handler) InsertWaistLine(w http.ResponseWriter, r *http.Request) {
	var input struct {
		Quantity float32 `json:"quantity"`
	}

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	response, err := h.service.Measurement.InsertWaistLine(WaistLine{
		ID:        uuid.New(),
		UserID:    userSession.Id,
		Quantity:  input.Quantity,
		CreatedAt: time.Now(),
		UpdatedAt: nil,
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

// UpdateWaistLine godoc
// @Summary      Update user waist line
// @Description  Update user waist line
// @Tags         measurements waistline
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "WaistLine"
// @Success      200  {object}   WaistLine
// @Router       /waistline/{id} [patch]
func (h Handler) UpdateWaistLine(w http.ResponseWriter, r *http.Request) {
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

	var updates map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	updates["UpdatedAt"] = time.Now()

	if err := h.service.Measurement.UpdateWaistLine(id, userSession.Id, updates); err != nil {
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// DeleteWaistLine godoc
// @Summary      Delete user waist line
// @Description  Delete user waist line
// @Tags         measurements waistline
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "WaistLine"
// @Success      200  {object}   WaistLine
// @Router       /waistline/{id} [delete]
func (h Handler) DeleteWaistLine(w http.ResponseWriter, r *http.Request) {
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

	err = h.service.Measurement.DeleteWaistLine(id, userSession.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
}

// GetWaistLine godoc
// @Summary      Get user waist line
// @Description  Get user waist line
// @Tags         measurements waistline
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "WaistLine"
// @Success      200  {object}   WaistLine
// @Router       /waistline/{id} [get]
func (h Handler) GetWaistLine(w http.ResponseWriter, r *http.Request) {
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

	res, err := h.service.Measurement.GetWaistLine(id, userSession.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}

// GetWaistLines godoc
// @Summary      Get user waist line list
// @Description  Get user waist line list
// @Tags         measurements waistline
// @Accept       json
// @Produce      json
// @Success      200  {object}   WaistLine
// @Router       /waistline [get]
func (h Handler) GetWaistLines(w http.ResponseWriter, r *http.Request) {
	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	res, err := h.service.Measurement.GetWaterIntakes(userSession.Id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(res)
}
