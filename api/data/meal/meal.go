package meal

import (
	"context"
	"elektron-canteen/api/data/addition"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Meal struct {
	ID        primitive.ObjectID `bson:"_id" json:"id"`
	Name      string             `bson:"name" json:"name"`
	Price     float32            `bson:"price" json:"price"`
	Additions []string           `bson:"additions" json:"additions"`
}

type NewMeal struct {
	Name      string   `bson:"name" json:"name"`
	Price     float32  `bson:"price" json:"price"`
	Additions []string `bson:"additions" json:"additions"`
}

type Response struct {
	ID        primitive.ObjectID  `bson:"_id" json:"id"`
	Name      string              `bson:"name" json:"name"`
	Price     float32             `bson:"price" json:"price"`
	Additions []addition.Addition `bson:"additions" json:"additions"`
}

func (m *Meal) ToResponse(ctx context.Context, model addition.Model) (*Response, error) {
	var mr Response
	mr.ID = m.ID
	mr.Name = m.Name
	mr.Price = m.Price

	for _, addition := range m.Additions {
		id, err := primitive.ObjectIDFromHex(addition)
		if err != nil {
			panic(err)
		}

		aTmp, err := model.QueryByID(ctx, id)
		if err != nil {
			return nil, err
		}

		mr.Additions = append(mr.Additions, *aTmp)
	}

	return &mr, nil
}
