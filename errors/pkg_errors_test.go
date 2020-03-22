package errors

import (
	errors2 "errors"
	"fmt"
	"github.com/pkg/errors"
	"testing"
)

func TestPkgErrors(t *testing.T) {
	err := errors.New("error a")
	fmt.Printf("%+v", err)
}

func TestPkgWithStack(t *testing.T) {
	err := errors2.New("error a")
	err2 := errors.WithStack(err)
	fmt.Printf("%+v\n", err2)
	err3 := errors.WithMessage(err, "error b")
	fmt.Printf("%+v\n", err3)
	err4 := errors.Cause(err3)
	fmt.Printf("%+v\n", err4)
	err5 := errors.Wrap(err2, "error c")
	fmt.Printf("%+v", err5)
}
