package controllers

import (
//	"context"
	"elektron-canteen/api/data/menu"
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
