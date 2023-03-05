package order

import (
	"elektron-canteen/api/controllers/utils"
	"errors"
	"regexp"
	"time"
)

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

func (v *Validator) ValidateUnixDate(unix string) error {
	d := utils.UnixToDate(unix)
	if d.Before(time.Now()) {
		return errors.New("order date is in the past")
	}

	return nil
}

func (v *Validator) ValidateDate(date string) error {
	re := regexp.MustCompile("(?i)\\d\\d\\d\\d-\\d\\d-\\d\\d")
	if !re.MatchString(date) {
		return errors.New("wrong date format")
	}

	return nil
}
