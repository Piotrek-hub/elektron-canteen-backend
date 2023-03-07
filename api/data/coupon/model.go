package coupon

import (
	"context"
	"elektron-canteen/foundation/database"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Model interface {
	Create(ctx context.Context, c Coupon) (primitive.ObjectID, error)
	Delete(ctx context.Context, code string) error

	QueryByCode(ctx context.Context, code string) (*Coupon, error)
}

type modelImpl struct {
	db *mongo.Client
}

var instance Model

func Instance() Model {
	if instance == nil {
		db, err := database.GetClient()
		if err != nil {
			panic(err)
		}
		instance = &modelImpl{
			db: db,
		}
	}

	return instance
}

func (m modelImpl) Create(ctx context.Context, c Coupon) (primitive.ObjectID, error) {
	var id primitive.ObjectID
	coll := m.db.Database("elektron_canteen").Collection("coupons")

	res, err := coll.InsertOne(ctx, c)
	if err != nil {
		return id, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func (m modelImpl) Delete(ctx context.Context, code string) error {
	coll := m.db.Database("elektron_canteen").Collection("coupons")

	_, err := coll.DeleteOne(ctx, bson.M{"code": code})

	return err
}

func (m modelImpl) QueryByCode(ctx context.Context, code string) (*Coupon, error) {
	coll := m.db.Database("elektron_canteen").Collection("coupons")

	var c Coupon
	err := coll.FindOne(ctx, bson.M{"code": code}).Decode(&c)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &c, errors.New("coupon doesnt exist")
		}
		return nil, err
	}

	return &c, nil
}
