# VK Callback API

VK API lib for Golang

Current VK API version: 5.80

Not everything is supported, see `TODO`s

Uses [easyjson](https://github.com/mailru/easyjson)

# Installation

`go generate` is required after installation

```
go get -v -u github.com/stek29/vk/... # would throw errors
go generate github.com/stek29/vk/...
go get github.com/stek29/vk/...
```

# Getting started

Minimal example -- making API calls:
```go
package main

import (
	"fmt"
	"os"
	"github.com/stek29/vk"
	"github.com/stek29/vk/vkapi"
)

func main() {
	client := vk.BaseAPIWithAccessToken(os.Getenv("VK_TOKEN"))
	users, _ := vkapi.Users{client}.Get(vkapi.UsersGetParams{
		UserIDs: []string{"1"},
		Fields:  []string{"followers_count"},
	})
	durov := users[0]
	fmt.Printf("Pavel Durov has %v followers\n", durov.FollowersCount)
}
```

For callback API server: See [this
gist](https://gist.github.com/stek29/7da818858713b7d82c1567800a478399)
-- server which deletes all comments containing cyrillic letters.
