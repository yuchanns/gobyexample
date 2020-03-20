package err

import (
	"fmt"
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/pkg/errors"
	"testing"
)

func TestErrorStack(t *testing.T) {
	cause := ErrorStack()
	testutil.AssertNotNil(t, cause)
	err := errors.WithStack(cause)
	fmt.Printf("%+v", err)
}

func TestRuntimeCaller(t *testing.T) {
	err := RuntimeCaller()
	testutil.AssertNil(t, err)
}
