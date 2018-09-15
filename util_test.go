package vk

import (
	"encoding/json"
	"net/url"
	"testing"
)

func TestMergeURLValues(t *testing.T) {
	first := url.Values{}
	second := url.Values{}

	first.Set("foo", "nameless")
	second.Set("bar", "king")

	MergeURLValues(first, second)

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

func assureBuildRequestParams(t *testing.T, p interface{}) url.Values {
	v, err := BuildRequestParams(p)
	if err != nil {
		t.Errorf("Unexpected error for BuildRequestParams(%v): %v", p, err)
	}
	return v
}

func TestBuildRequestParamsNil(t *testing.T) {
	nilValues := assureBuildRequestParams(t, nil)
	if len(nilValues) != 0 {
		t.Errorf("Expected empty map, got `%v` when building nil", nilValues)
	}
}

func TestBuildRequestParamsPrebuilt(t *testing.T) {
	const testKey = "foo"
	preBuilt := url.Values{}
	preBuilt.Set(testKey, "bar")
	postBuild := assureBuildRequestParams(t, preBuilt)

	expected := preBuilt.Get(testKey)
	actual := postBuild.Get(testKey)
	if actual != expected {
		t.Errorf("Expected `%v`, got `%v` for testKey", expected, actual)
	}
}

func TestBuildRequestParamsURLTagged(t *testing.T) {
	v := struct {
		TestField string `url:"test_field"`
	}{
		TestField: "ayylmao",
	}

	values := assureBuildRequestParams(t, v)
	testField := values.Get("test_field")

	if testField != v.TestField {
		t.Errorf("Expected `%v`, got `%v` for TestField", v.TestField, testField)
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
