package menu

import (
	"elektron-canteen/api/data/meal"
)

type Menu struct {
	Day            string      `json:"day"`
	Meals          []meal.Meal `json:"meals"`
	AvailableMeals []meal.Meal `json:"availableMeals"`
}
