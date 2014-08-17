package validations

import (
  "github.com/juju/errgo"
)

type Int64Func func(s int64) error

func Int64(field int64, validations []Int64Func) Func {
  return func() error {
    for _, validation := range validations {
      if err := validation(field); err != nil {
        return errgo.Mask(err)
      }
    }

    return nil
  }
}

func Range64(minvalue, maxvalue int64) Int64Func {
  return func(value int64) error {
    if value < minvalue || value > maxvalue {
      return errgo.Newf("range")
    }

    return nil
  }
}

func Positive64() Int64Func {
  return func(value int64) error {
    if value < 1 {
      return errgo.Newf("positive")
    }

    return nil
  }
}

func PositiveZero64() Int64Func {
  return func(value int64) error {
    if value < 0 {
      return errgo.Newf("positive zero")
    }

    return nil
  }
}

func MaxValue64(max int64) Int64Func {
  return func(value int64) error {
    if value > max {
      return errgo.Newf("max value")
    }

    return nil
  }
}
