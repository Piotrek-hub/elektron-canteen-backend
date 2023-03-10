package order

import (
	"context"
	"elektron-canteen/api/data/addition"
	"elektron-canteen/api/data/meal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Order status:
const (
	ACCEPTED = "accepted"
	DECLINED = "declined"
	WAITING  = "waiting"
	DONE     = "done"
	CANCELED = "cancelled"
)

// Payment status
const (
	ONLINE_PAYMENT     = "online_payment"
	STATIONARY_PAYMENT = "stationary_payment"
)

type Order struct {
	ID            primitive.ObjectID `bson:"_id" json:"id"`
	Meal          primitive.ObjectID `json:"meal"`
	User          primitive.ObjectID `json:"user"`
	Additions     []string           `json:"additions"`
	Status        string             `json:"status"`
	PaymentMethod string             `json:"paymentMethod"`
	DueTime       string             `json:"dueTime"`
	Date          string             `json:"date"`
	Price         float32            `json:"price"`
	Number        int                `json:"number"`
	Drink         bool               `json:"drink"`
}

type NewOrder struct {
	Meal          primitive.ObjectID `json:"meal"`
	User          primitive.ObjectID `json:"user"`
	Additions     []string           `json:"additions"`
	Status        string             `json:"status"`
	PaymentMethod string             `json:"paymentMethod"`
	DueTime       string             `json:"dueTime"`
	Date          string             `json:"date"`
	Price         float32            `json:"price"`
	Number        int                `json:"number"`
	Drink         bool               `json:"drink"`
}

type Response struct {
	Meal          meal.Meal           `json:"meal"`
	User          primitive.ObjectID  `json:"user"`
	Additions     []addition.Addition `json:"additions"`
	Status        string              `json:"status"`
	PaymentMethod string              `json:"paymentMethod"`
	DueTime       string              `json:"dueTime"`
	Date          string              `json:"date"`
	Price         float32             `json:"price"`
	Number        int                 `json:"number"`
	Drink         bool                `json:"drink"`
}

func (o *Order) ToResponse(ctx context.Context, mm meal.Model, am addition.Model) (*Response, error) {
	res := Response{}
	res.Status = o.Status
	res.User = o.User
	res.PaymentMethod = o.PaymentMethod
	res.DueTime = o.DueTime
	res.Date = o.Date
	res.Price = o.Price
	res.Number = o.Number
	res.Drink = o.Drink

	meal, err := mm.QueryByID(ctx, o.Meal)
	if err != nil {
		return nil, err
	}
	res.Meal = *meal

	var additions []addition.Addition

	for _, a := range o.Additions {
		aID, err := primitive.ObjectIDFromHex(a)
		if err != nil {
			return nil, err
		}

		addition, err := am.QueryByID(ctx, aID)
		if err != nil {
			return nil, err
		}

		additions = append(additions, *addition)

	}

	res.Additions = additions

	return &res, nil
}
