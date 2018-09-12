package vkCallbackApi

import (
	"encoding/json"
	"net/url"
	"testing"
)

func TestStringSlice(t *testing.T) {
	cases := []struct {
		in  CSVStringSlice
		out string
	}{
		{CSVStringSlice{}, ""},
		{CSVStringSlice{"foo"}, "foo"},
		{CSVStringSlice{"foo", "bar"}, "foo,bar"},
	}

	for _, tcase := range cases {
		v := url.Values{}
		tcase.in.EncodeValues("key", &v)
		got := v.Get("key")

		if got != tcase.out {
			t.Errorf("Error while encoding %v: got %v, expected: %v.", tcase.in, got, tcase.out)
		}
	}
}

func TestIntSlice(t *testing.T) {
	cases := []struct {
		in  CSVIntSlice
		out string
	}{
		{CSVIntSlice{}, ""},
		{CSVIntSlice{1337}, "1337"},
		{CSVIntSlice{13, -37, 29}, "13,-37,29"},
	}

	for _, tcase := range cases {
		v := url.Values{}
		tcase.in.EncodeValues("key", &v)
		got := v.Get("key")

		if got != tcase.out {
			t.Errorf("Error while encoding %v: got %v, expected: %v.", tcase.in, got, tcase.out)
		}
	}
}

func TestBoolIntDecodeFunc(t *testing.T) {
	cases := []struct {
		in  string
		out bool
	}{
		{"0", false},
		{"1", true},
		{"-0", false},
		{"42", true},
	}

	for _, tcase := range cases {
		got, err := decodeBoolIntResponse([]byte(tcase.in))

		if err != nil {
			t.Errorf("Unexpected error while decoding %v: %v", tcase.in, err)
		}

		if got != tcase.out {
			t.Errorf("Error while decoding %v: got %v, expected: %v.", tcase.in, got, tcase.out)
		}
	}
}

func TestBoolIntType(t *testing.T) {
	cases := []struct {
		in  string
		out bool
	}{
		{"0", false},
		{"1", true},
		{"-0", false},
		{"42", true},
	}

	for _, tcase := range cases {
		var got BoolInt
		err := json.Unmarshal([]byte(tcase.in), &got)

		if err != nil {
			t.Errorf("Unexpected error while decoding %v: %v", tcase.in, err)
		}

		if bool(got) != tcase.out {
			t.Errorf("Error while decoding %v: got %v, expected: %v.", tcase.in, got, tcase.out)
		}
	}
}

func intSlicesEqual(l, r []int) bool {
	if len(l) != len(r) {
		return false
	}

	for i, v := range l {
		if v != r[i] {
			return false
		}
	}

	return true
}

func TestArrayOnMethUnmarshal(t *testing.T) {
	cases := []struct {
		in  string
		out []int
	}{
		{`{}`, []int{}},
		{`{"1": 1}`, []int{1}},
		{`{"5": 1, "6": 1, "8": 1}`, []int{5, 6, 8}},
	}

	for _, tcase := range cases {
		var got ArrayOnMeth

		err := json.Unmarshal([]byte(tcase.in), &got)

		if err != nil {
			t.Errorf("Unexpected error while unmarshaling %v: %v", tcase.in, err)
		}

		if !intSlicesEqual([]int(got), tcase.out) {
			t.Errorf("Error while decoding %v: got %v, expected: %v.", tcase.in, got, tcase.out)
		}
	}
}

func TestGenTODOType(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Unmarshaling genTODOType did not panic")
		}
	}()

	var v genTODOType
	json.Unmarshal([]byte(`"cafebabe"`), &v)
}
