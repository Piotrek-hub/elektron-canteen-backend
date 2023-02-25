package controllers

import (
	"context"
	"elektron-canteen/api/data/user"
)

type AuthController struct {
  user user.Model	
  validator user.Validator 
}

func NewAuthController() *AuthController {
  return &AuthController{
	user: user.Instance(),
	validator: *user.NewValidator(),
  }
}

func (c AuthController) Add(nu user.NewUser) error {
  ctx := context.Background()
  
  if err := c.validator.ValidateUser(nu); err != nil {
	return err
  }

  c.user.Create(ctx, nu)

  return nil
}
