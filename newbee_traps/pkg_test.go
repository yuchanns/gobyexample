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

func TestMultiDimension(t *testing.T) {
	table := MultiDimension()
	assert.Equal(t, [][]int{
		{0, 0, 0, 0},
		{0, 0, 0, 0},
	}, table)
}

func TestImmutableStrings(t *testing.T) {
	x, y := ImmutableStrings()
	assert.Equal(t, "Test", x)
	assert.Equal(t, "世界", y)
}

func TestValidateStringAndLength(t *testing.T) {
	result, length, cLength := ValidateStringAndLength("♥\xfe")
	assert.False(t, result)
	assert.Equal(t, 4, length)
	assert.Equal(t, 2, cLength)
}

func TestNilChannel(t *testing.T) {
	NilChannel()
}

func TestJsonUnmarshalNumberic(t *testing.T) {
	status1, status2, status3 := JsonUnmarshalNumberic()
	assert.IsType(t, uint64(1), status1)
	assert.IsType(t, int64(1), status2)
	assert.IsType(t, uint64(1), status3)
}
