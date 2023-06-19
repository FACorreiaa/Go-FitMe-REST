package calculator

import (
	"encoding/json"
	"errors"
	"github.com/go-chi/chi/v5"
	"log"
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
type Measure string

const (
	maintenance Objective = "Maintenance"
	bulking     Objective = "Bulking"
	cutting     Objective = "Cutting"
)

const (
	m = "Male"
	f = "Female"
)

const (
	minAge                                 = 0
	maxAge                                 = 100
	minHeight                              = 50
	maxHeight                              = 250
	minWeight                              = 0
	maxWeight                              = 200
	maleAgeFactor                          = 5
	femaleAgeFactor                        = -161
	caloricDeficit                         = 450.0
	caloricExcedent                        = 350.0
	sedentaryActivityValue  ActivityValues = 1.2
	lightActivityValue      ActivityValues = 1.375
	moderateActivityValue   ActivityValues = 1.55
	heavyActivityValue      ActivityValues = 1.725
	extraHeavyActivityValue ActivityValues = 1.9
)

const (
	metric   Measure = "Metric"
	imperial Measure = "Imperial"
)

const (
	sedentaryActivity  Activity = "Sedentary"
	lightActivity      Activity = "Light Activity"
	moderateActivity   Activity = "Moderate"
	heavyActivity      Activity = "Heavy"
	extraHeavyActivity Activity = "Extra Heavy"
)
const (
	sedentaryActivityDescription  ActivityDescription = "Office Job. Very low activity during the day"
	lightActivityDescription      ActivityDescription = "Workout 1 to 2 days per week"
	moderateActivityDescription   ActivityDescription = "Workout 3-5 days per week"
	heavyActivityDescription      ActivityDescription = "Workout 5 to 7 days per week"
	extraHeavyActivityDescription ActivityDescription = "Giga Dog! Training twice a day!"
)

const (
	maintenanceDescription ObjectiveDescription = "You choose to keep your current weight."
	cuttingDescription     ObjectiveDescription = "You choose to loose some weight. Ideally 250 grams or half a pound per week"
	bulkingDescription     ObjectiveDescription = "You choose to gain weight. Ideally 300 grams or 0.70 pounds per week"
)

const (
	lowCarbs    CaloriesDistributionDescription = "Low Carbs diet. Protein 0.4, Fats 0.4, Carbs 0.2"
	mediumCarbs CaloriesDistributionDescription = "Moderate Carbs diet. Protein 0.3, Fats 0.35, Carbs 0.35"
	highCarbs   CaloriesDistributionDescription = "High Carbs diet. Protein 0.3, Fats 0.2, Carbs 0.5"
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

type MeasureList struct {
	Measure Measure `json:"metric"`
}

//const Low_Carb string = "Low Carb"
//const Moderate_Carb string = "Moderate Carb"
//const High_Carb string = "High Carb"
//
//const Metric_Height string = "cm"
//const Metric_Weight string = "kg"
//const Imperial_Height string = "lb"
//const Imperial_Weight string = "ft"
//
//const Male = "male"
//const Female = "female"

type UserData struct {
	Age    int64   `json:"age"`
	Height float64 `json:"height"`
	Weight float64 `json:"weight"`
	Gender string  `json:"gender"`
}

type Response struct {
	TDEE     float64 `json:"tdee"`
	Goal     float64 `json:"goal"`
	Activity string  `json:"activity"`
	Macros   Macros  `json:"macros"`
	BMF      float64 `json:"bmf"`
}

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

type MacrosInfo struct {
	CaloriesDistribution CaloriesDistribution `json:"calories-distribution"`
	Macros               Macros
}

type CaloriesObjective struct {
	Bulking     float64 `json:"bulking"`
	Cutting     float64 `json:"cutting"`
	Maintenance float64 `json:"maintenance"`
}

type GoalResponse struct {
	Goal float64 `json:"goal"`
}

type UserInfo struct {
	Metric        string `json:"metric"`
	UserData      UserData
	ActivityInfo  ActivityInfo
	ObjectiveInfo ObjectiveInfo
	BMR           float64 `json:"bmr"`
	TDEE          float64 `json:"tdee"`
	MacrosInfo    MacrosInfo
	Goal          float64 `json:"goal"`

	CaloriesObjective CaloriesObjective `json:"calories-objective"`
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

//var caloricDescriptionMap = map[CaloriesDistribution]CaloriesDistributionDescription{
//	LowCarbRatios:      lowCarbs,
//	ModerateCarbRatios: mediumCarbs,
//	HighCarbRatios:     highCarbs,
//}

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

func calculateBMR(userData UserData, measuringSystem Measure) float64 {
	var ageFactor float64
	if userData.Gender == m {
		ageFactor = maleAgeFactor
	} else {
		ageFactor = femaleAgeFactor
	}

	if measuringSystem == metric {
		return (10*userData.Weight + 6.25*userData.Height - 5.0*(float64(userData.Age))) + ageFactor
	} else {
		return (4.536*userData.Weight + 15.88*userData.Height - 5.0*(float64(userData.Age))) + ageFactor
	}
}

func calculateTDEE(bmr float64, activityValue ActivityValues) float64 {
	return bmr * float64(activityValue)
}

func calculateMacroDistribution(calorieFactor float64, calorieGoal float64, caloriesPerGram int) float64 {
	return (calorieFactor * calorieGoal) / float64(caloriesPerGram)
}

func calculateGoals(tdee float64) Goals {
	var fatLoss = tdee - caloricDeficit
	var bulk = tdee + caloricExcedent
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

//func getGoal(tdeeValue float64, objective Objective) GoalResponse {
//	goals := calculateGoals(tdeeValue)
//	mapGoals := make(map[Objective]float64)
//	mapGoals[maintenance] = goals.Maintenance
//	mapGoals[cutting] = goals.Cutting
//	mapGoals[bulking] = goals.Bulking
//	return GoalResponse{
//		Goal: mapGoals[objective],
//	}
//}

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

func CalculateMacros(w http.ResponseWriter, r *http.Request) {
	age, _ := strconv.ParseInt(chi.URLParam(r, "age"), 10, 64)
	height, _ := strconv.ParseFloat(chi.URLParam(r, "height"), 64)
	weight, _ := strconv.ParseFloat(chi.URLParam(r, "weight"), 64)
	gender := chi.URLParam(r, "gender")
	metric := chi.URLParam(r, "metric")
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
	bmr := calculateBMR(userInputData, Measure(metric))
	a, err := mapActivity(Activity(activity))
	o, err := mapObjective(Objective(objective))
	v, err := mapActivityValues(Activity(activity))
	tdee := calculateTDEE(bmr, v)
	//show struct after
	goal := getGoal(tdee, Objective(objective))
	println(goal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	macros := calculateMacroNutrients(tdee, CaloriesDistribution(distribution))
	response := UserInfo{
		Metric: metric,
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
			CaloriesDistribution: CaloriesDistribution(distribution),
			Macros:               macros,
		},
		Goal: goal,

		//CaloriesObjective: CaloriesObjective{
		//	Bulking:     macrosBulking,
		//	Cutting:     macrosCutting,
		//	Maintenance: macrosMaintenance,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}
