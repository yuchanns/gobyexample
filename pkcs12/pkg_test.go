package pkcs12

import (
	"github.com/coreos/etcd/pkg/testutil"
	"testing"
)

func TestKeyGen(t *testing.T) {
	testutil.AssertNil(t, KeyGen())
}
