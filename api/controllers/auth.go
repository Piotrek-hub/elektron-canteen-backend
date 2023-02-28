package controllers

import (
	"context"
	"elektron-canteen/api/data/user"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthController struct {
	user      user.Model
	validator user.Validator
}

func NewAuthController() *AuthController {
	return &AuthController{
		user:      user.Instance(),
		validator: *user.NewValidator(),
	}
}

func (c AuthController) Add(nu user.NewUser) error {
	ctx := context.Background()

	if err := c.validator.ValidateUser(nu); err != nil {
		return err
	}

	user, err := c.user.QueryByEmail(ctx, nu.Email)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}

	if user.Surname != "" {
		return errors.New("User already exists")
	}

	if _, err = c.user.Create(ctx, nu); err != nil {
		return err
	}

	return nil
}
