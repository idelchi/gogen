package config

import (
	"fmt"
	"os"

	"github.com/idelchi/godyl/pkg/pretty"
	"github.com/idelchi/gogen/pkg/validator"
	"github.com/spf13/viper"
)

// validate unmarshals the configuration and performs validation checks.
// If cfg.Show is true, prints the configuration and exits.
func validate(cfg *Config, validations ...any) error {
	if err := viper.Unmarshal(cfg); err != nil {
		return fmt.Errorf("unmarshalling config: %w", err)
	}

	if cfg.Show {
		pretty.PrintJSONMasked(cfg)
		os.Exit(0)
		return nil
	}

	for _, v := range validations {
		if err := Validate(v); err != nil {
			return fmt.Errorf("validating config: %w\nSee --help for more info on usage.", err)
		}
	}

	return nil
}

// registerMultiple32 adds a custom validator ensuring values are multiples of 32 bits (4 bytes).
// It registers both the validation logic and a human-readable error message.
func registerMultiple32(validator *validator.Validator) error {
	validator.Validator().RegisterValidation("multiple", validateMultiple32)

	if err := validator.RegisterValidationAndTranslation(
		"multiple",
		validateMultiple32,
		"{0} must be a multiple of {1} bytes",
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
	if value%4 != 0 {
		return false
	}

	return true
}
