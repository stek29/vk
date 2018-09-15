package vkbot

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strconv"
	"time"

	"github.com/stek29/vk"
	"github.com/stek29/vk/vkapi"
)

// LongPoller is a classic Bots Long Poll API based poller
type LongPoller struct {
	Wait time.Duration

	key    string
	server *url.URL
	ts     string
}

func (p *LongPoller) getServer(b *Bot) error {
	srv, err := vkapi.Groups{b}.GetLongPollServer(vkapi.GroupsGetLongPollServerParams{
		GroupID: b.GroupID,
	})
	if err != nil {
		return err
	}

	p.key = srv.Key
	p.ts = strconv.Itoa(srv.TS)

	if p.server, err = url.Parse(srv.Server); err != nil {
		return err
	}

	return nil
}

const (
	longPollErrorOk          int = 0
	longPollErrorNewTS           = 1
	longPollErrorKeyTooOld       = 2
	longPollErrorKeyTSTooOld     = 3
)

type longPollResponse struct {
	TS      string             `json:"ts"`
	Failed  int                `json:"failed"`
	Updates []vk.CallbackEvent `json:"updates"`
}

var errTryAgain = errors.New("vkbot/longpoll: try again")

func (p *LongPoller) getUpdates(ctx context.Context, b *Bot) ([]vk.CallbackEvent, error) {
	if p.server == nil || p.key == "" || p.ts == "" {
		oldTS := p.ts

		if err := p.getServer(b); err != nil {
			return nil, err
		}

		if oldTS != "" && p.ts != oldTS {
			p.ts = oldTS
		}
	}

	u := *p.server
	u.RawQuery = fmt.Sprintf("act=a_check&key=%v&ts=%v&wait=%v", p.key, p.ts, int(p.Wait/time.Second))

	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)

	r, err := b.HTTPClient().Do(req)
	if err != nil {
		return nil, err
	}
	defer r.Body.Close()

	resp := longPollResponse{}

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&resp); err != nil {
		return nil, err
	}

	switch resp.Failed {
	case longPollErrorOk:
		p.ts = resp.TS
		return resp.Updates, nil
	case longPollErrorNewTS:
		p.ts = resp.TS
		return nil, errTryAgain
	case longPollErrorKeyTooOld:
		p.key = ""
		return nil, errTryAgain
	case longPollErrorKeyTSTooOld:
		p.key = ""
		p.ts = ""
		return nil, errTryAgain
	default:
		return nil, fmt.Errorf("Longpoll: Unknown `failed` value %v", resp.Failed)
	}
}

// Poll conforms to Poller interface
func (p *LongPoller) Poll(ctx context.Context, b *Bot, dest chan<- vk.CallbackEvent) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			updates, err := p.getUpdates(ctx, b)
			if err == errTryAgain {
				continue
			}

			if err != nil {
				log.Printf("Error while trying to getUpdates: %v", err)
				continue
			}

			for _, upd := range updates {
				select {
				case <-ctx.Done():
					return
				case dest <- upd:
				}
			}
		}
	}
}
