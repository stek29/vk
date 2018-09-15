package vkbot_test

import (
	"context"
	"log"
	"os"
	"os/signal"
	"time"

	"github.com/spf13/pflag"
	"github.com/stek29/vk"
	"github.com/stek29/vk/vkapi"
	"github.com/stek29/vk/vkbot"
)

type godocDontStripMe struct{}

// Example_echobot is an echo bot which uses LongPoller
func Example_echobot() {
	_ = godocDontStripMe{}

	token := pflag.String("token", "", "VK Token (required)")
	groupID := pflag.Int("group-id", 0, "Group ID (optional)")

	pflag.Parse()

	if *token == "" {
		log.Fatal("token is required")
	}

	bot, err := vkbot.NewBot(vk.BaseAPIWithAccessToken(*token), vkbot.BotConfig{
		GroupID: *groupID,
		Poller: &vkbot.LongPoller{
			Wait: 10 * time.Second,
		},
	})
	if err != nil {
		log.Fatal("Cant create bot:", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	func(cancel context.CancelFunc) {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt)
		go func() {
			for range c {
				cancel()
			}
		}()
	}(cancel)

	events, err := bot.StartPolling(ctx, 0)
	if err != nil {
		log.Fatal("Cat start polling:", err)
	}

	for e := range events {
		// log.Printf("Got event: %+v", e)

		switch ev := e.Event.(type) {
		case vk.MessageNew:
			from := ev.PeerID
			text := ev.Text
			msgID := ev.ID

			log.Printf("New message(%v) from %v: `%v`", msgID, from, text)

			if text != "" {
				resp, err := vkapi.Messages{bot}.Send(vkapi.MessagesSendParams{
					PeerID:  from,
					Message: text,
					// ForwardMessages: ([]int{msgID}),
				})

				if err != nil {
					log.Printf("Cant send reply to (%v): %v", msgID, err)
				} else {
					log.Printf("Sent reply to (%v): reply id %v", msgID, resp)
				}
			}
		}
	}

	log.Printf("Bye!")
}
