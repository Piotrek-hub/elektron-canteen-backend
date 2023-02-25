package routers

import (
	"elektron-canteen/api/controllers"
	"elektron-canteen/api/data/user"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	router *gin.Engine
	controller controllers.AuthController
}

func NewAuthRouter(r *gin.Engine, c controllers.AuthController) *AuthRouter {
  return &AuthRouter{
	router: r,
	controller: c,
  }
} 

func (r *AuthRouter) Initialize() {
  ar := r.router.Group("/auth")

  ar.POST("/login", r.login)

  ar.POST("/register", r.register)
}

func (r *AuthRouter) login(c *gin.Context) {
	email := c.PostForm("email")
	name := c.PostForm("name")
	surname := c.PostForm("surname")
	password := c.PostForm("password")

	nu := user.NewUser{
	  Email: email,
	  Name: name,
	  Surname: surname,
	  Password: password,
	}
	

	if err := r.controller.Add(nu); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{
		"status": http.StatusBadRequest,
		"error": err.Error(),
	  })
	  return; 
	}

	c.JSON(http.StatusOK, gin.H{
	  "status": http.StatusOK,
	})
	return;
}

func (r *AuthRouter) register(c *gin.Context) {
	email := c.PostForm("email")
	name := c.PostForm("name")
	surname := c.PostForm("surname")
	password := c.PostForm("password")

	nu := user.NewUser{
	  Email: email,
	  Name: name,
	  Surname: surname,
	  Password: password,
	}
	

	if err := r.controller.Add(nu); err != nil {
	  c.JSON(http.StatusBadRequest, gin.H{
		"status": http.StatusBadRequest,
		"error": err.Error(),
	  })
	  return; 
	}

	c.JSON(http.StatusOK, gin.H{
	  "status": http.StatusOK,
	})
	return;
}
