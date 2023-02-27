package controllers

import (
	//	"context"
	"context"
	"elektron-canteen/api/data/menu"
	"time"
)

type MenuController struct {
	menu menu.Model
	validator menu.Validator
}

func NewMenuController() *MenuController {
  return &MenuController{
	menu: menu.Instance(),
  }
}

func (c MenuController) Add(menu menu.Menu) error {
//  ctx := context.Background()
  
  c.validator.ValidateMenu(menu)

  return nil
}

func (c MenuController) GetByDay(day time.Time) (*menu.Menu, error) {
  ctx := context.Background()

  m, err := c.menu.QueryByDay(ctx, day)
  if err != nil {
	return nil, err
  }

  return m, nil
}
