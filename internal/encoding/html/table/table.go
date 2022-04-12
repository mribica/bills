package table

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

type Decoder struct {
	input io.Reader
}

func NewDecoder(input io.Reader) *Decoder {
	return &Decoder{input: input}
}

func (d *Decoder) Decode(v *[]string) {

	tkn := html.NewTokenizer(d.input)

	for {
		tt := tkn.Next()
		if tt == html.ErrorToken {
			return
		}
		tag, _ := tkn.TagName()

		if string(tag) == "td" && tt == html.StartTagToken {
			*v = append(*v, getText(tkn))
		}
	}
}

func getText(tkn *html.Tokenizer) string {
	n := tkn.Next()
	if n == html.TextToken {
		return strings.TrimSpace(string(tkn.Text()))
	}
	return getText(tkn)
}
