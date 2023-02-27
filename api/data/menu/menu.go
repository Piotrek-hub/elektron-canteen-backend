package menu

import (
  "elektron-canteen/api/data/dish"
)

type Menu struct {
  Day string `json:"day"`
  Dishes []dish.Dish `json:"dishes"`
  AvailableDishes []dish.Dish `json:"availableDishes"`
}


