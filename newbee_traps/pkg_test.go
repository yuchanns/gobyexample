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

func TestJsonEncoderAddNewline(t *testing.T) {
	byteString, rawString := JsonEncoderAddNewline()
	assert.NotEqual(t, byteString, rawString)
}

func TestJsonEscapeHTML(t *testing.T) {
	rawString, b1String, b2String := JsonEscapeHTML()
	assert.Equal(t, "\"x \\u003c y\"", rawString)
	assert.Equal(t, "\"x \\u003c y\"\n", b1String)
	assert.Equal(t, "\"x < y\"\n", b2String)
}

func TestJsonUnmarshalNumberic(t *testing.T) {
	var data = []byte(`{"status": 200}`)
	status1, status2, status3 := JsonUnmarshalNumberic(data)
	assert.Equal(t, uint64(200), status1)
	assert.Equal(t, int64(200), status2)
	assert.Equal(t, uint64(200), status3)
}

func TestJsonUnmarshalUncertainType(t *testing.T) {
	records := [][]byte{
		[]byte(`{"status": 200, "tag": "one"}`),
		[]byte(`{"status": "ok", "tag": "two"}`),
	}
	JsonUnmarshalUncertainType(records)
}

func TestHiddenDataInSlice(t *testing.T) {
	raw, rawNew, rawCopy, rawFull := HiddenCapacityInSlice()
	assert.Equal(t, cap(raw), cap(rawNew))
	assert.Equal(t, 3, cap(rawCopy))
	assert.Equal(t, 3, cap(rawFull))
}

func TestDeferredExcutionTime(t *testing.T) {
	DeferredExcutionTime()
}

func TestNilInterface(t *testing.T) {
	data, in, in2 := NilInterface()
	assert.True(t, data == nil)
	assert.True(t, in == nil)
	assert.False(t, in2 == nil)
}
