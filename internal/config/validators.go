package config

import (
	"fmt"
	"strconv"

	"github.com/idelchi/gogen/pkg/validator"
)

// registerMultiple adds a custom validator ensuring values are multiples of the given value.
// It registers both the validation logic and a human-readable error message.
func registerMultiple(validator *validator.Validator) error {
	if err := validator.RegisterValidationAndTranslation(
		"multiple",
		validateMultiple,
		"{0} must be a multiple of {1}",
	); err != nil {
		return fmt.Errorf("registering validation: %w", err)
	}

	return nil
}

// validateMultiple checks if a field's value is a multiple of the number specified in the tag.
// Returns true if the value is valid, false otherwise.
func validateMultiple(fl validator.FieldLevel) bool {
	value := fl.Field().Int()

	// Get the parameter from the tag (the number after 'multiple=')
	param := fl.Param()

	multiplier, err := strconv.Atoi(param)
	if err != nil {
		return false // Invalid parameter
	}

	// Check if value is a multiple of the specified number
	return value%int64(multiplier) == 0
}
