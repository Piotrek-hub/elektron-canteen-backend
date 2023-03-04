package routers

import (
	"elektron-canteen/api/controllers"
	"elektron-canteen/api/data/order"
	"elektron-canteen/api/data/user"
	"elektron-canteen/api/mid"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type OrderRouter struct {
	router     *gin.Engine
	controller controllers.OrderController
}

func NewOrderRouter(r *gin.Engine, c controllers.OrderController) *OrderRouter {
	return &OrderRouter{
		router:     r,
		controller: c,
	}
}

func (r *OrderRouter) Initialize() {
	r.router.Use(mid.Auth())

	r.router.GET("/orders/user/:id", r.getUserOrders)
	r.router.GET("/orders/:id", r.getOrder)
	r.router.POST("/orders", r.addOrder)
	r.router.POST("/orders/cancel/:order_id", r.cancelOrder)

	authorizedRoutes := r.router.Group("/admin")
	authorizedRoutes.Use(mid.Role(user.ADMIN_ROLE))
	authorizedRoutes.GET("/", r.getAllOrders)
	authorizedRoutes.PATCH("/order/:id/:status", r.updateOrderStatus)

}

func (r *OrderRouter) cancelOrder(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Request.Header["user_id"][0])
	if err != nil {
		panic(err)
	}

	orderID, err := primitive.ObjectIDFromHex(c.Param("order_id"))
	if err != nil {
		panic(err)
	}

	if err = r.controller.CancelOrder(userID, orderID); err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order cancelled successfully"})
}

func (r *OrderRouter) updateOrderStatus(c *gin.Context) {
	orderID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		responseWithError(c, err)
		return
	}
	status := c.Param("status")

	if err := r.controller.UpdateOrderStatus(orderID, status); err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated sucessfully"})

}

func (r *OrderRouter) getOrder(c *gin.Context) {
	orderID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		responseWithError(c, err)
		return
	}

	order, err := r.controller.GetOrder(orderID)
	if err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func (r *OrderRouter) getUserOrders(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("id"))
	if err != nil {
		responseWithError(c, err)
		return
	}

	orders, err := r.controller.GetUserOrders(userID)
	if err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func (r *OrderRouter) getAllOrders(c *gin.Context) {
	orders, err := r.controller.GetAllOrders()
	if err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func (r *OrderRouter) addOrder(c *gin.Context) {
	var no order.NewOrder
	if err := c.BindJSON(&no); err != nil {
		responseWithError(c, err)
		return
	}

	orderID, err := r.controller.AddOrder(no)
	if err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "order added successfully", "orderId": orderID.Hex()})

}
