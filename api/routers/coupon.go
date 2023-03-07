package routers

import (
	"elektron-canteen/api/controllers"
	"elektron-canteen/api/data/coupon"
	"elektron-canteen/api/data/user"
	"elektron-canteen/api/mid"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"net/http"
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

	cr.POST("/generate", mid.Role(user.ADMIN_ROLE), r.generateCoupon)
	cr.POST("/redeem", r.redeemCoupon)
}

func (r *CouponRouter) redeemCoupon(c *gin.Context) {
	var coupon coupon.Coupon
	if err := c.BindJSON(&coupon); err != nil {
		responseWithError(c, err)
		return
	}

	userID, err := primitive.ObjectIDFromHex(c.Request.Header["user_id"][0])
	if err != nil {
		responseWithError(c, err)
		return
	}

	if err = r.controller.Redeem(userID, coupon.Code); err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "coupon redeemed successfully"})

}

func (r *CouponRouter) generateCoupon(c *gin.Context) {
	var coupon coupon.Coupon
	if err := c.BindJSON(&coupon); err != nil {
		responseWithError(c, err)
		return
	}

	createdCoupon, err := r.controller.Create(coupon.Value)
	if err != nil {
		responseWithError(c, err)
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "coupon created successfully", "couponID": createdCoupon.Code})

}
