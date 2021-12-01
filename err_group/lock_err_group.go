package err_group

import (
	"strings"
	"sync"
)

type LockErrGroup struct {
	errs   []error
	lock   *sync.Mutex
	hasErr bool
}

func NewLockErrGroup() *LockErrGroup {
	return &LockErrGroup{
		errs:   []error{},
		lock:   &sync.Mutex{},
		hasErr: false,
	}
}

func (eg *LockErrGroup) Add(err error) {
	eg.lock.Lock()
	defer eg.lock.Unlock()

	eg.hasErr = true
	eg.errs = append(eg.errs, err)
}

func (eg *LockErrGroup) HasErr() bool {
	return eg.hasErr
}

func (eg *LockErrGroup) Error() string {
	errs := make([]string, 0, len(eg.errs))
	for _, err := range eg.errs {
		errs = append(errs, err.Error())
	}

	return strings.Join(errs, ";")
}
