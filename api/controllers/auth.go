package controllers

import (
	"context"
	"elektron-canteen/api/data/user"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
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

func (c AuthController) Register(nu user.NewUser) error {
	ctx := context.Background()

	if err := c.validator.ValidateUser(nu); err != nil {
		return err
	}

	u, err := c.user.QueryByEmail(ctx, nu.Email)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}

	if u != nil {
		return errors.New("user already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(nu.Password), 10)
	if err != nil {
		return err
	}

	nu.Password = string(hashedPassword[:])
	nu.Role = user.NORMAL_ROLE
	nu.Points = 0

	if _, err = c.user.Create(ctx, nu); err != nil {
		return err
	}

	return err
}

func (c AuthController) Login(nu user.NewUser) (*user.User, error) {
	ctx := context.Background()

	if err := c.validator.ValidateUser(nu); err != nil {
		return nil, err
	}

	u, err := c.user.QueryByEmail(ctx, nu.Email)
	if err == mongo.ErrNoDocuments {
		return nil, errors.New("User doesn't exists")
	}

	if err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(nu.Password)); err != nil {
		return nil, err
	}

	return u, err
}
