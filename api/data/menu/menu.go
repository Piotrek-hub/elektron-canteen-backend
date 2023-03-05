package menu

import (
	"context"
	"elektron-canteen/api/data/meal"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type Menu struct {
	Day            string   `json:"day"`
	Meals          []string `json:"meals"`
	AvailableMeals []string `json:"availableMeals" bson:"availableMeals"`
}

type Response struct {
	Day            string      `json:"day"`
	Meals          []meal.Meal `json:"meals"`
	AvailableMeals []meal.Meal `json:"availableMeals" bson:"availableMeals"`
}

func (m *Menu) ToResponse(ctx context.Context, model meal.Model) (*Response, error) {
	var mr Response
	mr.Day = m.Day
	for i := 0; i < len(m.Meals); i++ {
		mealID, err := primitive.ObjectIDFromHex(m.Meals[i])
		if err != nil {
			panic(err)
		}

		mTmp, err := model.QueryByID(ctx, mealID)
		if err != nil {
			return nil, err
		}

		mr.Meals = append(mr.Meals, *mTmp)
	}

	log.Println(m)
	for i := 0; i < len(m.AvailableMeals); i++ {
		if m.AvailableMeals[i] == mr.Meals[i].ID.Hex() {
			mr.AvailableMeals = append(mr.AvailableMeals, mr.Meals[i])
		}
	}

	return &mr, nil
}
