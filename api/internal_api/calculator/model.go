package calculator

import "time"

type UserMacroDistribution struct {
	ID                              string    `json:"id,string" db:"id" pg:"default:gen_random_uuid()"`
	UserID                          int       `json:"user_id" db:"user_id"`
	Age                             uint8     `json:"age" db:"age"`
	Height                          uint8     `json:"height" db:"height"`
	Weight                          uint16    `json:"weight" db:"weight"`
	Gender                          string    `json:"gender" db:"gender"`
	System                          string    `json:"system" db:"system"`
	Activity                        string    `json:"activity" db:"activity"`
	ActivityDescription             string    `json:"activity_description" db:"activity_description"`
	Objective                       string    `json:"objective" db:"objective"`
	ObjectiveDescription            string    `json:"objective_description" db:"objective_description"`
	CaloriesDistribution            string    `json:"calories_distribution" db:"calories_distribution"`
	CaloriesDistributionDescription string    `json:"calories_distribution_description" db:"calories_distribution_description"`
	Protein                         uint16    `json:"protein" db:"protein"`
	Fats                            uint16    `json:"fats" db:"fats"`
	Carbs                           uint16    `json:"carbs" db:"carbs"`
	BMR                             uint16    `json:"bmr" db:"bmr"`
	TDEE                            uint16    `json:"tdee" db:"tdee"`
	Goal                            uint16    `json:"goal" db:"goal"`
	CreatedAt                       time.Time `json:"created_at" db:"created_at"`
}

type UserInfo struct {
	System        string `json:"system" db:"system"`
	UserData      UserData
	ActivityInfo  ActivityInfo
	ObjectiveInfo ObjectiveInfo
	BMR           uint16 `json:"bmr" db:"bmr"`
	TDEE          uint16 `json:"tdee" db:"tdee"`
	MacrosInfo    MacrosInfo
	Goal          uint16 `json:"dietGoal" db:"goal"`
}

type UserParams struct {
	Age          uint8  `json:"age" db:"age"`
	Height       uint8  `json:"height" db:"height"`
	Weight       uint16 `json:"weight" db:"weight"`
	Gender       string `json:"gender" db:"gender"`
	System       string `json:"system" db:"system"`
	Activity     string `json:"activity" db:"activity"`
	Objective    string `json:"objective" db:"objective"`
	CaloriesDist string `json:"calories-distribution" db:"calorie_distribution"`
}

type Goals struct {
	Bulking     uint16 `json:"bulking"`
	Cutting     uint16 `json:"cutting"`
	Maintenance uint16 `json:"maintenance"`
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
	Age    uint8  `json:"age"`
	Height uint8  `json:"height"`
	Weight uint16 `json:"weight"`
	Gender string `json:"gender"`
}

type ActivityInfo struct {
	Activity    Activity            `json:"activity" db:"activity"`
	Description ActivityDescription `json:"description" db:"activity_description"`
}
type ObjectiveInfo struct {
	Objective   Objective            `json:"objective" db:"objective"`
	Description ObjectiveDescription `json:"description" db:"objective_description"`
}

type Macros struct {
	Protein uint16 `json:"protein"`
	Fats    uint16 `json:"fats"`
	Carbs   uint16 `json:"carbs"`
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
	Bulking     uint16 `json:"bulking"`
	Cutting     uint16 `json:"cutting"`
	Maintenance uint16 `json:"maintenance"`
}

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
