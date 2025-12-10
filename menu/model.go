package menu

type MenuItem struct {
	ItemName        string  `json:"item_name"`
	Category        string  `json:"category"`
	Calories        int     `json:"calories"`
	TasteProfile    string  `json:"taste_profile"`
	PopularityScore float64 `json:"popularity_score"`
}

type Combo struct {
	ComboID       string  `json:"combo_id"`
	Main          string  `json:"main"`
	Side          string  `json:"side"`
	Drink         string  `json:"drink"`
	CalorieCount  int     `json:"calorie_count"`
	PopularityAvg float64 `json:"popularity_score"`
	Reasoning     string  `json:"reasoning"`
}

type DailyMenu struct {
	Day    string  `json:"day"`
	Combos []Combo `json:"combos"`
}

type MenuPlan struct {
	MenuPlan []DailyMenu `json:"menu_plan"`
}
