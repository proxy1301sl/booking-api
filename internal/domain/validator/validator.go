package validator

import "github.com/go-playground/validator/v10"

type Validatorr struct {
	validator *validator.Validate
}

func New() *Validatorr {
	return &Validatorr{validator: validator.New()}
}

func (v *Validatorr) Validate(s any) error {
	if err := v.validator.Struct(s); err != nil {
		return err
	}
	return nil

}
