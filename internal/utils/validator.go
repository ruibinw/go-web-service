package utils

import (
	"errors"
	"github.com/go-playground/validator/v10"
	"strings"
)

type Validator interface {
	Validate(i any) error
}

type RequestValidator struct {
	validator *validator.Validate
}

func NewRequestValidator(v *validator.Validate) Validator {
	return &RequestValidator{validator: v}
}

func (s RequestValidator) Validate(i any) error {
	if err := s.validator.Struct(i); err != nil {
		return newValidationDetailsError(err)
	}
	return nil
}

func newValidationDetailsError(err error) error {
	errs := err.(validator.ValidationErrors)
	msg := make([]string, len(errs))
	for i, e := range errs {
		msg[i] = e.Field() + getMsgByTag(e.Tag())
	}
	return errors.New(strings.Join(msg, ", "))
}

func getMsgByTag(tag string) string {
	switch tag {
	case "required":
		return ": input is required"
	}
	return ": input invalid" // default error
}
