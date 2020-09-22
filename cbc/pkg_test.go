package cbc

import (
	"github.com/coreos/etcd/pkg/testutil"
	"github.com/panjf2000/ants/v2"
	"sync"
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

func BenchmarkDecrypt(b *testing.B) {
	b.StopTimer()

	keyCoder, _ := NewKeyCoder()

	toEncryptBytes := []byte("This is a string to encrypt")
	encryptedBytes, iv, _ := keyCoder.Encrypt(toEncryptBytes)

	wg := &sync.WaitGroup{}
	b.StartTimer()

	for t := 0; t <= 59; t++ {
		for i := 0; i < 13; i++ {
			wg.Add(1)
			go func() {
				decryptedBytes, _ := keyCoder.Decrypt(encryptedBytes, keyCoder.Key(), iv)
				_ = decryptedBytes
				wg.Done()
			}()
		}
		//time.Sleep(time.Second)
	}

	wg.Wait()
}

func BenchmarkPoolDecrypt(b *testing.B) {
	defer ants.Release()
	b.StopTimer()

	keyCoder, _ := NewKeyCoder()

	toEncryptBytes := []byte("This is a string to encrypt")
	encryptedBytes, iv, _ := keyCoder.Encrypt(toEncryptBytes)

	wg := sync.WaitGroup{}

	decryptFunc := func() {
		decryptedBytes, _ := keyCoder.Decrypt(encryptedBytes, keyCoder.Key(), iv)
		_ = decryptedBytes
	}

	p, _ := ants.NewPoolWithFunc(5, func(i interface{}) {
		decryptFunc()
		wg.Done()
	})
	defer p.Release()

	for t := 0; t <= 59; t++ {
		b.StartTimer()
		for i := 0; i < 13; i++ {
			wg.Add(1)
			_ = p.Invoke(i)
		}
		b.StopTimer()
		//time.Sleep(time.Second)
	}
	wg.Wait()
}
