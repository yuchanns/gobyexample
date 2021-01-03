package errors

import (
	"github.com/pkg/errors"
	"os"
	"testing"
)

func TestCheckAndOpen(t *testing.T) {
	errs := []struct {
		Err   error
		AsErr error
	}{
		{checkAndOpen(""), &emptyPathError{}},
		{checkAndOpen("none-existing"), &os.PathError{}},
	}

	for _, err := range errs {
		if !errors.As(err.Err, &err.AsErr) {
			t.Fatalf("err [%s] is not [%s]", err.Err, err.AsErr)
		}
	}
}
