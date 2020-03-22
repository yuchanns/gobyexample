package errors

import (
	errors2 "errors"
	"fmt"
	"github.com/coreos/etcd/pkg/testutil"
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
	fmt.Printf("%+v\n", errors2.Unwrap(err5))
	fmt.Printf("%+v\n", errors.Cause(err5))
}

type whateverErr struct {
	msg string
}

func (e *whateverErr) Error() string {
	return e.msg

}

func TestPkgErrorAs(t *testing.T) {
	err1 := &whateverErr{msg: "error 1"}
	err2 := errors.WithStack(err1)
	var err3 *whateverErr
	testutil.AssertTrue(t, errors2.As(err2, &err3))
}
