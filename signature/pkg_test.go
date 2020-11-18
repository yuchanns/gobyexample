package signature

import (
	"fmt"
	"testing"
)

func TestSum(t *testing.T) {
	result := Sum(1, 2)
	if result != 3 {
		t.Error("failed.")
	} else {
		fmt.Println(result)
	}
}
