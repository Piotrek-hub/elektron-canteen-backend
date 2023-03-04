package routers

import (
	"elektron-canteen/api/controllers"
	"elektron-canteen/api/data/user"
	jwtutil "elektron-canteen/foundation/jwt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	router     *gin.Engine
	controller controllers.AuthController
}

func NewAuthRouter(r *gin.Engine, c controllers.AuthController) *AuthRouter {
	return &AuthRouter{
		router:     r,
		controller: c,
	}
}

func (r *AuthRouter) Initialize() {
	ar := r.router.Group("/auth")

	ar.POST("/login", r.login)
	ar.POST("/register", r.register)
}

func (r *AuthRouter) login(c *gin.Context) {
	var nu user.NewUser
	if err := c.BindJSON(&nu); err != nil {
		responseWithError(c, err)
	}

	u, err := r.controller.Login(nu)
	if err != nil {
		responseWithError(c, err)
		return
	}

	jwtToken, err := jwtutil.Generate(*u)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successfull",
		"token":   jwtToken,
	})
	return
}

func (r *AuthRouter) register(c *gin.Context) {
	var nu user.NewUser
	err := c.BindJSON(&nu)
	if err != nil {
		responseWithError(c, err)
	}

	if err := r.controller.Register(nu); err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Register successfull",
	})
	return
}
