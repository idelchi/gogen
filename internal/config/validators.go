package config

import (
	"fmt"

	"github.com/idelchi/gogen/pkg/validator"
)

// registerMultiple32 adds a custom validator ensuring values are multiples of 32 bits (4 bytes).
// It registers both the validation logic and a human-readable error message.
func registerMultiple32(validator *validator.Validator) error {
	if err := validator.RegisterValidationAndTranslation(
		"multiple",
		validateMultiple32,
		"{0} must be a multiple of {1}",
	); err != nil {
		return fmt.Errorf("registering validation: %w", err)
	}

	return nil
}

// validateMultiple32 checks if a field's value is a multiple of 4 bytes (32 bits).
// Returns true if the value is valid, false otherwise.
func validateMultiple32(fl validator.FieldLevel) bool {
	value := fl.Field().Int()

	// Check if value is a multiple of 4 bytes
	return value%4 == 0
}
