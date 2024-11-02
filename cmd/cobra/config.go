package main

import (
	"errors"
	"fmt"

	"github.com/idelchi/gogen/pkg/validator"
)

var ErrUsage = errors.New("usage error")

type Generate struct {
	Length int
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

func (c *Config) Validate() error {
	errs := validator.NewValidator().Validate(c)

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
