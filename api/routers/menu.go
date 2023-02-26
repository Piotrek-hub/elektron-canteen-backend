package routers

import (
	"elektron-canteen/api/controllers"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MenuRouter struct {
	router *gin.Engine
	controller controllers.MenuController
}

func NewMenuRouter(r *gin.Engine, c controllers.MenuController) *MenuRouter{
  return &MenuRouter{
	router: r,
	controller: c,
  }
}

func (r *MenuRouter) Initialize() {
  r.router.GET("/menu", r.getMenu)
  r.router.POST("/menu", r.addMenu)
  r.router.PATCH("/menu", r.updateMenu)
}

func (r *MenuRouter) getMenu(c *gin.Context) {
  var f interface{}
	  

  if err := c.BindJSON(&f); err != nil {
	log.Println(err)
  }

  log.Println("F", f)

  c.JSON(http.StatusOK, gin.H{"message": "get menu"})
}
func (r *MenuRouter) addMenu(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"message": "add menu"})
}
func (r *MenuRouter) updateMenu(c *gin.Context) {
  c.JSON(http.StatusOK, gin.H{"message": "update menu"})
}
