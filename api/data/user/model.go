package user

import (
	"context"
	"log"
)

type Model interface {
  Create(ctx context.Context, nu NewUser) (*User, error)
 // Update(ctx context.Context, id primitive.ObjectID, uu UpdateUser) (*User, error)
  //Delete(ctx context.Context, id primitive.ObjectID) error

  QueryAll(ctx context.Context) ([]User, error)
  // QueryByID(ctx context.Context, id primitive.ObjectID) (User, error)
}

type modelImpl struct {
  // db *mongo.Client
}

var instance Model 

func Instance() Model {
  if instance == nil {
	instance = &modelImpl{

	}
  }

  return instance
}


func (m modelImpl) Create(ctx context.Context, nu NewUser) (*User, error) {
  log.Println(nu)
  return nil, nil
}

func (m modelImpl) QueryAll(ctx context.Context) ([]User, error) {
  return nil, nil 
}





