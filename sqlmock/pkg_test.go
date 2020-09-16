package sqlmock

import (
	"github.com/coreos/etcd/pkg/testutil"
	"testing"
)

func TestInsertOrUpdatePrintSql(t *testing.T) {
	err := InsertOrUpdatePrintSql()
	testutil.AssertNil(t, err)
}

func TestQueryRows(t *testing.T) {
	err := QueryRows()
	testutil.AssertNil(t, err)
}
