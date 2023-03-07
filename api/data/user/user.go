package user

import "go.mongodb.org/mongo-driver/bson/primitive"

const (
	NORMAL_ROLE string = "normal_role"
	ADMIN_ROLE  string = "admin_role"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id" json:"id"`
	Email    string             `json:"email"`
	Role     string             `json:"role"`
	Name     string             `json:"name"`
	Surname  string             `json:"surname"`
	Points   float32            `json:"points"`
	Password string             `json:"password"`
}

type NewUser struct {
	Email    string `json:"email"`
	Name     string `json:"name"`
	Role     string `json:"role"`
	Surname  string `json:"surname"`
	Points   int    `json:"points"`
	Password string `json:"password"`
}

type UpdateUser struct{}
