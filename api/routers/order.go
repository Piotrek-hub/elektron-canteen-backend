package routers

import (
	"elektron-canteen/api/controllers"
	"elektron-canteen/api/data/order"
	"elektron-canteen/api/data/user"
	"elektron-canteen/api/mid"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
)

type OrderRouter struct {
	router       *gin.Engine
	controller   controllers.OrderController
	orderChannel chan *order.Response
}

func NewOrderRouter(r *gin.Engine, c controllers.OrderController) *OrderRouter {
	return &OrderRouter{
		router:       r,
		controller:   c,
		orderChannel: make(chan *order.Response),
	}
}

func (r *OrderRouter) Initialize() {
	or := r.router.Group("/order")
	or.Use(mid.Auth())

	or.GET("/:order_id", r.getOrder)
	or.GET("/date/:date", r.getOrdersByDate)
	or.GET("/user/:user_id", r.getUserOrders)
	or.POST("/add", r.createOrder)
	or.POST("/cancel/:order_id", r.cancelOrder)

	or.GET("/all", mid.Role(user.ADMIN_ROLE), r.getAllOrders)
	or.GET("/all/ws", mid.Role(user.ADMIN_ROLE), r.getAllOrdersLive)
	or.PATCH("/:order_id/:status", mid.Role(user.ADMIN_ROLE), r.updateOrderStatus)

}

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func (r *OrderRouter) getAllOrdersLive(c *gin.Context) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		responseWithError(c, err)
		return
	}
	defer ws.Close()

	go r.controller.ListenForOrders(r.orderChannel)
	for {
		no := <-r.orderChannel
		if no == nil {
			responseWithError(c, errors.New("error with websocket"))
			return
		}

		orderJson, err := json.Marshal(no)
		if err != nil {
			responseWithError(c, err)
			return
		}

		ws.WriteMessage(1, orderJson)
	}
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

	orderID, err := r.controller.AddOrder(no)
	if err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "order added successfully", "order_id": orderID.Hex()})
}
