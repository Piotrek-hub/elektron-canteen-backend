package menu

import (
	"context"
	"elektron-canteen/foundation/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Model interface {
	Create(ctx context.Context, menu Menu) (primitive.ObjectID, error)
	Update(ctx context.Context, menu Menu) error
	Delete(ctx context.Context, day string) error

	QueryAll(ctx context.Context) ([]Menu, error)
	QueryByDay(ctx context.Context, day string) (*Menu, error)
	QueryRanged(ctx context.Context, days []string) ([]Menu, error)
}

type modelImpl struct {
	db *mongo.Client
}

var instance Model

func Instance() Model {
	db, err := database.GetClient()
	if err != nil {
		panic(err)
	}

	if instance == nil {
		instance = &modelImpl{db: db}
	}

	return instance
}

func (m modelImpl) Create(ctx context.Context, menu Menu) (primitive.ObjectID, error) {
	coll := m.db.Database("elektron_canteen").Collection("menus")

	res, err := coll.InsertOne(ctx, menu)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func (m modelImpl) Update(ctx context.Context, menu Menu) error {
	coll := m.db.Database("elektron_canteen").Collection("menus")

	_, err := coll.UpdateOne(ctx, bson.D{{"day", menu.Day}}, bson.D{{"$set", bson.M{"day": menu.Day, "meals": menu.Meals, "availableMeals": menu.AvailableMeals}}})
	return err
}

func (m modelImpl) Delete(ctx context.Context, day string) error {
	coll := m.db.Database("elektron_canteen").Collection("menus")

	_, err := coll.DeleteOne(ctx, bson.D{{"day", day}})
	return err
}

func (m modelImpl) QueryAll(ctx context.Context) ([]Menu, error) {
	coll := m.db.Database("elektron_canteen").Collection("menus")

	var menus []Menu
	cur, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	if err = cur.All(ctx, &menus); err != nil {
		if err == mongo.ErrNoDocuments {
			return menus, nil
		}
		panic(err)
	}

	return menus, nil
}

func (m modelImpl) QueryByDay(ctx context.Context, day string) (*Menu, error) {
	coll := m.db.Database("elektron_canteen").Collection("menus")

	var menu Menu
	coll.FindOne(ctx, bson.D{{"day", day}}).Decode(&menu)

	return &menu, nil
}

func (m modelImpl) QueryRanged(ctx context.Context, days []string) ([]Menu, error) {
	coll := m.db.Database("elektron_canteen").Collection("menus")

	filter := bson.M{
		"day": bson.M{"$in": days},
	}

	var menus []Menu
	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	if err = cur.All(ctx, &menus); err != nil {
		if err == mongo.ErrNoDocuments {
			return menus, nil
		}
		panic(err)
	}

	return menus, nil
}
