package vk

import (
	"net/url"

	"github.com/google/go-querystring/query"
	"github.com/mailru/easyjson/jlexer"
)

// MergeURLValues merges mergeWith into base
//
// Can be useful when implementing API.Request
func MergeURLValues(base, mergeWith url.Values) {
	for k, v := range mergeWith {
		if old, ok := base[k]; ok {
			base[k] = append(old, v...)
		} else {
			base[k] = v
		}
	}
}

// BuildRequestParams is a helper function which can be used when implementing API.Request
//
// params can be:
// 	- nil
//  - url.Values
//  - url tagged struct (https://godoc.org/github.com/google/go-querystring/query)
func BuildRequestParams(params interface{}) (url.Values, error) {
	switch v := params.(type) {
	case nil:
		return make(url.Values), nil
	case url.Values:
		return v, nil
	default:
		return query.Values(params)
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
