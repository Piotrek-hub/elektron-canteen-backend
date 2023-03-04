package order

import "errors"

type Validator struct {
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v *Validator) ValidateOrder(order NewOrder) error {
	if order.Meal.String() == "" {
		return errors.New("meal is missing")
	}
	if order.User.String() == "" {
		return errors.New("user is missing")
	}
	if order.Status == "" {
		return errors.New("status is missing")
	}
	if order.PaymentMethod == "" {
		return errors.New("paymentMethod is missing")
	}
	if order.DueTime == "" {
		return errors.New("time")
	}

	return nil
}
