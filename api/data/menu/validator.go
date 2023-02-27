package menu

import (
	"errors"
)

type Validator struct {

}

func NewValidator() *Validator {
  return &Validator{}
}

func (v Validator) ValidateMenu(menu Menu) error {
  if menu.Day == "" {
	return errors.New("menu day is missing")
  }
  if len(menu.Dishes) == 0 {
	return errors.New("dishes are missing")
  }

  return nil
} 
