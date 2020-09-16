package beego

import (
	"github.com/coreos/etcd/pkg/testutil"
	"testing"
)

func TestInsertOrUpdatePrintSql(t *testing.T) {
	err := InsertOrUpdatePrintSql()
	testutil.AssertNil(t, err)
}
