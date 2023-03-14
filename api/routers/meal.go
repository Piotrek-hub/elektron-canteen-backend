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
	mr := r.router.Group("/meal")
	mr.GET("/", r.getMeals)
	mr.GET("/meal/:id", r.getMeal)

	mr.POST("", mid.Role(user.ADMIN_ROLE), r.createMeal)
	mr.PATCH("/:id", mid.Role(user.ADMIN_ROLE), r.updateMeal)
	mr.DELETE("/:id", mid.Role(user.ADMIN_ROLE), r.deleteMeal)
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

func (r *MealRouter) createMeal(c *gin.Context) {
	var nm meal.NewMeal
	if err := c.BindJSON(&nm); err != nil {
		responseWithError(c, err)
		return
	}

	if err := r.controller.Add(nm); err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{
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

func (r *MealRouter) deleteMeal(c *gin.Context) {
	mealID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		responseWithError(c, err)
		return
	}

	if err := r.controller.Delete(mealID); err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Meal deleted successfully"})

}
