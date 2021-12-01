package err_group

import "strings"

type ChanErrGroup struct {
	errs   []error
	ch     chan error
	hasErr bool
}

func NewChanErrGroup() *ChanErrGroup {
	errGroup := &ChanErrGroup{
		errs:   []error{},
		ch:     make(chan error, 10),
		hasErr: false,
	}
	errGroup.consume()

	return errGroup
}

func (eg *ChanErrGroup) Add(err error) {
	eg.ch <- err
}

func (eg *ChanErrGroup) consume() {
	go func() {
		for {
			err := <-eg.ch
			eg.hasErr = true
			eg.errs = append(eg.errs, err)
		}
	}()
}

func (eg *ChanErrGroup) HasErr() bool {
	return eg.hasErr
}

func (eg *ChanErrGroup) Error() string {
	errs := make([]string, 0, len(eg.errs))
	for _, err := range eg.errs {
		errs = append(errs, err.Error())
	}

	return strings.Join(errs, ";")
}
