package validations

import (
	"github.com/juju/errgo"
)

// If only runs the validations if the condition it's true.
func If(condition bool, validations []Func) error {
	if !condition {
		return nil
	}

	for _, validation := range validations {
		if err := validation(); err != nil {
			return errgo.Mask(err)
		}
	}

	return nil
}
