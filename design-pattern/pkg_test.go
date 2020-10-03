package design_pattern

import (
	"github.com/coreos/etcd/pkg/testutil"
	"testing"
)

func TestBridge(t *testing.T) {
	brush := &brushPen{}
	pencil := &pencilPen{}
	blue := &blue{}
	red := &red{}

	testutil.AssertEqual(t, brush.Draw(red), "draw red color with brush")
	testutil.AssertEqual(t, pencil.Draw(blue), "draw blue color with pencil")
}
