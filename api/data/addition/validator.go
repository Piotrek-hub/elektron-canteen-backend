package addition

import "errors"

type Validator struct {
}

func NewValidator() *Validator {
	return &Validator{}
}

func (v Validator) ValidateAddition(a Addition) error {
	if a.Name == "" {
		return errors.New("name can't be empty")
	}
	if a.Price < 0 {
		return errors.New("price can't be smaller than 0")
	}

	return nil
}
