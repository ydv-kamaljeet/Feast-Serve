package menu

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func CategorizeMenu(items []MenuItem) map[string][]MenuItem {
	categorized := make(map[string][]MenuItem)
	for _, item := range items {
		categorized[item.Category] = append(categorized[item.Category], item)
	}
	return categorized
}

func CalculateComboMetrics(main, side, drink MenuItem) (int, float64) {
	totalCalories := main.Calories + side.Calories + drink.Calories
	avgPopularity := (main.PopularityScore + side.PopularityScore + drink.PopularityScore) / 3.0
	return totalCalories, avgPopularity
}

func IsValidCombo(main, side, drink MenuItem, minCalories, maxCalories int, popularityTolerance float64) bool {
	totalCalories, _ := CalculateComboMetrics(main, side, drink)

	if totalCalories < minCalories || totalCalories > maxCalories {
		return false
	}

	popularityScores := []float64{main.PopularityScore, side.PopularityScore, drink.PopularityScore}
	sort.Float64s(popularityScores)

	return (popularityScores[len(popularityScores)-1] - popularityScores[0]) <= popularityTolerance
}

func GenerateReasoning(main, side, drink MenuItem, totalCalories int, avgPopularity float64) string {
	tasteProfiles := map[string]bool{
		main.TasteProfile:  true,
		side.TasteProfile:  true,
		drink.TasteProfile: true,
	}

	tasteDesc := "a mixed taste profile"
	if len(tasteProfiles) == 1 {
		for k := range tasteProfiles {
			tasteDesc = fmt.Sprintf("a %s profile", k)
		}
	}

	return fmt.Sprintf(
		"This combo features %s, consists of popular choices (average popularity: %.2f), and meets the calorie target (%d kcal).",
		tasteDesc, math.Round(avgPopularity*100)/100, totalCalories,
	)
}

func Signature(main, side, drink MenuItem) string {
	names := []string{main.ItemName, side.ItemName, drink.ItemName}
	sort.Strings(names)
	return strings.Join(names, "_")
}
