package utils

import "go.mongodb.org/mongo-driver/bson/primitive"

type ArrayType interface {
	int64 | string | float64 | primitive.ObjectID
}

func ArrayContains[T ArrayType](arr []T, elem T) bool {
	for _, v := range arr {
		if v == elem {
			return true
		}
	}
	return false
}
