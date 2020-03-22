package errors

import (
	errors2 "errors"
	"fmt"
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/pkg/errors"
	"os"
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

func TestPkgErrorAsIs(t *testing.T) {
	err := virtualErr()
	var err2 *os.PathError
	testutil.AssertTrue(t, errors2.As(err, &err2))
	fmt.Println("file not found")
	fmt.Printf("%+v\n", err)
	err3 := virtualErr2()
	var err4 *whateverErr
	testutil.AssertTrue(t, errors2.As(err3, &err4))
	fmt.Println("display a cutom method")
	fmt.Println(err4.CustomError())
}

type whateverErr struct {
	msg string
}

func (w *whateverErr) Error() string {
	return w.msg
}

func (w *whateverErr) CustomError() string {
	return "a text that tells file not found~"
}

func virtualErr() error {
	if _, err := os.Open("non-existing"); err != nil {
		return errors.WithStack(err)
	}
	return errors.WithStack(&whateverErr{msg: "this is a whatever error"})
}

func virtualErr2() error {
	_, err := os.Open("non-existing")
	if err != nil {
		return errors.WithStack(&whateverErr{msg: "this is a whatever error"})
	}
	return errors.WithStack(err)
}
