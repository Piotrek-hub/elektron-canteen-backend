package controllers

import (
	//	"context"
	"context"
	"elektron-canteen/api/data/meal"
	"elektron-canteen/api/data/menu"
	"errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type MenuController struct {
	menu      menu.Model
	meal      meal.Model
	validator menu.Validator
}

func NewMenuController() *MenuController {
	return &MenuController{
		menu:      menu.Instance(),
		meal:      meal.Instance(),
		validator: *menu.NewValidator(),
	}
}

func (c MenuController) Add(mr menu.AddRequest) error {
	ctx := context.Background()

	menus := mr.Menus
	if err := c.validator.ValidateMenus(menus); err != nil {
		return err
	}

	for _, menu := range menus {
		if _, err := c.menu.Create(ctx, menu); err != nil {
			if mongo.IsDuplicateKeyError(err) {
				return errors.New("duplicated menu, day: " + menu.Day)
			}
			return err
		}
	}

	return nil
}

func (c MenuController) Update(menu menu.Menu) error {
	ctx := context.Background()

	if err := c.validator.ValidateMenu(menu); err != nil {
		return err
	}

	if err := c.menu.Update(ctx, menu); err != nil {
		return err
	}

	return nil
}

func (c MenuController) Delete(day string) error {
	return c.menu.Delete(context.TODO(), day)
}

func (c MenuController) GetByDay(day string) (*menu.Response, error) {
	ctx := context.Background()

	if err := c.validator.ValidateDay(day); err != nil {
		return nil, err
	}

	menu, err := c.menu.QueryByDay(ctx, day)
	if err != nil {
		return nil, err
	}

	mr, err := menu.ToResponse(ctx, c.meal)
	if err != nil {
		return nil, err
	}

	return mr, nil
}
