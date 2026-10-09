package validator

import "github.com/go-playground/validator/v10"

type Validator struct {
	validator *validator.Validate
}

func New() *Validator {
	return &Validator{validator: validator.New()}
}

func (v *Validator) Validate(s any) error {
	if err := v.validator.Struct(s); err != nil {
		return err
	}
	return nil

}
