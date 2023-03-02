package controllers

import (
	"context"
	"elektron-canteen/api/data/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController struct {
	user user.Model
}

func NewUserController() *UserController {
	return &UserController{
		user: user.Instance(),
	}
}

func (c UserController) Get(userID primitive.ObjectID) (*user.User, error) {
	return c.user.QueryByID(context.TODO(), userID)
}
