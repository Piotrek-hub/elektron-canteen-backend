package routers

import (
	"elektron-canteen/api/controllers"
	"elektron-canteen/api/data/menu"
	"elektron-canteen/api/data/user"
	"elektron-canteen/api/mid"
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
	r.router.GET("/menu/:day", r.getMenu)

	securedRoutes := r.router.Group("/menu")
	securedRoutes.Use(mid.Role(user.ADMIN_ROLE))

	securedRoutes.POST("", r.addMenu)
	securedRoutes.PATCH("", r.updateMenu)
	securedRoutes.DELETE("/:day", r.deleteMenu)
}

func (r *MenuRouter) getMenu(c *gin.Context) {
	day := c.Param("day")

	m, err := r.controller.GetByDay(day)
	if err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": m})
}

func (r *MenuRouter) deleteMenu(c *gin.Context) {
	day := c.Param("day")

	if err := r.controller.Delete(day); err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "deleted successfully"})
}

func (r *MenuRouter) addMenu(c *gin.Context) {
	var mr menu.AddRequest

	if err := c.BindJSON(&mr); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := r.controller.Add(mr); err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "add menu"})
}

func (r *MenuRouter) updateMenu(c *gin.Context) {
	var m menu.Menu

	if err := c.BindJSON(&m); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := r.controller.Update(m); err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": m})
}
