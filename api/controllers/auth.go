package controllers

import (
	"context"
	"elektron-canteen/api/config"
	"elektron-canteen/api/data/user"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
)

type AuthController struct {
	cfg       config.Config
	user      user.Model
	validator user.Validator
}

func NewAuthController() *AuthController {
	return &AuthController{
		user:      user.Instance(),
		validator: *user.NewValidator(),
		cfg:       config.Load(),
	}
}

func (c AuthController) Register(nu user.NewUser) error {
	ctx := context.Background()

	if err := c.validator.ValidateUser(nu); err != nil {
		return err
	}

	u, err := c.user.QueryByEmail(ctx, nu.Email)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}

	if u.Surname != "" {
		return errors.New("User already exists")
	}

	nu.Role = user.NORMAL_ROLE

	if _, err = c.user.Create(ctx, nu); err != nil {
		return err
	}

	return nil
}

func (c AuthController) Login(nu user.NewUser) (user.User, error) {
	ctx := context.Background()

	if err := c.validator.ValidateUser(nu); err != nil {
		return user.User{}, err
	}

	user, err := c.user.QueryByEmail(ctx, nu.Email)
	if err == mongo.ErrNoDocuments {
		return user, errors.New("User doesn't exists")
	}

	if nu.Password != user.Password {
		return user, errors.New("Wrong password")
	}

	return user, err
}
