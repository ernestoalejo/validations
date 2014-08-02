package validations

import (
	"reflect"
	"time"

	"github.com/juju/errgo"
)

type TimeFunc func(s time.Time) (time.Time, error)

func Time(field *time.Time, validations []TimeFunc) Func {
	return func() error {
		privField := time.Time{}

		for _, validation := range validations {
			if newField, err := validation(privField); err != nil {
				return errgo.Mask(err)
			} else {
				privField = newField
			}
		}

		reflect.ValueOf(field).Elem().Set(reflect.ValueOf(privField))

		return nil
	}
}

func ParseTime(source, format string) TimeFunc {
	return func(value time.Time) (time.Time, error) {
		parsed, err := time.Parse(format, source)
		if err != nil {
			return time.Time{}, errgo.Mask(err)
		}

		return parsed, nil
	}
}

func NotBefore(otherTime time.Time) TimeFunc {
	return func(value time.Time) (time.Time, error) {
		if value.Before(otherTime) {
			return time.Time{}, errgo.Newf("not before")
		}

		return value, nil
	}
}
