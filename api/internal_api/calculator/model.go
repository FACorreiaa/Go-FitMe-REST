package calculator

type UserParams struct {
	Age          int64   `json:"age"`
	Height       float64 `json:"height"`
	Weight       float64 `json:"weight"`
	Gender       string  `json:"gender"`
	System       string  `json:"system"`
	Activity     string  `json:"activity"`
	Objective    string  `json:"objective"`
	CaloriesDist string  `json:"calories_distribution"`
}

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
