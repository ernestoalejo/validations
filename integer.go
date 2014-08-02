package validations

import (
	"github.com/juju/errgo"
)

type IntFunc func(s int) error

func Int(field int, validations []IntFunc) Func {
	return func() error {
		for _, validation := range validations {
			if err := validation(field); err != nil {
				return errgo.Mask(err)
			}
		}

		return nil
	}
}

func Range(minvalue, maxvalue int) IntFunc {
	return func(value int) error {
		if value < minvalue || value > maxvalue {
			return errgo.Newf("range")
		}

		return nil
	}
}

func Positive() IntFunc {
	return func(value int) error {
		if value < 1 {
			return errgo.Newf("positive")
		}

		return nil
	}
}

func PositiveZero() IntFunc {
	return func(value int) error {
		if value < 0 {
			return errgo.Newf("positive zero")
		}

		return nil
	}
}

func MaxValue(max int) IntFunc {
	return func(value int) error {
		if value > max {
			return errgo.Newf("max value")
		}

		return nil
	}
}
