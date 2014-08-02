package validations

import (
	"github.com/juju/errgo"
)

type Func func() error

func Run(validations []Func) error {
	for _, validation := range validations {
		if err := validation(); err != nil {
			return errgo.Mask(err)
		}
	}

	return nil
}
