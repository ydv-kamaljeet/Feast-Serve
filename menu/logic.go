package menu

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func GenerateMenuSuggestions(masterMenu []MenuItem, numDays, combosPerDay, minCalories, maxCalories int) MenuPlan {
	rand.Seed(time.Now().UnixNano())

	categorizedMenu := CategorizeMenu(masterMenu)
	menuPlan := MenuPlan{MenuPlan: []DailyMenu{}}

	day1UsedItems := map[string]bool{}
	comboUsage := map[string]int{}
	globalCounter := 0

	dayNames := []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

	for dayIndex := 0; dayIndex < numDays; dayIndex++ {
		dailyCombos := GenerateDailyCombos(
			categorizedMenu, combosPerDay, minCalories, maxCalories,
			dayIndex, &globalCounter, &day1UsedItems, comboUsage,
		)

		menuPlan.MenuPlan = append(menuPlan.MenuPlan, DailyMenu{
			Day:    dayNames[dayIndex],
			Combos: dailyCombos,
		})
	}

	return menuPlan
}

func GenerateDailyCombos(
	categorizedMenu map[string][]MenuItem,
	combosPerDay, minCalories, maxCalories int,
	dayIndex int,
	globalCounter *int,
	day1UsedItems *map[string]bool,
	comboUsage map[string]int,
) []Combo {

	mains := categorizedMenu["main"]
	sides := categorizedMenu["side"]
	drinks := categorizedMenu["drink"]

	dayUsed := map[string]bool{}
	daily := []Combo{}

	for len(daily) < combosPerDay {
		main := mains[rand.Intn(len(mains))]
		side := sides[rand.Intn(len(sides))]
		drink := drinks[rand.Intn(len(drinks))]

		sig := Signature(main, side, drink)

		// 3-day rule
		if lastDay, ok := comboUsage[sig]; ok && (dayIndex-lastDay < 3) {
			continue
		}

		// Day-1 unique item rule
		if dayIndex == 0 {
			if (*day1UsedItems)[main.ItemName] || (*day1UsedItems)[side.ItemName] || (*day1UsedItems)[drink.ItemName] {
				continue
			}
		}

		// Current day unique rule
		if dayUsed[main.ItemName] || dayUsed[side.ItemName] || dayUsed[drink.ItemName] {
			continue
		}

		if !IsValidCombo(main, side, drink, minCalories, maxCalories, 0.15) {
			continue
		}

		total, avg := CalculateComboMetrics(main, side, drink)
		*globalCounter++

		combo := Combo{
			ComboID:       fmt.Sprintf("combo_%d", *globalCounter),
			Main:          main.ItemName,
			Side:          side.ItemName,
			Drink:         drink.ItemName,
			CalorieCount:  total,
			PopularityAvg: math.Round(avg*100) / 100,
			Reasoning:     GenerateReasoning(main, side, drink, total, avg),
		}

		daily = append(daily, combo)
		dayUsed[main.ItemName] = true
		dayUsed[side.ItemName] = true
		dayUsed[drink.ItemName] = true
		comboUsage[sig] = dayIndex

		if dayIndex == 0 {
			(*day1UsedItems)[main.ItemName] = true
			(*day1UsedItems)[side.ItemName] = true
			(*day1UsedItems)[drink.ItemName] = true
		}
	}

	return daily
}
