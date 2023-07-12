package calculator

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"log"
	"math"
	"net/http"
	"strconv"
)

type Objective string
type ObjectiveDescription string
type CaloriesDistribution string
type CaloriesDistributionDescription string
type Activity string
type ActivityDescription string
type ActivityValues float64
type System string

const (
	maintenance Objective = "Maintenance"
	bulking     Objective = "Bulking"
	cutting     Objective = "Cutting"

	m             = "Male"
	metric System = "Metric"

	sedentaryActivity  Activity = "Sedentary"
	lightActivity      Activity = "LightActivity"
	moderateActivity   Activity = "Moderate"
	heavyActivity      Activity = "Heavy"
	extraHeavyActivity Activity = "ExtraHeavy"

	sedentaryActivityDescription  ActivityDescription = "Office Job. Very low activity during the day"
	lightActivityDescription      ActivityDescription = "Workout 1 to 2 days per week"
	moderateActivityDescription   ActivityDescription = "Workout 3-5 days per week"
	heavyActivityDescription      ActivityDescription = "Workout 5 to 7 days per week"
	extraHeavyActivityDescription ActivityDescription = "Giga Dog! Training twice a day!"

	maintenanceDescription ObjectiveDescription = "You choose to keep your current weight."
	cuttingDescription     ObjectiveDescription = "You choose to lose some weight. Ideally 250 grams or half a pound per week"
	bulkingDescription     ObjectiveDescription = "You choose to gain weight. Ideally 300 grams or 0.70 pounds per week"

	lowCarbs      CaloriesDistributionDescription = "Low Carbs diet. Protein 0.4, Fats 0.4, Carbs 0.2"
	moderateCarbs CaloriesDistributionDescription = "Moderate Carbs diet. Protein 0.3, Fats 0.35, Carbs 0.35"
	highCarbs     CaloriesDistributionDescription = "High Carbs diet. Protein 0.3, Fats 0.2, Carbs 0.5"
)

const (
	minAge                                 = 0
	maxAge                                 = 100
	minHeight                              = 0
	maxHeight                              = 250
	minWeight                              = 0
	maxWeight                              = 500
	maleAgeFactor                          = 5
	femaleAgeFactor                        = -161
	caloricDeficit                         = 450.0
	caloricPlus                            = 350.0
	sedentaryActivityValue  ActivityValues = 1.2
	lightActivityValue      ActivityValues = 1.375
	moderateActivityValue   ActivityValues = 1.55
	heavyActivityValue      ActivityValues = 1.725
	extraHeavyActivityValue ActivityValues = 1.9
)

const (
	HighCarbRatios     CaloriesDistribution = "High"
	ModerateCarbRatios CaloriesDistribution = "Moderate"
	LowCarbRatios      CaloriesDistribution = "Low"
)

var macroRatios = map[CaloriesDistribution]struct {
	ProteinRatio float64
	FatRatio     float64
	CarbRatio    float64
}{
	HighCarbRatios: {
		ProteinRatio: 0.3,
		FatRatio:     0.2,
		CarbRatio:    0.5,
	},
	ModerateCarbRatios: {
		ProteinRatio: 0.3,
		FatRatio:     0.35,
		CarbRatio:    0.35,
	},
	LowCarbRatios: {
		ProteinRatio: 0.4,
		FatRatio:     0.4,
		CarbRatio:    0.2,
	},
}

var (
	fatGramValue     = 9
	proteinGramValue = 4
	carbGramValue    = 4
)

type Goals struct {
	Bulking     float64 `json:"bulking"`
	Cutting     float64 `json:"cutting"`
	Maintenance float64 `json:"maintenance"`
}

type ActivityList struct {
	Activity    Activity            `json:"activity"`
	Description ActivityDescription `json:"description"`
}

type ObjectiveList struct {
	Objective   Objective            `json:"objective"`
	Description ObjectiveDescription `json:"description"`
}

type SystemList struct {
	System System `json:"metric"`
}

type UserData struct {
	Age    int64   `json:"age"`
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
	Gender string  `json:"gender"`
}

type ActivityInfo struct {
	Activity    Activity            `json:"activity"`
	Description ActivityDescription `json:"description"`
}
type ObjectiveInfo struct {
	Objective   Objective            `json:"objective"`
	Description ObjectiveDescription `json:"description"`
}

type Macros struct {
	Protein float64 `json:"protein"`
	Fats    float64 `json:"fats"`
	Carbs   float64 `json:"carbs"`
}

type CaloriesInfo struct {
	CaloriesDistribution            CaloriesDistribution            `json:"carbDistribution"`
	CaloriesDistributionDescription CaloriesDistributionDescription `json:"carbDistributionDescription"`
}
type MacrosInfo struct {
	CaloriesInfo CaloriesInfo
	Macros       Macros
}

type CaloriesObjective struct {
	Bulking     float64 `json:"bulking"`
	Cutting     float64 `json:"cutting"`
	Maintenance float64 `json:"maintenance"`
}

type UserInfo struct {
	System        string `json:"system"`
	UserData      UserData
	ActivityInfo  ActivityInfo
	ObjectiveInfo ObjectiveInfo
	BMR           float64 `json:"bmr"`
	TDEE          float64 `json:"tdee"`
	MacrosInfo    MacrosInfo
	Goal          float64 `json:"caloricGoal"`
}

var activityDescriptionMap = map[Activity]ActivityDescription{
	sedentaryActivity:  sedentaryActivityDescription,
	lightActivity:      lightActivityDescription,
	moderateActivity:   moderateActivityDescription,
	heavyActivity:      heavyActivityDescription,
	extraHeavyActivity: extraHeavyActivityDescription,
}
var activityValuesMap = map[Activity]ActivityValues{
	sedentaryActivity:  sedentaryActivityValue,
	lightActivity:      lightActivityValue,
	moderateActivity:   moderateActivityValue,
	heavyActivity:      heavyActivityValue,
	extraHeavyActivity: extraHeavyActivityValue,
}
var objectiveDescriptionMap = map[Objective]ObjectiveDescription{
	maintenance: maintenanceDescription,
	bulking:     bulkingDescription,
	cutting:     cuttingDescription,
}

var carbsDistribution = map[CaloriesDistribution]CaloriesDistributionDescription{
	HighCarbRatios:     highCarbs,
	LowCarbRatios:      lowCarbs,
	ModerateCarbRatios: moderateCarbs,
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
	age, _ := strconv.ParseInt(chi.URLParam(r, "age"), 10, 64)
	height, _ := strconv.ParseFloat(chi.URLParam(r, "height"), 64)
	weight, _ := strconv.ParseFloat(chi.URLParam(r, "weight"), 64)
	gender := chi.URLParam(r, "gender")
	system := chi.URLParam(r, "system")
	activity := chi.URLParam(r, "activity")
	objective := chi.URLParam(r, "objective")
	distribution := chi.URLParam(r, "calories-distribution")
	validAge, _ := validateAge(age)
	validHeight, _ := validateHeight(height)
	validWeight, _ := validateWeight(weight)
	userInputData := UserData{
		Age:    age,
		Height: height,
		Weight: weight,
		Gender: gender,
	}
	bmr := calculateBMR(userInputData, System(system))
	a, err := mapActivity(Activity(activity))
	o, err := mapObjective(Objective(objective))
	v, err := mapActivityValues(Activity(activity))
	d, err := mapDistribution(CaloriesDistribution(distribution))
	tdee := calculateTDEE(bmr, v)
	goal := getGoal(tdee, Objective(objective))

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	macros := calculateMacroNutrients(tdee, CaloriesDistribution(distribution))
	response := UserInfo{
		System: system,
		UserData: UserData{
			Age:    validAge,
			Height: validHeight,
			Weight: validWeight,
			Gender: gender,
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
