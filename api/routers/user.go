package routers

import (
	"elektron-canteen/api/controllers"
	"elektron-canteen/api/mid"
	jwtutil "elektron-canteen/foundation/jwt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
)

type UserRouter struct {
	router     *gin.Engine
	controller controllers.UserController
}

func NewUserRouter(r *gin.Engine, c controllers.UserController) *UserRouter {
	return &UserRouter{
		router:     r,
		controller: c,
	}
}

func (r *UserRouter) Initialize() {
	r.router.Use(mid.Auth())
	r.router.GET("/user", r.getUserData)
}

func (r *UserRouter) getUserData(c *gin.Context) {
	cookie, err := c.Cookie("token")
	if err != nil {
		responseWithError(c, err)
	}

	claims, err := jwtutil.DecodeIntoClaims(cookie)
	if err != nil {
		responseWithError(c, err)
	}

	userID, err := primitive.ObjectIDFromHex(claims["user"].(string))
	if err != nil {
		panic(err)
	}

	user, err := r.controller.Get(userID)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			c.JSON(http.StatusForbidden, gin.H{
				"error": "User doesnt exists",
			})
			return
		}
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}
