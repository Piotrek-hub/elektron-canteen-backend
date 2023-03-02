package controllers

import (
	"context"
	"elektron-canteen/api/data/meal"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type MealController struct {
	meal      meal.Model
	validator meal.Validator
}

func NewMealController() *MealController {
	return &MealController{
		meal:      meal.Instance(),
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

func (c MealController) GetAll() ([]meal.Meal, error) {
	ctx := context.Background()

	return c.meal.QueryAll(ctx)
}

func (c MealController) GetByID(id primitive.ObjectID) (*meal.Meal, error) {
	ctx := context.Background()

	return c.meal.QueryByID(ctx, id)
}
