package user

import "errors"

type Validator struct {

}

func NewValidator() *Validator {
	return &Validator{}
}

func (v Validator) ValidateUser(user NewUser) error {
	if user.Email == "" {
		return errors.New("email is missing")
	}
  	if user.Name == "" {
		return errors.New("name is missing")
	}
	if user.Surname == "" {
		return errors.New("surname is missing")
	}
	if user.Password == "" {
		return errors.New("password is missing")
	}


	return nil
}
