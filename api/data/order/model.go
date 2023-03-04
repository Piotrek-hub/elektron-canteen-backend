package order

import (
	"context"
	"elektron-canteen/foundation/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Model interface {
	Create(ctx context.Context, no NewOrder) (primitive.ObjectID, error)
	//Update(ctx context.Context, order Order, newMeal primitive.ObjectID) error
	//Delete(ctx context.Context, id primitive.ObjectID) error

	UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error
	QueryAll(ctx context.Context) ([]Order, error)
	QueryByNotStatus(ctx context.Context, status string) ([]Order, error)
	QueryByUser(ctx context.Context, userID primitive.ObjectID) ([]Order, error)
	QueryByID(ctx context.Context, orderID primitive.ObjectID) ([]Order, error)
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

func (m modelImpl) Create(ctx context.Context, no NewOrder) (primitive.ObjectID, error) {
	coll := m.db.Database("elektron_canteen").Collection("orders")

	res, err := coll.InsertOne(ctx, no)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func (m modelImpl) UpdateStatus(ctx context.Context, id primitive.ObjectID, status string) error {
	coll := m.db.Database("elektron_canteen").Collection("orders")

	_, err := coll.UpdateOne(ctx, bson.D{{"_id", id}}, bson.D{{"$set", bson.D{{"status", status}}}})
	return err
}

func (m modelImpl) QueryAll(ctx context.Context) ([]Order, error) {
	coll := m.db.Database("elektron_canteen").Collection("orders")

	cur, err := coll.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}

	var orders []Order
	if err = cur.All(ctx, &orders); err != nil {
		if err == mongo.ErrNoDocuments {
			return orders, nil
		}
		panic(err)
	}

	return orders, nil
}

func (m modelImpl) QueryByNotStatus(ctx context.Context, status string) ([]Order, error) {
	coll := m.db.Database("elektron_canteen").Collection("orders")

	filter := bson.D{{"status", bson.D{{"$not", bson.D{{"$eq", status}}}}}}

	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var orders []Order
	if err = cur.All(ctx, &orders); err != nil {
		if err == mongo.ErrNoDocuments {
			return orders, nil
		}
		panic(err)
	}

	return orders, nil
}

func (m modelImpl) QueryByUser(ctx context.Context, userID primitive.ObjectID) ([]Order, error) {
	coll := m.db.Database("elektron_canteen").Collection("orders")

	filter := bson.D{{"user", userID}}

	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var orders []Order
	if err = cur.All(ctx, &orders); err != nil {
		if err == mongo.ErrNoDocuments {
			return orders, nil
		}
		panic(err)
	}

	return orders, nil
}

func (m modelImpl) QueryByID(ctx context.Context, orderID primitive.ObjectID) ([]Order, error) {
	coll := m.db.Database("elektron_canteen").Collection("orders")

	filter := bson.D{{"_id", orderID}}

	cur, err := coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	var orders []Order
	if err = cur.All(ctx, &orders); err != nil {
		if err == mongo.ErrNoDocuments {
			return orders, nil
		}
		panic(err)
	}

	return orders, nil
}
