# VK API

[![GoDoc](https://godoc.org/github.com/stek29/vk?status.svg)](https://godoc.org/github.com/stek29/vk)

VK API lib for Golang

Current VK API version: 5.92

Almost everything is supported, but not everything is tested (VK API Schema has a lot of issues).

See [`TODO`s](https://github.com/stek29/vk/search?q=TODO) and [unsupported types](https://github.com/stek29/vk/search?q=genTODOType).

Uses [easyjson](https://github.com/mailru/easyjson)
Inspired by [telebot](https://github.com/tucnak/telebot)

# Overview

Library consists of three packages:
- vk: Core package, defines API interface, provides BaseAPI
	implementation and defines most of types used by VK API
- vkapi: Automatically generated wrappers for API
- vkbot: Various helpers for making VK Bots -- using Callback API or
	Bots Long Poll API to automate communities

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
	client, _ := vk.NewBaseAPI(vk.BaseAPIConfig{AccessToken: os.Getenv("VK_TOKEN")})
	users, _ := vkapi.Users{client}.Get(vkapi.UsersGetParams{
		UserIDs: []string{"1"},
		Fields:  []string{"followers_count"},
	})
	fmt.Printf("Pavel Durov has %v followers\n", users[0].FollowersCount)
}
```

For bot example: See [echobot](examples/echobot)

Also see [nocyril](examples/nocyril): A bit more advanced "bot" which supports multiple groups and works via callback poller.

## Uploading files
Currently there's no wrapper to handle file uploads, since it's trivial to do it manually.

See [vidloader](examples/vidloader) for an example of video uploader program.
