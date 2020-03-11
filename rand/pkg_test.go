package rand

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/coreos/etcd/pkg/testutil"
	"math/big"
	"testing"
)

func TestCryptoRand(t *testing.T) {
	// rand int
	n, err := rand.Int(rand.Reader, big.NewInt(100))
	testutil.AssertNil(t, err)
	fmt.Println(n.Int64())
	// generate token
	tokenBytes := make([]byte, 32)
	_, err = rand.Read(tokenBytes)
	testutil.AssertNil(t, err)
	fmt.Println(base64.StdEncoding.EncodeToString(tokenBytes))
}
