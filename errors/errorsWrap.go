package err

import (
	"fmt"
	"github.com/pkg/errors"
	"runtime"
)

func ErrorStack() error {
	err := errors.New("error happens")
	return err
}

func RuntimeCaller() error {
	pc, file, line, ok := runtime.Caller(0)
	if !ok {
		return errors.New("runtime.Caller not ok")
	}
	fmt.Println("function name:", runtime.FuncForPC(pc).Name())
	fmt.Println("file name:", file)
	fmt.Println("line number:", line)

	return nil
}
