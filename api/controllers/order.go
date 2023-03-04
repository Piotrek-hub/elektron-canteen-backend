package controllers

import (
	"context"
	"elektron-canteen/api/data/order"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderController struct {
	order     order.Model
	validator *order.Validator
}

func NewOrderController() *OrderController {
	return &OrderController{
		order:     order.Instance(),
		validator: order.NewValidator(),
	}
}

func (c *OrderController) AddOrder(no order.NewOrder) (primitive.ObjectID, error) {
	ctx := context.Background()

	no.Status = order.WAITING

	if err := c.validator.ValidateOrder(no); err != nil {
		return primitive.ObjectID{}, err
	}

	return c.order.Create(ctx, no)
}

func (c *OrderController) UpdateOrderStatus(orderID primitive.ObjectID, status string) error {
	ctx := context.Background()
	return c.order.UpdateStatus(ctx, orderID, status)
}

func (c *OrderController) GetAllOrders() ([]order.Order, error) {
	ctx := context.Background()
	return c.order.QueryAll(ctx)
}

func (c *OrderController) GetOrder(orderID primitive.ObjectID) ([]order.Order, error) {
	ctx := context.Background()
	return c.order.QueryByID(ctx, orderID)
}

func (c *OrderController) CancelOrder(userID, orderID primitive.ObjectID) error {
	ctx := context.Background()
	userOrders, err := c.GetUserOrders(userID)
	if err != nil {
		return err
	}

	for _, uo := range userOrders {
		if uo.ID == orderID {
			if uo.Status == order.WAITING {
				c.order.UpdateStatus(ctx, uo.ID, order.CANCELED)
				return nil
			} else {
				return errors.New("Cannot cancel order, order status: " + uo.Status)
			}
		}
	}

	return errors.New("order not found")
}

func (c *OrderController) GetUserOrders(userID primitive.ObjectID) ([]order.Order, error) {
	ctx := context.Background()
	return c.order.QueryByUser(ctx, userID)

}
