package validator

import (
	"github.com/google/uuid"
	"github.com/sdkopen/sdkopen-go-webbase/types"
)

func registerCustomTypes() {
	registerUUIDCustomType()
	registerDateTimeCustomType()
	registerIsoDateCustomType()
	registerIsoTimeCustomType()
}

func registerUUIDCustomType() {
	instance.formDecoder.RegisterCustomTypeFunc(func(vals []string) (any, error) {
		return uuid.Parse(vals[0])
	}, uuid.UUID{})
}

func registerDateTimeCustomType() {
	instance.formDecoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return types.ParseDateTime(vals[0])
	}, types.DateTime{})
}

func registerIsoDateCustomType() {
	instance.formDecoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return types.ParseIsoDate(vals[0])
	}, types.IsoDate{})
}

func registerIsoTimeCustomType() {
	instance.formDecoder.RegisterCustomTypeFunc(func(vals []string) (interface{}, error) {
		return types.ParseIsoTime(vals[0])
	}, types.IsoTime{})
}
