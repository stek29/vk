package vk

import (
	"net/url"

	"github.com/mailru/easyjson/jlexer"
)

// urlValuesMerge merges with into base
func urlValuesMerge(base, with url.Values) {
	for k, v := range with {
		if old, ok := base[k]; ok {
			base[k] = append(old, v...)
		} else {
			base[k] = v
		}
	}
}

// BoolInt is bool type which conforms to easyjson.Unmarshaler interface
// and unmarshals from VK's favorite 1/0 int bools
type BoolInt bool

// UnmarshalJSON implements json.Unmarshaler interface
func (v *BoolInt) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&r)
	return r.Error()
}

// UnmarshalEasyJSON implements easyjson.Unmarshaler interface
func (v *BoolInt) UnmarshalEasyJSON(in *jlexer.Lexer) {
	*v = in.Int() != 0
}
