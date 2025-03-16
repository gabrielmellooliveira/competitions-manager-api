package common

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func Validate(input any) error {
	validate := validator.New()

	err := validate.Struct(input)
	if err != nil {
		for _, e := range err.(validator.ValidationErrors) {
			return errors.New("campo '" + e.Field() + "' inválido " + e.Tag())
		}
	}

	return nil
}
