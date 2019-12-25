package rsa

import (
	"fmt"
	"testing"
)

func TestEncrypt(t *testing.T) {
	data := map[string]interface{}{
		"name":   "yuchanns",
		"age":    21,
		"gender": "male",
	}

	encrypted, err := Encrypt(&data)

	if err != nil {
		fmt.Println("error:", err.Error())
	} else {
		fmt.Println("encrypted:", encrypted)
	}
}

func BenchmarkEncrypt(b *testing.B) {
	data := map[string]interface{}{
		"name":   "yuchanns",
		"age":    21,
		"gender": "male",
	}

	encrypted, err := Encrypt(&data)

	if err != nil {
		fmt.Println("error:", err.Error())
	} else {
		fmt.Println("encrypted:", encrypted)
	}
}
