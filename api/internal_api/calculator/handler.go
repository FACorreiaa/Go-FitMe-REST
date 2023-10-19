package calculator

import (
	"encoding/json"
	"errors"
	"github.com/FACorreiaa/Stay-Healthy-Backend/api/internal_api/auth"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"log"
	"math"
	"net/http"
	"time"
)

type Handler struct {
	service *StructCalculator
}

func NewCalculatorHandler(s *StructCalculator) *Handler {
	return &Handler{
		service: s,
	}
}

func mapActivity(activity Activity) (*ActivityList, error) {
	description, valid := activityDescriptionMap[activity]
	if !valid {
		return nil, errors.New("invalid activity")
	}
	return &ActivityList{
		Activity:    activity,
		Description: description,
	}, nil
}

func mapActivityValues(activity Activity) (ActivityValues, error) {
	value, _ := activityValuesMap[activity]
	return value, nil
}

func mapObjective(objective Objective) (*ObjectiveList, error) {
	description, valid := objectiveDescriptionMap[objective]
	if !valid {
		return nil, errors.New("invalid objective")
	}
	return &ObjectiveList{
		Objective:   objective,
		Description: description,
	}, nil
}

func mapDistribution(distribution CaloriesDistribution) (*CaloriesInfo, error) {
	description, valid := carbsDistribution[distribution]
	if !valid {
		return nil, errors.New("invalid objective")
	}
	return &CaloriesInfo{
		CaloriesDistribution:            distribution,
		CaloriesDistributionDescription: description,
	}, nil
}

func validateAge(age uint8) (uint8, error) {
	if age <= minAge || age >= maxAge {
		return 0, errors.New("invalid age")
	}
	return age, nil
}

func validateWeight(weight uint16) (uint16, error) {
	if weight <= minWeight || weight > maxWeight {
		return 0, errors.New("invalid weight")
	}

	return weight, nil
}

func validateHeight(height uint8) (uint8, error) {
	if height <= minHeight || height > maxHeight {
		return 0, errors.New("invalid height")
	}

	return height, nil
}

func convertWeight(weight uint16, system System) float64 {
	if system == metric {
		return float64(weight)
	}
	return float64(weight) / 0.453592 // 1 lb = 0.453592 kg
}
func convertHeight(height uint8, system System) float64 {
	if system == metric {
		return float64(height)
	}
	return float64(height) / 2.54 // 1 in = 2.54 cm
}

func calculateBMR(userData UserData, system System) float64 {
	var ageFactor float64
	weight := convertWeight(userData.Weight, system)
	height := convertHeight(userData.Height, system)
	if userData.Gender == m {
		ageFactor = maleAgeFactor
	} else {
		ageFactor = femaleAgeFactor
	}

	if system == metric {
		return math.Round((10*weight + 6.25*height - 5.0*(float64(userData.Age))) + ageFactor)
	} else {
		return math.Round((4.536*weight + 15.88*height - 5.0*(float64(userData.Age))) + ageFactor)
	}
}

func calculateTDEE(bmr float64, activityValue ActivityValues) float64 {
	return math.Round(bmr * float64(activityValue))
}

func calculateGoals(tdee float64) Goals {
	var fatLoss = tdee - caloricDeficit
	var bulk = tdee + caloricPlus
	return Goals{
		Cutting:     uint16(fatLoss),
		Maintenance: uint16(tdee),
		Bulking:     uint16(bulk),
	}
}

func getGoal(tdeeValue float64, objective Objective) uint16 {
	goals := calculateGoals(tdeeValue)
	mapGoals := make(map[Objective]uint16)
	mapGoals[maintenance] = goals.Maintenance
	mapGoals[cutting] = goals.Cutting
	mapGoals[bulking] = goals.Bulking
	return mapGoals[objective]
}

func calculateMacroNutrients(calorieGoal float64, distribution CaloriesDistribution) Macros {
	if ratios, ok := macroRatios[distribution]; ok {
		protein := calculateMacroDistribution(ratios.ProteinRatio, calorieGoal, proteinGramValue)
		fats := calculateMacroDistribution(ratios.FatRatio, calorieGoal, fatGramValue)
		carbs := calculateMacroDistribution(ratios.CarbRatio, calorieGoal, carbGramValue)

		return Macros{
			Protein: uint16(protein),
			Fats:    uint16(fats),
			Carbs:   uint16(carbs),
		}
	}

	return Macros{}
}

func calculateMacroDistribution(calorieFactor float64, calorieGoal float64, caloriesPerGram int) float64 {
	return math.Round((calorieFactor * calorieGoal) / float64(caloriesPerGram))
}

func validateUserInput(params UserParams) (UserData, error) {
	validAge, _ := validateAge(params.Age)
	validHeight, _ := validateHeight(params.Height)
	validWeight, _ := validateWeight(params.Weight)
	userInputData := UserData{
		Age:    validAge,
		Height: validHeight,
		Weight: validWeight,
		Gender: params.Gender,
	}
	return userInputData, nil
}

func calculateUserPersonalMacros(params UserParams) (UserInfo, error) {
	userData, err := validateUserInput(params)
	bmr := calculateBMR(userData, System(params.System))
	a, err := mapActivity(Activity(params.Activity))
	o, err := mapObjective(Objective(params.Objective))
	v, err := mapActivityValues(Activity(params.Activity))
	d, err := mapDistribution(CaloriesDistribution(params.CaloriesDist))
	tdee := calculateTDEE(bmr, v)
	goal := getGoal(tdee, Objective(params.Objective))
	if err != nil {
		return UserInfo{}, err
	}

	macros := calculateMacroNutrients(tdee, CaloriesDistribution(params.CaloriesDist))
	return UserInfo{
		System: params.System,
		UserData: UserData{
			Age:    userData.Age,
			Height: userData.Height,
			Weight: userData.Weight,
			Gender: userData.Gender,
		},
		ActivityInfo: ActivityInfo{
			Activity:    a.Activity,
			Description: a.Description,
		},
		ObjectiveInfo: ObjectiveInfo{
			Objective:   o.Objective,
			Description: o.Description,
		},
		BMR:  uint16(bmr),
		TDEE: uint16(tdee),
		MacrosInfo: MacrosInfo{
			CaloriesInfo: CaloriesInfo{
				CaloriesDistribution:            d.CaloriesDistribution,
				CaloriesDistributionDescription: d.CaloriesDistributionDescription,
			},
			Macros: macros,
		},
		Goal: goal,
	}, nil

}

func CalculateMacrosOffline(w http.ResponseWriter, r *http.Request) {
	var params UserParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	response, err := calculateUserPersonalMacros(params)
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

// CalculateMacros godoc
// @Summary      Create all diet macros from user
// @Description  Create all diet macros from user
// @Tags         macros calculator
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {array}   UserMacroDistribution
// @Router       /{user_id} [post]
func (h Handler) CalculateMacros(w http.ResponseWriter, r *http.Request) {
	var params UserParams
	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}

	macros, err := calculateUserPersonalMacros(params)
	response, err := h.service.Calculator.Create(UserMacroDistribution{
		ID:                              uuid.New(),
		UserID:                          userSession.Id,
		Age:                             params.Age,
		Height:                          params.Height,
		Weight:                          params.Weight,
		Gender:                          params.Gender,
		System:                          params.System,
		Activity:                        string(macros.ActivityInfo.Activity),
		ActivityDescription:             string(macros.ActivityInfo.Description),
		Objective:                       string(macros.ObjectiveInfo.Objective),
		ObjectiveDescription:            string(macros.ObjectiveInfo.Description),
		CaloriesDistribution:            string(macros.MacrosInfo.CaloriesInfo.CaloriesDistribution),
		CaloriesDistributionDescription: string(macros.MacrosInfo.CaloriesInfo.CaloriesDistributionDescription),
		Protein:                         macros.MacrosInfo.Macros.Protein,
		Fats:                            macros.MacrosInfo.Macros.Fats,
		Carbs:                           macros.MacrosInfo.Macros.Carbs,
		BMR:                             macros.BMR,
		TDEE:                            macros.TDEE,
		Goal:                            macros.Goal,
		CreatedAt:                       time.Now(),
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

// GetAllDietMacros godoc
// @Summary      Get all diet macros from user
// @Description  Get all diet macros from user
// @Tags         macros calculator
// @Accept       json
// @Produce      json
// @Param        id   path      int  true  "User ID"
// @Success      200  {array}   UserMacroDistribution
// @Router       /{user_id} [get]
func (h Handler) GetAllDietMacros(w http.ResponseWriter, r *http.Request) {
	userSession, ok := r.Context().Value(auth.SessionManagerKey{}).(*auth.UserSession)
	if !ok {
		http.Error(w, "User session not found", http.StatusUnauthorized)
		return
	}

	userMacros, err := h.service.Calculator.GetAll(r.Context(), userSession.Id)
	if err != nil {
		http.Error(w, "Error finding user plan", http.StatusInternalServerError)
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(userMacros)
}

// GetDietMacros godoc
// @Summary      Get diet macros
// @Description  Get diet macros
// @Tags         macros calculator
// @Accept       json
// @Produce      json
// @Param        id   path      string  true  "uuid formatted ID."
// @Success      200  {array}   UserMacroDistribution
// @Router       /plan/{id} [get]
func (h Handler) GetDietMacros(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Error parsing id", http.StatusInternalServerError)
	}

	userMacros, err := h.service.Calculator.Get(r.Context(), id)
	println("err: ", err)
	if err != nil {
		http.Error(w, "Error finding user plan", http.StatusInternalServerError)
	}

	// Serialize the response as JSON and write to the response writer
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(userMacros)
}
