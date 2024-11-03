// Package validator provides a wrapper around go-playground/validator with
// built-in translation support for validation error messages. It simplifies
// struct validation and error handling by providing formatted, human-readable
// error messages in English.
package validator

import (
	"errors"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

type FieldLevel = validator.FieldLevel

// Validator wraps the go-playground/validator functionality with translation support.
// It holds both the validator instance and a translator for converting validation
// errors into human-readable messages.
type Validator struct {
	validate   *validator.Validate
	translator ut.Translator
}

// Validator returns the underlying validator instance.
// This method provides access to the raw validator when needed for advanced usage.
func (v *Validator) Validator() *validator.Validate {
	return v.validate
}

// NewValidator creates and initializes a new Validator instance with English translations.
// It sets up the universal translator with English locale and registers the default
// English translations for validation error messages.
func NewValidator() *Validator {
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("en")

	validate := validator.New()
	en_translations.RegisterDefaultTranslations(validate, trans)

	return &Validator{
		validate:   validate,
		translator: trans,
	}
}

// FormatErrors converts validator.ValidationErrors into a slice of standard errors.
// It translates the validation errors into human-readable messages using the
// configured translator. Returns nil if there are no errors to format.
func (v *Validator) FormatErrors(err error) []error {
	errs := err.(validator.ValidationErrors)

	if len(errs) == 0 {
		return nil
	}

	errMap := errs.Translate(v.translator)

	totalErrors := make([]error, 0, len(errMap))

	for _, err := range errMap {
		totalErrors = append(totalErrors, errors.New(err))
	}

	return totalErrors
}

// Validate performs validation on the provided struct and returns any validation errors.
// It uses struct tags to determine validation rules and returns formatted error messages.
// If validation passes, it returns nil.
//
// Example usage:
//
//	type User struct {
//	    Name  string `validate:"required"`
//	    Email string `validate:"required,email"`
//	}
//
//	validator := validator.NewValidator()
//	user := User{Name: "John", Email: "invalid-email"}
//	if errs := validator.Validate(user); errs != nil {
//	    // Handle validation errors
//	}
func (v *Validator) Validate(c any) []error {
	if err := v.validate.Struct(c); err != nil {
		return v.FormatErrors(err)
	}

	return nil
}

// RegisterValidationAndTranslation registers both a validation function and its error message translation.
// It simplifies the process of adding custom validations with proper error messages.
//
// Parameters:
//   - tag: the validation tag to use in struct field tags
//   - fn: the validation function that implements the validation logic
//   - msgTemplate: the error message template (use {0} for the field name and {1} for the parameter)
//
// Example usage:
//
//	validator.RegisterValidationAndTranslation(
//	    "multiple",
//	    validateMultiple,
//	    "{0} must be a multiple of {1}"
//	)
func (v *Validator) RegisterValidationAndTranslation(tag string, fn validator.Func, msgTemplate string) error {
	// Register the validation function
	if err := v.validate.RegisterValidation(tag, fn); err != nil {
		return err
	}

	// Register the translation
	return v.validate.RegisterTranslation(tag, v.translator,
		// RegisterTranslation
		func(ut ut.Translator) error {
			return ut.Add(tag, msgTemplate, true)
		},
		// Translation
		func(ut ut.Translator, fe validator.FieldError) string {
			param := fe.Param()
			t, _ := ut.T(tag, fe.Field(), param)
			return t
		},
	)
}
