package meal

import "go.mongodb.org/mongo-driver/bson/primitive"

type Meal struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Price     float32            `bson:"price" json:"price"`
	Additions []string           `bson:"additions" json:"additions"`
	Salads    []string           `bson:"salads" json:"salads"`
}

type NewMeal struct {
	Name      string   `json:"name"`
	Price     float32  `json:"price"`
	Additions []string `json:"additions"`
	Salads    []string `json:"salads"`
}
