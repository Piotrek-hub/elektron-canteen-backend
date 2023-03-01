package routers

import (
	"elektron-canteen/api/controllers"
	"elektron-canteen/api/data/user"
	"elektron-canteen/foundation/jwt"
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
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
	}
	user, err := r.controller.Login(nu)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	jwtToken, err := jwt.GenerateJWT(user)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successfull",
		"token":   jwtToken,
	})
	return
}

func (r *AuthRouter) register(c *gin.Context) {
	var nu user.NewUser
	c.BindJSON(&nu)

	if err := r.controller.Register(nu); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Register successfull",
	})
	return
}
