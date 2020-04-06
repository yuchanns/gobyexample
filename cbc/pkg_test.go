package cbc

import (
	"github.com/coreos/etcd/pkg/testutil"
	"testing"
)

func TestEncryptDecrypt(t *testing.T) {
	keyCoder, err := NewKeyCoder()
	testutil.AssertNil(t, err)

	toEncryptBytes := []byte("This is a string to encrypt")
	encryptedBytes, iv, err := keyCoder.Encrypt(toEncryptBytes)
	testutil.AssertNil(t, err)

	decryptedBytes, err := keyCoder.Decrypt(encryptedBytes, keyCoder.Key(), iv)

	testutil.AssertNil(t, err)
	testutil.AssertEqual(t, toEncryptBytes, decryptedBytes)
}
