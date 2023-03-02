package user

import (
	"context"
	"elektron-canteen/api/config"
	"elektron-canteen/foundation/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Model interface {
	Create(ctx context.Context, nu NewUser) (primitive.ObjectID, error)
	// Update(ctx context.Context, id primitive.ObjectID, uu UpdateUser) (*User, error)
	//Delete(ctx context.Context, id primitive.ObjectID) error

	QueryAll(ctx context.Context) ([]User, error)
	QueryByEmail(ctx context.Context, email string) (*User, error)
	QueryByID(ctx context.Context, id primitive.ObjectID) (*User, error)
}

type modelImpl struct {
	db  *mongo.Client
	cfg config.Config
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

func (m modelImpl) Create(ctx context.Context, nu NewUser) (primitive.ObjectID, error) {
	coll := m.db.Database("elektron_canteen").Collection("users")

	result, err := coll.InsertOne(ctx, nu)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (m modelImpl) QueryAll(ctx context.Context) ([]User, error) {
	coll := m.db.Database("elektron_canteen").Collection("users")
	cur, err := coll.Find(ctx, bson.D{})

	var users []User
	if err = cur.All(ctx, &users); err != nil {
		panic(err)
	}

	return users, nil
}

func (m modelImpl) QueryByEmail(ctx context.Context, email string) (*User, error) {
	coll := m.db.Database("elektron_canteen").Collection("users")

	var user User
	err := coll.FindOne(ctx, bson.D{{"email", email}}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		panic(err)
	}

	return &user, nil
}

func (m modelImpl) QueryByID(ctx context.Context, id primitive.ObjectID) (*User, error) {
	coll := m.db.Database("elektron_canteen").Collection("users")

	var user User
	err := coll.FindOne(ctx, bson.D{{"_id", id}}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, err
		}
		panic(err)
	}

	return &user, nil
}
