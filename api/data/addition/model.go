package addition

import (
	"context"
	"elektron-canteen/foundation/database"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Model interface {
	Create(ctx context.Context, na NewAddition) (primitive.ObjectID, error)
	Update(ctx context.Context, id primitive.ObjectID, na Addition) error
	Delete(ctx context.Context, id primitive.ObjectID) error

	QueryAll(ctx context.Context) ([]Addition, error)
	QueryByName(ctx context.Context, name string) (*Addition, error)
	QueryByID(ctx context.Context, id primitive.ObjectID) (*Addition, error)
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

func (m modelImpl) Create(ctx context.Context, na NewAddition) (primitive.ObjectID, error) {
	coll := m.db.Database("elektron_canteen").Collection("additions")

	res, err := coll.InsertOne(ctx, na)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return res.InsertedID.(primitive.ObjectID), nil
}

func (m modelImpl) Update(ctx context.Context, id primitive.ObjectID, na Addition) error {
	//coll := m.db.Database("elektron_canteen").Collection("additions")

	// err := coll.UpdateOne(ctx, bson.M{"id"})
	panic("FUNCTION NOT IMPLEMENTED")
	return nil
}

func (m modelImpl) Delete(ctx context.Context, id primitive.ObjectID) error {
	coll := m.db.Database("elektron_canteen").Collection("additions")

	_, err := coll.DeleteOne(ctx, bson.M{"_id": id})

	return err
}

func (m modelImpl) QueryAll(ctx context.Context) ([]Addition, error) {
	coll := m.db.Database("elektron_canteen").Collection("additions")

	var additions []Addition
	cur, err := coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cur.All(ctx, &additions); err != nil {
		if err == mongo.ErrNoDocuments {
			return additions, nil
		}
		panic(err)
	}

	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, err
		}
		panic(err)
	}

	return additions, nil
}

func (m modelImpl) QueryByName(ctx context.Context, name string) (*Addition, error) {
	coll := m.db.Database("elektron_canteen").Collection("additions")

	var a *Addition
	err := coll.FindOne(ctx, bson.M{"name": name}).Decode(&a)
	if err != nil {
		if err != mongo.ErrNoDocuments {
			return nil, errors.New("addition not found, passed addition name: " + name)
		}
		panic(err)
	}

	return a, nil
}

func (m modelImpl) QueryByID(ctx context.Context, id primitive.ObjectID) (*Addition, error) {
	coll := m.db.Database("elektron_canteen").Collection("additions")

	var a *Addition
	err := coll.FindOne(ctx, bson.M{"_id": id}).Decode(&a)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("addition not found, passed id: " + id.Hex())
		}
		panic(err)
	}

	return a, nil
}
