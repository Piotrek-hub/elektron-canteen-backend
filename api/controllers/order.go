package controllers

import (
	"context"
	"elektron-canteen/api/controllers/utils"
	"elektron-canteen/api/data/menu"
	"elektron-canteen/api/data/order"
	"errors"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

type OrderController struct {
	order     order.Model
	menu      menu.Model
	validator *order.Validator
}

func NewOrderController() *OrderController {
	return &OrderController{
		order:     order.Instance(),
		menu:      menu.Instance(),
		validator: order.NewValidator(),
	}
}

func (c *OrderController) AddOrder(no order.NewOrder) (primitive.ObjectID, error) {
	ctx := context.Background()

	no.Status = order.WAITING

	if err := c.validator.ValidateOrder(no); err != nil {
		return primitive.ObjectID{}, err
	}

	if err := c.validator.ValidateUnixDate(no.DueTime); err != nil {
		return primitive.ObjectID{}, err
	}

	todayMenu, err := c.menu.QueryByDay(ctx, utils.UnixToFormattedDate(no.DueTime))
	if err != nil {
		return primitive.ObjectID{}, err
	}

	log.Println(utils.UnixToFormattedDate(no.DueTime))
	var isFound = false
	for _, am := range todayMenu.AvailableMeals {
		if no.Meal.Hex() == am {
			isFound = true
			break
		}
	}

	if !isFound {
		return primitive.ObjectID{}, errors.New("meal is not available")
	}

	return c.order.Create(ctx, no)
}

func (c *OrderController) UpdateOrderStatus(orderID primitive.ObjectID, status string) error {
	ctx := context.Background()

	if status != order.WAITING && status != order.CANCELED && status != order.ACCEPTED && status != order.DECLINED && status != order.DONE {
		return errors.New("Wrong order status")
	}

	o, err := c.order.QueryByID(ctx, orderID)
	if err != nil {
		return err
	}

	if o.Status == order.CANCELED {
		return errors.New("can't change order status, order is cancelled")
	}

	return c.order.UpdateStatus(ctx, orderID, status)
}

func (c *OrderController) GetAllOrders() ([]order.Order, error) {
	ctx := context.Background()
	return c.order.QueryAll(ctx)
}

func (c *OrderController) GetOrdersByDate(date string) ([]order.Order, error) {
	ctx := context.Background()

	if err := c.validator.ValidateDate(date); err != nil {
		return nil, err
	}

	return c.order.QueryByDate(ctx, date)
}

func (c *OrderController) GetOrder(orderID primitive.ObjectID) (*order.Order, error) {
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
