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

func (v Validator) ValidatePassword(passwd string) error {
	if len(passwd) < 6 {
		return errors.New("password is too short")
	}
	var containsCapitalLetter, containsNormalLetter, containsNumber bool

	for _, l := range passwd {
		if l >= 65 && l <= 90 {
			containsCapitalLetter = true
			continue
		}

		if l >= 97 && l <= 122 {
			containsNormalLetter = true
			continue
		}

		if l >= 48 && l <= 57 {
			containsNumber = true
			continue
		}
	}

	if !containsCapitalLetter {
		return errors.New("password doesn't contains capital letter")
	}
	if !containsNormalLetter {
		return errors.New("password doesn't contains normal letter")
	}
	if !containsNumber {
		return errors.New("password doesn't contains number")
	}

	return nil

}
