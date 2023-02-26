package menu

import (
  "elektron-canteen/api/data/dish"
  "time"
)

type Menu struct {
  Day time.Time
  Dishes []dish.Dish 
  AvailableDishes []dish.Dish
}


