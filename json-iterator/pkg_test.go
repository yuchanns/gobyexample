package json_iterator

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshal(t *testing.T) {
	_, err := Marshal()
	assert.Nil(t, err)
}

func BenchmarkMarshal(b *testing.B) {
	_, _ = Marshal()
}

func TestMonkeyPatch(t *testing.T) {
	_, err := MonkeyPatch()
	assert.Nil(t, err)
}

func BenchmarkMonkeyPatch(b *testing.B) {
	_, _ = MonkeyPatch()
}

func TestRegisterEncoder(t *testing.T) {
	_, err := RegisterEncoder()
	assert.Nil(t, err)
}

func BenchmarkRegisterEncoder(b *testing.B) {
	_, _ = RegisterEncoder()
}
