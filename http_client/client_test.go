package httpclient

import (
	"fmt"
	"testing"
)

func TestSend(t *testing.T) {
	content, err := Send("posts/2019/12/16/web-structure-in-go.html", "yuchanns")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(content)
	}
}
