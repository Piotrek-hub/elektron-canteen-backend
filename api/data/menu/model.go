package menu

import (
  "log"
  "context")

type Model interface {
  Create(ctx context.Context, menu Menu) (*Menu, error)

//  QueryAll(ctx context.Context) ([]Menu, error)
//  QueryByDay(ctx context.Context, id primitive.ObjectID) (User, error)
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


func (m modelImpl) Create(ctx context.Context, menu Menu) (*Menu, error) {
  log.Println(menu)
  return nil, nil
}



