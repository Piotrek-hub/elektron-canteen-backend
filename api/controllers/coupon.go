package controllers

import (
	"context"
	"elektron-canteen/api/controllers/utils"
	"elektron-canteen/api/data/coupon"
	"elektron-canteen/api/data/user"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CouponController struct {
	coupon    coupon.Model
	user      user.Model
	validator coupon.Validator
}

func NewCouponController() *CouponController {
	return &CouponController{
		coupon:    coupon.Instance(),
		user:      user.Instance(),
		validator: *coupon.NewValidator(),
	}
}

const COUPON_LENGTH = 6

func (c *CouponController) Create(value float32) (*coupon.Coupon, error) {
	ctx := context.Background()

	//c.validator.ValidateCoupon(coupon)

	var coupon coupon.Coupon
	coupon.Value = value
	coupon.Code = utils.RandomString(COUPON_LENGTH)

	if _, err := c.coupon.Create(ctx, coupon); err != nil {
		return nil, err
	}

	return &coupon, nil
}

func (c *CouponController) Redeem(userID primitive.ObjectID, code string) error {
	ctx := context.Background()

	coupon, err := c.coupon.QueryByCode(ctx, code)
	if err != nil {
		return err
	}

	user, err := c.user.QueryByID(ctx, userID)
	if err != nil {
		return err
	}

	newPoints := user.Points + coupon.Value
	err = c.user.UpdatePoints(ctx, user.ID, newPoints)
	if err != nil {
		return err
	}

	if err = c.coupon.Delete(ctx, coupon.Code); err != nil {
		return err
	}

	return nil
}
