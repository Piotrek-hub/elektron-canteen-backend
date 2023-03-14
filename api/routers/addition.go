package routers

import (
	"elektron-canteen/api/controllers"
	"elektron-canteen/api/data/addition"
	"elektron-canteen/api/data/user"
	"elektron-canteen/api/mid"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type AdditionRouter struct {
	router     *gin.Engine
	controller controllers.AdditionController
}

func NewAdditionRouter(r *gin.Engine, c controllers.AdditionController) *AdditionRouter {
	return &AdditionRouter{
		router:     r,
		controller: c,
	}
}

func (r *AdditionRouter) Initialize() {
	ar := r.router.Group("/addition")

	ar.GET("/id/:id", r.getById)
	ar.GET("/name/:name", r.getByName)
	ar.GET("/all", r.getAll)

	ar.POST("/create", mid.Auth(), mid.Role(user.ADMIN_ROLE), r.createAddition)
	ar.POST("/update/:id", mid.Auth(), mid.Role(user.ADMIN_ROLE))
	ar.DELETE("/delete/:id", mid.Auth(), mid.Role(user.ADMIN_ROLE), r.deleteAddition)

}

func (r *AdditionRouter) getAll(c *gin.Context) {
	additions, err := r.controller.GetAll()
	if err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"additions": additions})
}

func (r *AdditionRouter) getById(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}

	a, err := r.controller.GetById(id)
	if err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"addition": a})
}

func (r *AdditionRouter) getByName(c *gin.Context) {
	name := c.Param("name")

	a, err := r.controller.GetByName(name)
	if err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"addition": a})
}

func (r *AdditionRouter) createAddition(c *gin.Context) {
	var na addition.NewAddition
	if err := c.BindJSON(&na); err != nil {
		responseWithError(c, err)
		return
	}

	id, err := r.controller.Create(na)
	if err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Addition created successfully", "addition_id": id.Hex()})
}

func (r *AdditionRouter) deleteAddition(c *gin.Context) {
	id, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		panic(err)
	}

	if err := r.controller.Delete(id); err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "addition deleted"})
}
