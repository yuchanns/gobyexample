package wire

import (
	"bytes"
	"github.com/coreos/etcd/pkg/testutil"
	"io"
	"os"
	"testing"
)

func TestInitialEvent(t *testing.T) {
	old := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	event := InitialEvent()
	event.Start()

	outC := make(chan string)
	go func() {
		var buf bytes.Buffer
		_, err := io.Copy(&buf, r)
		testutil.AssertNil(t, err)
		outC <- buf.String()
	}()

	_ = w.Close()
	os.Stdout = old
	out := <-outC
	testutil.AssertEqual(t, out, "hello wire!\n")
}
