validations
===========

Go validations

Install
-------

```shell
go get github.com/ernestoalejo/validations
```

API reference
-------------

http://godoc.og/github.com/ernestoalejo/validations


Examples
--------

#### Validate strings, emails and integers

```go
package main

import (
  vals "github.com/ernestoalejo/validations"
)

type requestData struct {
  Foo string
  FooEmail string
  Bar int
}

func (req *requestData) Validate() error {
  return vals.Run([]vals.Func{
    vals.Str(req.Foo, []vals.StrFunc{
      vals.Required(),
    }),

    vals.Str(req.FooEmail, []vals.StrFunc{
      vals.Required(),
      vals.Email(),
    }),

    vals.Int(req.Bar, []vals.IntFunc{
      vals.Range(10, 20),
    }),
  })
}
```


#### Validate slices

```go
package main

import (
  vals "github.com/ernestoalejo/validations"
)

type requestData struct {
  Foo []*sliceData
}

type sliceData struct {
  Foo string
}

func (req *requestData) Validate() error {
  return vals.Run([]vals.Func{
    vals.Slice(req.Rooms, []vals.SliceFunc{
      vals.SliceMinLen(1),

      vals.Each(func(item *sliceData) error {
        return vals.Run([]vals.Func{
          vals.Str(item.Foo, []vals.StrFunc{
            vals.Required(),
          }),
        })
      }),
    }),
  })
}
```

#### Validate times

```go
package main

import (
  "time"

  vals "github.com/ernestoalejo/validations"
)

type requestData struct {
  EntryRaw, DepartureRaw string

  entry, departure time.Time
}

func (req *requestData) Validate() error {
  return vals.Run([]vals.Func{
    vals.Time(&req.entry, []vals.TimeFunc{
      vals.ParseTime(req.EntryRaw, "2006-01-02"),
      vals.NotBefore(time.Now()),
    }),

    vals.Time(&req.departure, []vals.TimeFunc{
      vals.ParseTime(req.DepartureRaw, "2006-01-02"),
      vals.NotBefore(req.entry),
    }),
  })
}
```
