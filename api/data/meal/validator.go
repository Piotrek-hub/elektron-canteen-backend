package meal

import (
	"errors"
)

type Validator struct {
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v Validator) ValidateMeal(meal NewMeal) error {
	if meal.Name == "" {
		return errors.New("meal name is missing")
	}
	if meal.Price == 0 {
		return errors.New("price is 0")
	}

	return nil
}
