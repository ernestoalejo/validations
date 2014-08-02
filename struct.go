package validations

import (
	"reflect"

	"github.com/juju/errgo"
)

type StructFunc func(s interface{}) error

func Struct(field interface{}, validations []StructFunc) Func {
	return func() error {
		for _, validation := range validations {
			if err := validation(field); err != nil {
				return errgo.Mask(err)
			}
		}

		return nil
	}
}

func NotNil() StructFunc {
	return func(value interface{}) error {
		reflectedValue := reflect.ValueOf(value)
		if reflectedValue.Kind() != reflect.Ptr {
			return errgo.Newf("not a pointer")
		}

		if reflectedValue.IsNil() {
			return errgo.Newf("not nil")
		}

		return nil
	}
}
