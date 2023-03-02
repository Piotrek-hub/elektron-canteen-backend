package routers

import (
	"elektron-canteen/api/controllers"
	"elektron-canteen/api/data/menu"
	"elektron-canteen/api/mid"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MenuRouter struct {
	router     *gin.Engine
	controller controllers.MenuController
}

func NewMenuRouter(r *gin.Engine, c controllers.MenuController) *MenuRouter {
	return &MenuRouter{
		router:     r,
		controller: c,
	}
}

func (r *MenuRouter) Initialize() {
	r.router.Use(mid.Auth())
	r.router.GET("/menu", r.getMenu)
	r.router.POST("/menu", r.addMenu)
	r.router.PATCH("/menu", r.updateMenu)
}

func (r *MenuRouter) getMenu(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "get menu"})
}

func (r *MenuRouter) addMenu(c *gin.Context) {
	var mr menu.AddRequest

	if err := c.BindJSON(&mr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	log.Println("menu", mr)

	c.JSON(http.StatusOK, gin.H{"message": "add menu"})
}

func (r *MenuRouter) updateMenu(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "update menu"})
}
