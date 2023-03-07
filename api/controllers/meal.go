package controllers

import (
	"context"
	"elektron-canteen/api/data/addition"
	"elektron-canteen/api/data/meal"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MealController struct {
	meal      meal.Model
	addition  addition.Model
	validator meal.Validator
}

func NewMealController() *MealController {
	return &MealController{
		meal:      meal.Instance(),
		addition:  addition.Instance(),
		validator: *meal.NewValidator(),
	}
}

func (c MealController) Add(nm meal.NewMeal) error {
	ctx := context.Background()
	err := c.validator.ValidateMeal(nm)
	if err != nil {
		return err
	}

	m, err := c.meal.QueryByName(ctx, nm.Name)
	if err != nil && err != mongo.ErrNoDocuments {
		return err
	}

	if m != nil {
		return errors.New("Meal already exists")
	}

	_, err = c.meal.Create(ctx, nm)

	return err
}

func (c MealController) Update(id primitive.ObjectID, nm meal.NewMeal) error {
	ctx := context.Background()

	if err := c.validator.ValidateMeal(nm); err != nil {
		return err
	}

	if err := c.meal.Update(ctx, id, nm); err != nil {
		return err
	}

	return nil
}

func (c MealController) Delete(id primitive.ObjectID) error {
	ctx := context.Background()

	return c.meal.Delete(ctx, id)
}

func (c MealController) GetAll() ([]meal.Response, error) {
	ctx := context.Background()
	meals, err := c.meal.QueryAll(ctx)
	if err != nil {
		return nil, err
	}

	var mr []meal.Response
	for _, meal := range meals {
		r, err := meal.ToResponse(ctx, c.addition)
		if err != nil {
			return nil, err
		}

		mr = append(mr, *r)
	}

	return mr, nil
}

func (c MealController) GetByID(id primitive.ObjectID) (*meal.Response, error) {
	ctx := context.Background()

	meal, err := c.meal.QueryByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return meal.ToResponse(ctx, c.addition)
}
