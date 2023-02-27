package menu

import (
	"context"
	"time"
)

type Model interface {
  Create(ctx context.Context, menu Menu) (*Menu, error)

  QueryAll(ctx context.Context) ([]Menu, error)
  QueryByDay(ctx context.Context, day time.Time) (*Menu, error)
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
  return nil, nil
}

func (m modelImpl) QueryAll(ctx context.Context) ([]Menu, error) {
  return nil, nil 
}

func (m modelImpl) QueryByDay(ctx context.Context, day time.Time) (*Menu, error) {
  return nil, nil
}


