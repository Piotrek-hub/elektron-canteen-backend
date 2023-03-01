package user

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	ID       primitive.ObjectID `bson:"_id",json:"id"`
	Email    string             `json:"email""`
	Name     string             `json:"name""`
	Surname  string             `json:"surname""`
	Password string             `json:"password""`
}

type NewUser struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Password string `json:"password"`
}

type UpdateUser struct{}
