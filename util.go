package vkCallbackApi

import (
	"net/url"
	"strconv"
	"strings"

	"github.com/mailru/easyjson"
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
