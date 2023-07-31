package calculator

import (
	"encoding/json"
	"errors"
	"log"
	"math"
	"net/http"
)

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

func validateAge(age int64) (int64, error) {
	if age <= minAge || age >= maxAge {
		return 0, errors.New("invalid age")
	}
	return age, nil
}

func validateWeight(weight float64) (float64, error) {
	if weight <= minWeight || weight > maxWeight {
		return 0, errors.New("invalid weight")
	}

	return weight, nil
}

func validateHeight(height float64) (float64, error) {
	if height <= minHeight || height > maxHeight {
		return 0, errors.New("invalid height")
	}

	return height, nil
}

func convertWeight(weight float64, system System) float64 {
	if system == metric {
		return weight
	}
	return weight / 0.453592 // 1 lb = 0.453592 kg
}
func convertHeight(height float64, system System) float64 {
	if system == metric {
		return height
	}
	return height / 2.54 // 1 in = 2.54 cm
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
		Cutting:     fatLoss,
		Maintenance: tdee,
		Bulking:     bulk,
	}
}

func getGoal(tdeeValue float64, objective Objective) float64 {
	goals := calculateGoals(tdeeValue)
	mapGoals := make(map[Objective]float64)
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
			Protein: protein,
			Fats:    fats,
			Carbs:   carbs,
		}
	}

	return Macros{}
}

func calculateMacroDistribution(calorieFactor float64, calorieGoal float64, caloriesPerGram int) float64 {
	return math.Round((calorieFactor * calorieGoal) / float64(caloriesPerGram))
}

func CalculateMacros(w http.ResponseWriter, r *http.Request) {
	var params UserParams
	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		http.Error(w, "Bad Request", http.StatusBadRequest)
		return
	}
	validAge, _ := validateAge(params.Age)
	validHeight, _ := validateHeight(params.Height)
	validWeight, _ := validateWeight(params.Weight)
	userInputData := UserData{
		Age:    params.Age,
		Height: params.Height,
		Weight: params.Weight,
		Gender: params.Gender,
	}
	bmr := calculateBMR(userInputData, System(params.System))
	a, err := mapActivity(Activity(params.Activity))
	o, err := mapObjective(Objective(params.Objective))
	v, err := mapActivityValues(Activity(params.Activity))
	d, err := mapDistribution(CaloriesDistribution(params.CaloriesDist))
	tdee := calculateTDEE(bmr, v)
	goal := getGoal(tdee, Objective(params.Objective))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	macros := calculateMacroNutrients(tdee, CaloriesDistribution(params.CaloriesDist))
	response := UserInfo{
		System: params.System,
		UserData: UserData{
			Age:    validAge,
			Height: validHeight,
			Weight: validWeight,
			Gender: params.Gender,
		},
		ActivityInfo: ActivityInfo{
			Activity:    a.Activity,
			Description: a.Description,
		},
		ObjectiveInfo: ObjectiveInfo{
			Objective:   o.Objective,
			Description: o.Description,
		},
		BMR:  bmr,
		TDEE: tdee,
		MacrosInfo: MacrosInfo{
			CaloriesInfo: CaloriesInfo{
				CaloriesDistribution:            d.CaloriesDistribution,
				CaloriesDistributionDescription: d.CaloriesDistributionDescription,
			},
			Macros: macros,
		},
		Goal: goal,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
