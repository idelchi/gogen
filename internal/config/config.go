package config

import (
	"errors"
	"fmt"

	"github.com/idelchi/gogen/pkg/validator"
)

// ErrUsage indicates an error in command-line usage or configuration.
var ErrUsage = errors.New("usage error")

// Generate holds parameters for key generation.
type Generate struct {
	// Length specifies the key length in bytes (32-512, must be multiple of 4)
	Length int `validate:"min=32,max=512,multiple=4"`
}

// Password holds parameters for password hashing operations.
type Password struct {
	// Password is the input password to be hashed
	Password string `mapstructure:"-"`
	// Cost is the bcrypt work factor (4-31)
	Cost int `validate:"min=4,max=31"`
	// Benchmark indicates whether to run performance benchmarks
	Benchmark bool
}

// Config holds the application's configuration parameters.
type Config struct {
	// Show enables output display
	Show bool
	// Generate contains key generation settings
	Generate Generate `mapstructure:",squash"`
	// Password contains password hashing settings
	Password Password `mapstructure:",squash"`
}

// Display returns the value of the Show field.
func (c Config) Display() bool {
	return c.Show
}

// Validate performs configuration validation using the validator package.
// It returns a wrapped ErrUsage if any validation rules are violated.
func Validate(config any) error {
	validator := validator.NewValidator()

	if err := registerMultiple32(validator); err != nil {
		return fmt.Errorf("registering multiple32: %w", err)
	}

	errs := validator.Validate(config)

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
