package vkbot

import (
	"context"

	"github.com/stek29/vk"
)

// <3 github.com/tucnak/telebot

// Poller is a provider of Events.
// All pollers must implement Poll(), which accepts bot
// pointer and subscription channel and starts polling
// *synchronously*.
type Poller interface {
	// Poll is supposed to take the bot object,
	// subscription channel and start polling
	// for Events immediately.
	// When ctx is Done, Poller should stop gracefully.
	// *dest channel might be closed when ctx is Done*
	Poll(ctx context.Context, b *Bot, dest chan<- vk.CallbackEvent)
}
