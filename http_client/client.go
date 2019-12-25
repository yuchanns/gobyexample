package httpclient

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	baseUri = "https://www.yuchanns.xyz"
)

func Send(target string, data string) (content string, err error) {
	var resp *http.Response
	resp, err = http.PostForm(strings.Join([]string{
		baseUri,
		target,
	}, "/"), url.Values{"name": {data}})
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)

	if err == nil {
		content = bytes.NewBuffer(body).String()
	}

	return
}
