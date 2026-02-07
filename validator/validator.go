package validator

import (
	form "github.com/go-playground/form/v4"
	playValidator "github.com/go-playground/validator/v10"
)

type Validator struct {
	validator   *playValidator.Validate
	formEncoder *form.Encoder
	formDecoder *form.Decoder
}

var instance *Validator

func Initialize() {
	instance = &Validator{
		validator:   playValidator.New(),
		formEncoder: form.NewEncoder(),
		formDecoder: form.NewDecoder(),
	}

	registerCustomTypes()
}

func Struct(object any) error {
	return instance.validator.Struct(object)
}

func FormEncode(object any) (map[string][]string, error) {
	return instance.formEncoder.Encode(object)
}

func FormEncodeToQueryString(object any) (string, error) {
	values, err := instance.formEncoder.Encode(object)
	if err != nil {
		return "", err
	}

	return values.Encode(), nil
}

func FormDecode(object any, values map[string][]string) error {
	return instance.formDecoder.Decode(object, values)
}
