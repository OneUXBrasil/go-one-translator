package one_translator

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func Translate(content string, from string, to string) (string, error) {
	tt := ""
	content = url.PathEscape(content)

	url := fmt.Sprintf("http://translate.googleapis.com/translate_a/single?client=gtx&sl=%s&tl=%s&dt=t&q=%s&ie=UTF-8&oe=UTF-8", from, to, content)

	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	s, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return "", err
	}

	tt = string(s)
	sIndex := strings.Index(tt[4:], "\"")
	r := string(tt[4 : sIndex+4])

	return r, err
}
