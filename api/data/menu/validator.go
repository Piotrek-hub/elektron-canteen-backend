package menu

import (
	"errors"
	"regexp"
)

type Validator struct {
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v Validator) ValidateMenus(menus []Menu) error {

	for _, menu := range menus {
		if err := v.ValidateMenu(menu); err != nil {
			return err
		}
	}

	for i := 0; i < len(menus); i++ {
		for j := 0; j < len(menus); j++ {
			if i != j {
				if menus[i].Day == menus[j].Day {
					return errors.New("menu copied, menu day: " + menus[i].Day)
				}
			}
		}
	}

	return nil
}

func (v Validator) ValidateMenu(menu Menu) error {
	if menu.Day == "" {
		return errors.New("menu day is missing")
	}
	if len(menu.Meals) == 0 {
		return errors.New("meals are missing")
	}

	return nil
}

func (v Validator) ValidateDay(day string) error {
	re := regexp.MustCompile("(?i)\\d\\d\\d\\d-\\d\\d-\\d\\d")
	if !re.MatchString(day) {
		return errors.New("wrong day format")
	}

	return nil
}
