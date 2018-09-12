package vkCallbackApi

import (
	"errors"
	"net/url"
	"strconv"
	"strings"

	"github.com/mailru/easyjson"
	"github.com/mailru/easyjson/jlexer"
)

// CSVStringSlice is a string slice which gets encoded
// as comma-separated string
// API methods often accept arrays of strings, which
// should be encoded as one comma-separated parameter
// This is helper type which implements query.Encoder
type CSVStringSlice []string

// EncodeValues conforms to query.Encoder inteface
func (csv CSVStringSlice) EncodeValues(key string, v *url.Values) error {
	if val := []string(csv); len(val) != 0 {
		encoded := strings.Join(val, ",")
		v.Set(key, encoded)
	}

	return nil
}

// CSVIntSlice is an int slice which gets encoded
// as comma-separated string
// API methods sometimes accept arrays of ints, which
// should be encoded as one comma-separated parameter
// This is helper type which implements query.Encoder
type CSVIntSlice []int

// EncodeValues conforms to query.Encoder inteface
func (csv CSVIntSlice) EncodeValues(key string, v *url.Values) error {
	strCSV := make(CSVStringSlice, len(csv))

	for i, v := range csv {
		strCSV[i] = strconv.Itoa(v)
	}

	return strCSV.EncodeValues(key, v)
}

func decodeBoolIntResponse(r easyjson.RawMessage) (bool, error) {
	resp, err := strconv.Atoi(string(r))
	return resp == 1, err
}

func urlValuesMerge(base, with url.Values) {
	for k, v := range with {
		if old, ok := base[k]; ok {
			base[k] = append(old, v...)
		} else {
			base[k] = v
		}
	}
}

// ArrayOnMeth is array which is represented as object by VK (see messages.delete)
type ArrayOnMeth []int

// UnmarshalJSON implements json.Unmarshaler interface
func (v *ArrayOnMeth) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	v.UnmarshalEasyJSON(&r)
	return r.Error()
}

// UnmarshalEasyJSON implements easyjson.Unmarshaler interface
func (v *ArrayOnMeth) UnmarshalEasyJSON(in *jlexer.Lexer) {
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeString()
		if id, err := strconv.Atoi(key); err == nil {
			*v = append(*v, id)
		}
		in.WantColon()
		in.SkipRecursive()
		in.WantComma()
	}
	in.Delim('}')
}

// genTODOType is placeholder type for unimplemented types in codegen
type genTODOType struct {
	fill bool // easyjson panics with zero division error on empty structs
}

// UnmarshalJSON implements json.Unmarshaler interface
func (genTODOType) UnmarshalJSON(data []byte) error {
	panic(errors.New("Trying to unmarshal genTODOType"))
}

// UnmarshalEasyJSON implements easyjson.Unmarshaler interface
func (genTODOType) UnmarshalEasyJSON(in *jlexer.Lexer) {
	panic(errors.New("Trying to unmarshal genTODOType"))
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
