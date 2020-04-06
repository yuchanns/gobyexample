package ecb

import (
	"bytes"
	"crypto/des"
	"crypto/rand"
	"errors"
	"fmt"
)

type keyCoder struct {
	keyBytes []byte
}

func (k *keyCoder) Encrypt(data []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(k.keyBytes)
	if err != nil {
		return nil, err
	}

	bs := block.BlockSize()

	data = zeroPadding(data, bs)

	encrypted := make([]byte, len(data))
	dst := encrypted
	for len(data) > 0 {
		block.Encrypt(dst, data[:bs])
		data = data[bs:]
		dst = dst[bs:]
	}

	return encrypted, nil
}

func (k *keyCoder) Key() []byte {
	return k.keyBytes
}

func (k *keyCoder) Decrypt(data, keyBytes []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(keyBytes)
	if err != nil {
		return nil, err
	}

	bs := block.BlockSize()

	if len(data)%bs != 0 {
		return nil, errors.New(fmt.Sprintf("输入的数据长度不满足%d的倍数", bs))
	}
	decrypted := make([]byte, len(data))
	dst := decrypted
	for len(data) > 0 {
		block.Decrypt(dst, data)
		data = data[bs:]
		dst = dst[bs:]
	}

	return zeroUnpadding(decrypted), nil
}

// Deprecated: ecb encryption is unsafe. Do not use it unless necessary
func NewKeyCoder() (*keyCoder, error) {
	keyBytes := make([]byte, 24)
	if _, err := rand.Read(keyBytes); err != nil {
		return nil, err
	}
	return &keyCoder{keyBytes: keyBytes}, nil
}

func zeroPadding(data []byte, bs int) []byte {
	padding := bs - len(data)%bs
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(data, padtext...)
}

func zeroUnpadding(data []byte) []byte {
	return bytes.TrimFunc(data, func(r rune) bool {
		return r == rune(0)
	})
}
