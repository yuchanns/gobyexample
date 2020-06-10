package errors

import (
	"errors"
	"github.com/coreos/etcd/pkg/testutil"
	perrors "github.com/pkg/errors"
	"os"
	"testing"
)

type emptyPathError struct {
	err string
}

func (e *emptyPathError) Error() string {
	return e.err
}

func checkAndOpen(path string) error {
	if path == "" {
		return perrors.Wrap(&emptyPathError{"it is an empty error"}, "it is a wrap error")
	}

	_, err := os.Open(path)

	return err
}

func TestCheckAndOpen(t *testing.T) {
	errs := []error{
		checkAndOpen(""),
		checkAndOpen("none-existing"),
	}

	for k, err := range errs {
		if k == 0 {
			var eErr *emptyPathError
			testutil.AssertTrue(t, errors.As(err, &eErr))
		} else {
			var oErr *os.PathError
			testutil.AssertTrue(t, errors.As(err, &oErr))
		}
	}
}
