package main

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

func main() {
	token := pflag.String("token", "", "VK Token (required)")
	groupID := pflag.Int("group-id", 0, "Group ID (optional)")

	pflag.Parse()

	if *token == "" {
		log.Fatal("token is required")
	}

	baseAPI, err := vk.NewBaseAPI(vk.BaseAPIConfig{
		AccessToken: *token,
	})
	if err != nil {
		log.Fatal("Cant create baseAPI:", err)
	}

	bot, err := vkbot.NewBot(baseAPI, vkbot.BotConfig{
		GroupID: *groupID,
		Poller: &vkbot.LongPoller{
			Wait: 10 * time.Second,
		},
	})
	if err != nil {
		log.Fatal("Cant create bot:", err)
	}

	ctx, cancel := context.WithCancel(context.Background())

	go func(cancel context.CancelFunc) {
		sigint := make(chan os.Signal, 1)
		signal.Notify(sigint, os.Interrupt)
		<-sigint
		cancel()
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
