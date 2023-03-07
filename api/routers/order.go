package routers

import (
	"elektron-canteen/api/controllers"
	"elektron-canteen/api/controllers/utils"
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

	r.router.GET("/orders/:order_id", r.getOrder)
	r.router.GET("/orders/date/:date", r.getOrdersByDate)
	r.router.GET("/orders/user/:user_id", r.getUserOrders)
	r.router.POST("/orders/add", r.createOrder)
	r.router.POST("/orders/cancel/:order_id", r.cancelOrder)

	r.router.GET("/orders/all", mid.Role(user.ADMIN_ROLE), r.getAllOrders)
	r.router.PATCH("/order/:order_id/:status", mid.Role(user.ADMIN_ROLE), r.updateOrderStatus)

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
	orderID, err := primitive.ObjectIDFromHex(c.Param("order_id"))
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
	orderID, err := primitive.ObjectIDFromHex(c.Param("order_id"))
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

func (r *OrderRouter) getOrdersByDate(c *gin.Context) {
	date := c.Param("date")

	orders, err := r.controller.GetOrdersByDate(date)
	if err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"date": date, "orders": orders})
}

func (r *OrderRouter) getUserOrders(c *gin.Context) {
	userID, err := primitive.ObjectIDFromHex(c.Param("user_id"))
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

func (r *OrderRouter) createOrder(c *gin.Context) {
	var no order.NewOrder
	if err := c.BindJSON(&no); err != nil {
		responseWithError(c, err)
		return
	}

	userID, err := primitive.ObjectIDFromHex(c.Request.Header["user_id"][0])
	if err != nil {
		responseWithError(c, err)
		return
	}

	no.User = userID
	no.Date = utils.UnixToFormattedDate(no.DueTime)

	orderID, err := r.controller.AddOrder(no)
	if err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "order added successfully", "order_id": orderID.Hex()})
}
