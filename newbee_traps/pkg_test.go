package newbee_traps

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNilInitVariableWithExplicitType(t *testing.T) {
	x := NilInitVariableWithExplicitType()
	assert.Nil(t, x)
}

func TestNilInitSlicesAndMaps(t *testing.T) {
	m, s := NilInitSlicesAndMaps()
	assert.Equal(t, map[string]int{"one": 1}, m)
	assert.Equal(t, []int{1}, s)
}

func TestInitStrings(t *testing.T) {
	s := InitStrings()
	assert.NotEmpty(t, s)
}

func TestRangeSlices(t *testing.T) {
	RangeSlices()
}
