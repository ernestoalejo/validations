package validations

import (
	"github.com/juju/errgo"
)

type FloatFunc func(s float64) error

func Float(field float64, validations []FloatFunc) Func {
	return func() error {
		for _, validation := range validations {
			if err := validation(field); err != nil {
				return errgo.Mask(err)
			}
		}

		return nil
	}
}

func FloatRange(minvalue, maxvalue float64) FloatFunc {
	return func(value float64) error {
		if value < minvalue || value > maxvalue {
			return errgo.Newf("range")
		}

		return nil
	}
}

func FloatPositive() FloatFunc {
	return func(value float64) error {
		if value < 1 {
			return errgo.Newf("positive")
		}

		return nil
	}
}

func FloatPositiveZero() FloatFunc {
	return func(value float64) error {
		if value < 0 {
			return errgo.Newf("positive zero")
		}

		return nil
	}
}
