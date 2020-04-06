package cbc

import (
	"bytes"
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
)

type keyCoder struct {
	keyBytes []byte
}

func (k *keyCoder) Encrypt(data []byte) ([]byte, []byte, error) {
	block, err := des.NewTripleDESCipher(k.keyBytes)
	if err != nil {
		return nil, nil, err
	}

	bs := block.BlockSize()
	iv := make([]byte, bs)
	if _, err := rand.Read(iv); err != nil {
		return nil, nil, err
	}

	mode := cipher.NewCBCEncrypter(block, iv)

	data = zeroPadding(data, bs)

	encrypted := make([]byte, len(data))

	mode.CryptBlocks(encrypted, data)

	return encrypted, iv, nil
}

func (k *keyCoder) Key() []byte {
	return k.keyBytes
}

func (k *keyCoder) Decrypt(data, keyBytes, iv []byte) ([]byte, error) {
	block, err := des.NewTripleDESCipher(keyBytes)
	if err != nil {
		return nil, err
	}

	mode := cipher.NewCBCDecrypter(block, iv)

	decrypted := make([]byte, len(data))

	mode.CryptBlocks(decrypted, data)

	return zeroUnpadding(decrypted), nil
}

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
