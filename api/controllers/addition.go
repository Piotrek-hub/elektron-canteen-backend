package controllers

import (
	"context"
	"elektron-canteen/api/data/addition"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type AdditionController struct {
	addition  addition.Model
	validator addition.Validator
}

func NewAdditionController() *AdditionController {
	return &AdditionController{
		addition:  addition.Instance(),
		validator: *addition.NewValidator()}
}

func (c AdditionController) GetAll() ([]addition.Addition, error) {
	ctx := context.Background()

	return c.addition.QueryAll(ctx)
}

func (c AdditionController) GetById(id primitive.ObjectID) (*addition.Addition, error) {
	ctx := context.Background()

	return c.addition.QueryByID(ctx, id)
}

func (c AdditionController) GetByName(name string) (*addition.Addition, error) {
	ctx := context.Background()

	return c.addition.QueryByName(ctx, name)
}

func (c AdditionController) Create(na addition.NewAddition) (primitive.ObjectID, error) {
	ctx := context.Background()

	err := c.validator.ValidateAddition(na)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return c.addition.Create(ctx, na)
}

func (c AdditionController) Delete(id primitive.ObjectID) error {
	ctx := context.Background()

	return c.addition.Delete(ctx, id)
}
