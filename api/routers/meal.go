package routers

import (
	"elektron-canteen/api/controllers"
	"elektron-canteen/api/data/meal"
	"elektron-canteen/api/data/user"
	"elektron-canteen/api/mid"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type MealRouter struct {
	router     *gin.Engine
	controller controllers.MealController
}

func NewMealRouter(r *gin.Engine, c controllers.MealController) *MealRouter {
	return &MealRouter{
		router:     r,
		controller: c,
	}
}

func (r *MealRouter) Initialize() {
	r.router.Use(mid.Auth())
	r.router.GET("/meal", r.getMeals)
	r.router.GET("/meal/:id", r.getMeal)

	securedRoutes := r.router.Group("/meal")
	securedRoutes.Use(mid.Role(user.ADMIN_ROLE))

	securedRoutes.POST("", r.addMeal)
	securedRoutes.PATCH("/:id", r.updateMeal)
}

func (r *MealRouter) getMeals(c *gin.Context) {
	meals, err := r.controller.GetAll()
	if err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"meals": meals,
	})
}

func (r *MealRouter) getMeal(c *gin.Context) {
	mealID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		responseWithError(c, err)
		return
	}

	m, err := r.controller.GetByID(mealID)
	if err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"meal": m,
	})
}

func (r *MealRouter) addMeal(c *gin.Context) {
	var nm meal.NewMeal
	if err := c.BindJSON(&nm); err != nil {
		responseWithError(c, err)
		return
	}

	if err := r.controller.Add(nm); err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Meal added successfully",
	})
}

func (r *MealRouter) updateMeal(c *gin.Context) {
	mealID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		responseWithError(c, err)
		return
	}

	var nm meal.NewMeal
	if err := c.BindJSON(&nm); err != nil {
		responseWithError(c, err)
		return
	}

	if err := r.controller.Update(mealID, nm); err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Meal updated successfully"})

}
