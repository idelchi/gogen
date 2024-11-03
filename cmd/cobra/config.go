package main

import (
	"errors"
	"fmt"

	"github.com/idelchi/gogen/pkg/validator"
)

var ErrUsage = errors.New("usage error")

type Generate struct {
	Length int `validate:"min=32,max=512,multiple=4"`
}

type Password struct {
	Password  string `mapstructure:"-"`
	Cost      int    `validate:"min=4,max=31"`
	Benchmark bool
}

type Config struct {
	Show bool

	Generate Generate `mapstructure:",squash"`
	Password Password `mapstructure:",squash"`
}

func Validate(c any) error {
	validator := validator.NewValidator()
	validator.Validator().RegisterValidation("multiple", validateMultiple32)

	// Register the multiple validation with a nice error message
	if err := validator.RegisterValidationAndTranslation(
		"multiple",
		validateMultiple32,
		"{0} must be a multiple of {1} bytes",
	); err != nil {
		return fmt.Errorf("registering validation: %w", err)
	}

	errs := validator.Validate(c)

	switch {
	case errs == nil:
		return nil
	case len(errs) == 1:
		return fmt.Errorf("%w: %w", ErrUsage, errs[0])
	case len(errs) > 1:
		return fmt.Errorf("%ws:\n%w", ErrUsage, errors.Join(errs...))

	}

	return nil
}

func validateMultiple32(fl validator.FieldLevel) bool {
	value := fl.Field().Int()

	// Must be a multiple of 32 bits (4 bytes)
	if value%4 != 0 {
		return false
	}

	return true
}
