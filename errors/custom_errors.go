package errors

import (
	"github.com/pkg/errors"
	"os"
)

type emptyPathError struct {
	err string
}

func (e *emptyPathError) Error() string {
	return e.err
}

func checkAndOpen(path string) error {
	if path == "" {
		return errors.Wrap(&emptyPathError{"it is an empty error"}, "it is a wrap error")
	}

	_, err := os.Open(path)

	return err
}
