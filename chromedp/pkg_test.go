package chromedp

import (
	"context"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestRender(t *testing.T) {
	ms := Mathmatics{
		{ID: "m1", Expr: "T_{pq}", Inline: true},
		{ID: "m2", Expr: "T^n", Inline: true},
		{ID: "m3", Expr: "a_{3}=b_{2}q+a_{2}q+a_{2}p", Inline: false},
		{ID: "m4", Expr: "\\sum\\limits_{n=a}^{b}f(n)=f(a)+\\cdots+f(b)", Inline: false},
	}
	mathematics, style, svg, err := Render(context.Background(), ms)
	assert.Nil(t, err)
	assert.NotEqual(t, "", style)
	assert.NotEqual(t, "", svg)
	assert.Equal(t, false, reflect.DeepEqual(ms, mathematics))
}
