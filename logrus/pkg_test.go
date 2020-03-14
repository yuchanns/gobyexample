package logrus

import (
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/sirupsen/logrus"
	"testing"
)

func TestInitLog(t *testing.T) {
	logger, err := InitLog()
	testutil.AssertNil(t, err)
	logger.WithFields(logrus.Fields{
		"name": "yuchanns",
		"age":  26,
		"kind": "gopher",
	}).Info()
}
