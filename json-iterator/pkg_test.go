// +build jsoniter

package json_iterator

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
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

func TestSetupRoter(t *testing.T) {
	router := SetupRoter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/jsoniter", nil)

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, fmt.Sprintln(`{"code":0,"data":{"id":1,"age":27,"gender":1,"name":"yuchanns","location":"China Guangdong Shenzhen Nanshan"}}`), w.Body.String())
}
