package validations

import (
	"reflect"

	"github.com/juju/errgo"
)

type SliceFunc func(s interface{}) error

func Slice(field interface{}, validations []SliceFunc) Func {
	return func() error {
		for _, validation := range validations {
			if err := validation(field); err != nil {
				return errgo.Mask(err)
			}
		}

		return nil
	}
}

func SliceMinLen(length int) SliceFunc {
	return func(value interface{}) error {
		reflectedValue := reflect.ValueOf(value)
		if reflectedValue.Kind() != reflect.Slice {
			return errgo.Newf("not a slice")
		}

		if reflectedValue.Len() < length {
			return errgo.Newf("slice minlen")
		}

		return nil
	}
}

func Each(fn interface{}) SliceFunc {
	return func(slice interface{}) error {
		fnType := reflect.TypeOf(fn)
		sliceType := reflect.TypeOf(slice)
		fnValue := reflect.ValueOf(fn)
		sliceValue := reflect.ValueOf(slice)

		if fnType.Kind() != reflect.Func {
			return errgo.Newf("fn not a function")
		}

		if sliceType.Kind() != reflect.Slice {
			return errgo.Newf("value not a slice")
		}

		if fnType.NumIn() != 1 || fnType.NumOut() != 1 {
			return errgo.Newf("fn should receive and element and return an error")
		}

		if fnType.In(0) != sliceType.Elem() {
			return errgo.Newf("fn should receive a param with the same type as the slice")
		}

		if fnType.Out(0) != reflect.TypeOf((*error)(nil)).Elem() {
			return errgo.Newf("fn should return an error")
		}

		total := sliceValue.Len()
		for i := 0; i < total; i++ {
			result := fnValue.Call([]reflect.Value{sliceValue.Index(i)})
			if !result[0].IsNil() {
				return result[0].Interface().(error)
			}
		}

		return nil
	}
}
