package user

import (
	"context"
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
	QueryByEmail(ctx context.Context, email string) (User, error)
}

type modelImpl struct {
	db *mongo.Client
}

var instance Model

func Instance() Model {
	if instance == nil {
		db, err := database.GetClient("users")
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
	coll := m.db.Database("users").Collection("users")

	result, err := coll.InsertOne(context.TODO(), nu)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return result.InsertedID.(primitive.ObjectID), nil
}

func (m modelImpl) QueryAll(ctx context.Context) ([]User, error) {
	coll := m.db.Database("users").Collection("users")
	cur, err := coll.Find(ctx, bson.D{})

	var users []User
	if err = cur.All(ctx, &users); err != nil {
		panic(err)
	}

	return users, nil
}

func (m modelImpl) QueryByEmail(ctx context.Context, email string) (User, error) {
	coll := m.db.Database("users").Collection("users")

	var user User
	err := coll.FindOne(ctx, bson.D{{"email", email}}).Decode(&user)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return User{}, err
		}
		panic(err)
	}

	return user, nil
}
