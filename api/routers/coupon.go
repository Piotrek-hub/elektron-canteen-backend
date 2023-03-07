package routers

import (
	"elektron-canteen/api/controllers"
	"elektron-canteen/api/mid"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
	"strconv"
)

type CouponRouter struct {
	router     *gin.Engine
	controller controllers.CouponController
}

func NewCouponRouter(r *gin.Engine, c controllers.CouponController) *CouponRouter {
	return &CouponRouter{
		router:     r,
		controller: c,
	}
}

func (r *CouponRouter) Initialize() {
	cr := r.router.Group("/coupon")
	cr.Use(mid.Auth())

	cr.POST("/generate/:value", r.generateCoupon)
	cr.POST("/redeem/:code", r.redeemCoupon)
}

func (r *CouponRouter) redeemCoupon(c *gin.Context) {
	code := c.Param("code")

	userID, err := primitive.ObjectIDFromHex(c.Request.Header["user_id"][0])
	if err != nil {
		responseWithError(c, err)
		return
	}

	if err = r.controller.Redeem(userID, code); err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "coupon redeemed successfully"})

}

func (r *CouponRouter) generateCoupon(c *gin.Context) {
	value, err := strconv.ParseFloat(c.Param("value"), 32)
	if err != nil {
		responseWithError(c, err)
		return
	}
	coupon, err := r.controller.Create(float32(value))
	if err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "coupon created successfully", "couponID": coupon.Code})

}
