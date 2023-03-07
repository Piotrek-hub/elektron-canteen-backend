package meal

import (
	"context"
	"elektron-canteen/foundation/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type Model interface {
	Create(ctx context.Context, nm NewMeal) (primitive.ObjectID, error)
	Update(ctx context.Context, id primitive.ObjectID, nm NewMeal) error
	Delete(ctx context.Context, id primitive.ObjectID) error

	QueryAll(ctx context.Context) ([]Meal, error)
	QueryByID(ctx context.Context, id primitive.ObjectID) (*Meal, error)
	QueryByName(ctx context.Context, name string) (*Meal, error)
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

func (m modelImpl) Create(ctx context.Context, nm NewMeal) (primitive.ObjectID, error) {
	coll := m.db.Database("elektron_canteen").Collection("meals")

	result, err := coll.InsertOne(ctx, nm)
	if err != nil {
		return primitive.ObjectID{}, err
	}

	return result.InsertedID.(primitive.ObjectID), nil
}

func (m modelImpl) Update(ctx context.Context, id primitive.ObjectID, nm NewMeal) error {
	coll := m.db.Database("elektron_canteen").Collection("meals")
	log.Println(id)
	_, err := coll.UpdateOne(ctx, bson.D{{"_id", id}}, bson.D{{"$set", bson.M{"name": nm.Name, "price": nm.Price, "additions": nm.Additions}}})
	return err
}

func (m modelImpl) Delete(ctx context.Context, id primitive.ObjectID) error {
	coll := m.db.Database("elektron_canteen").Collection("meals")

	_, err := coll.DeleteOne(ctx, bson.D{{"_id", id}})
	return err
}

func (m modelImpl) QueryAll(ctx context.Context) ([]Meal, error) {
	coll := m.db.Database("elektron_canteen").Collection("meals")
	cur, err := coll.Find(ctx, bson.D{})

	var meals []Meal
	if err = cur.All(ctx, &meals); err != nil {
		if err == mongo.ErrNoDocuments {
			return meals, nil
		}
		panic(err)
	}

	return meals, nil
}

func (m modelImpl) QueryByID(ctx context.Context, id primitive.ObjectID) (*Meal, error) {
	coll := m.db.Database("elektron_canteen").Collection("meals")

	var meal Meal
	err := coll.FindOne(ctx, bson.D{{"_id", id}}).Decode(&meal)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		panic(err)
	}

	return &meal, nil
}

func (m modelImpl) QueryByName(ctx context.Context, name string) (*Meal, error) {
	coll := m.db.Database("elektron_canteen").Collection("meals")

	var meal Meal
	err := coll.FindOne(ctx, bson.D{{"name", name}}).Decode(&meal)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		panic(err)
	}

	return &meal, nil
}
