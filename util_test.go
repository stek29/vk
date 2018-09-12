package vk

import (
	"encoding/json"
	"net/url"
	"testing"
)

func TestURLValuesMerge(t *testing.T) {
	first := url.Values{}
	second := url.Values{}

	first.Set("foo", "nameless")
	second.Set("bar", "king")

	urlValuesMerge(first, second)

	if first.Get("foo") != "nameless" {
		t.Errorf("Key foo disappeared from first after merge")
	}
	if second.Get("bar") != "king" {
		t.Errorf("Key bar disappeared from second after merge")
	}
	if first.Get("bar") != "king" {
		t.Errorf("Key bar did not appear in first after merge")
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
