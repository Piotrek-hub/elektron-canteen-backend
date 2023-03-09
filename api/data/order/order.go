package order

import "go.mongodb.org/mongo-driver/bson/primitive"

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
}
