package vk

import (
	"encoding/json"
	"testing"
)

func TestAttachmentUnmarshal(t *testing.T) {
	const url = "http://example.com"
	const title = "Example"

	var test Attachment
	json.Unmarshal([]byte(`
		{
			"link": {
				"url": "`+url+`",
				"title": "`+title+`"
			}
		}
	`), &test)

	var link Link

	if v, ok := test.Val.(Link); ok {
		link = v
	} else {
		t.Errorf("Expected to find Link after unmarshal, found %T", test)
	}

	if link.URL != url {
		t.Errorf("Wrong URL: Expected %v, got %v", url, link.URL)
	}

	if link.Title != title {
		t.Errorf("Wrong Title: Expected %v, got %v", title, link.Title)
	}
}
