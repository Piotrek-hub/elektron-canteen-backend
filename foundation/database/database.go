package database

import (
	"context"
	"elektron-canteen/api/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func GetClient(dbName string, cfg config.Config) (*mongo.Client, error) {

	uri := "mongodb://" + cfg.Get("mongo_username") + ":" + cfg.Get("mongo_password") + "@" + cfg.Get("mongo_server") + "/" + dbName

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, err
	}

	return client, nil
}
