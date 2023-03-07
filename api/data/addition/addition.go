package addition

import "go.mongodb.org/mongo-driver/bson/primitive"

type Addition struct {
	ID    primitive.ObjectID `bson:"_id" json:"id"`
	Name  string             `bson:"name" json:"name"`
	Price float32            `bson:"price" json:"price"`
}
