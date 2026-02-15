package validator

import (
	"github.com/google/uuid"
)

func registerCustomTypes() {
	registerUUIDCustomType()
}

func registerUUIDCustomType() {
	instance.formDecoder.RegisterCustomTypeFunc(func(vals []string) (any, error) {
		return uuid.Parse(vals[0])
	}, uuid.UUID{})
}
