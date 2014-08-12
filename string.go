package validations

import (
	"regexp"

	"github.com/juju/errgo"
)

type StrFunc func(s string) error

func Str(field string, validations []StrFunc) Func {
	return func() error {
		for _, validation := range validations {
			if err := validation(field); err != nil {
				return errgo.Mask(err)
			}
		}

		return nil
	}
}

func Required() StrFunc {
	return func(value string) error {
		if value == "" {
			return errgo.Newf("required")
		}

		return nil
	}
}

func MinLength(length int) StrFunc {
	return func(value string) error {
		if value == "" {
			return nil
		}

		if len(value) < length {
			return errgo.Newf("minlength")
		}

		return nil
	}
}

func MaxLength(length int) StrFunc {
	return func(value string) error {
		if len(value) > length {
			return errgo.Newf("maxlength")
		}

		return nil
	}
}

func Length(minlength, maxlength int) StrFunc {
	return func(value string) error {
		l := len(value)
		if l < minlength || l > maxlength {
			return errgo.Newf("length")
		}

		return nil
	}
}

func RegExp(expression string) StrFunc {
	return func(value string) error {
		if matched, err := regexp.MatchString(expression, value); err != nil {
			return errgo.Mask(err)
		} else if !matched {
			return errgo.Newf("regexp")
		}

		return nil
	}
}

func Email() StrFunc {
	// Regexp comming from Angular.JS, to have the same validation
	// Backtick has to be escaped apart (+ "`" +)
	return RegExp(`(?i)^[a-z0-9!#$%&'*+\/=?^_` + "`" + `{|}~.-]+@[a-z0-9]([a-z0-9-]*[a-z0-9])?(\.[a-z0-9]([a-z0-9-]*[a-z0-9])?)*$`)
}

func In(allowed ...string) StrFunc {
	return func(value string) error {
		for _, item := range allowed {
			if item == value {
				return nil
			}
		}

		return errgo.Newf("in")
	}
}
