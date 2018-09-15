package vkbot

import (
	"context"
	"errors"
	"strconv"

	"github.com/stek29/vk"
	"github.com/stek29/vk/vkapi"
)

// BotConfig represents configuration used for Bot creation
type BotConfig struct {
	// Poller to be used when StartPolling is called
	Poller Poller
	// GroupID this bot is running as -- optional if group access token is used
	GroupID int
}

// Bot represents VK Bot instance
//
// Conforms to vk.API interface and can be used in vkapi
type Bot struct {
	vk.API
	BotConfig

	me *vk.Group
}

// NewBot tries to instantiate a bot which uses baseAPI for API requests
func NewBot(baseAPI vk.API, cfg BotConfig) (*Bot, error) {
	b := &Bot{
		API:       baseAPI,
		BotConfig: cfg,
	}

	if _, err := b.GetMe(true); err != nil {
		return nil, err
	}

	return b, nil
}

// GetMe returns VK Group this bot is running as
//
// Result is cached, pass flush=true to force new request
func (b *Bot) GetMe(flush bool) (*vk.Group, error) {
	if b.me != nil && !flush {
		return b.me, nil
	}

	groups, err := vkapi.Groups{b}.GetByID(vkapi.GroupsGetByIDParams{
		GroupID: strconv.Itoa(b.GroupID),
	})

	if err != nil {
		return nil, err
	}

	if len(groups) != 1 {
		return nil, errors.New("VK did not return group we needed")
	}

	b.me = &groups[0]
	b.GroupID = b.me.ID

	return b.me, nil
}

// StartPolling starts polling for events in background
//
// Returns channel to read events from, which is closed when ctx is Done
//
// Usage:
//
//   events, _ := b.StartPolling(ctx, 0)
//   for event := range events {
//   	// handle event here
//   }
//   // ctx is Done
func (b *Bot) StartPolling(ctx context.Context, Cap int) (<-chan vk.CallbackEvent, error) {
	if b.Poller == nil {
		return nil, errors.New("Poller is required")
	}

	// stopPoll := make(chan struct{})
	events := make(chan vk.CallbackEvent, Cap)

	go func() {
		<-ctx.Done()
		close(events)
	}()

	go b.Poller.Poll(ctx, b, events)

	return events, nil
}
