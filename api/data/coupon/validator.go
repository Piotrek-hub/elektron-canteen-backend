package coupon

import "errors"

type Validator struct {
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v Validator) ValidateCoupon(c Coupon) error {
	if c.Value == 0 {
		return errors.New("coupon value = 0")
	}
	if c.Code == "" {
		return errors.New("empty code")
	}

	return nil
}
