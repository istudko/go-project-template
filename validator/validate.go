package validator

import (
	"github.com/go-playground/validator/v10"
)

var v *validator.Validate

func Init() {
	v = validator.New()
}

func Validate(r interface{}) error {
	err := v.Struct(r)
	if err != nil && len(err.(validator.ValidationErrors)) > 0 {
		return err
	}
	return nil
}
