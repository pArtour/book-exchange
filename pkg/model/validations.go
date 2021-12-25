package model

import validation "github.com/go-ozzo/ozzo-validation"

func RequiredIf(c bool) validation.RuleFunc {
	return func(v interface{}) error {
		if c == true {
			return validation.Validate(v, validation.Required)
		}
		return nil
	}
}
