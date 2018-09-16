package vkbot

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/stek29/vk"
)

// CallbackGroupConfig is a configuration for one group
//
// Secret and Confirmation should be set in group settings
type CallbackGroupConfig struct {
	GroupID      int
	Secret       string
	Confirmation string
}

// CallbackPoller is Callback API based poller
//
// If Listen is not empty, it starts an http server with that Addr
// Otherwise, it's up to caller to add CallbackPoller to http Mux
//
// GroupConfigs is slice of groups this poller should process events for
type CallbackPoller struct {
	Listen string
	// XXX: use map[int] instead of slice?
	GroupConfigs []CallbackGroupConfig

	dest chan<- vk.CallbackEvent
	ctx  context.Context
}

// ServeHTTP confroms to http.Handler interface
func (p *CallbackPoller) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	event := vk.CallbackEvent{}
	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&event); err != nil {
		log.Printf("Cant unmarshal event: %v", err)
		return
	}

	var foundCfg *CallbackGroupConfig
	for _, cfg := range p.GroupConfigs {
		if cfg.GroupID == event.GroupID {
			foundCfg = &cfg
			break
		}
	}

	if foundCfg == nil {
		log.Printf("There's no CallbackGroupConfig for Group %d, dropping", event.GroupID)
		return
	}

	if foundCfg.Secret != "" && foundCfg.Secret != event.Secret {
		log.Printf("Secret mismatch, dropping event for Group %v", event.GroupID)
		return
	}

	if _, ok := event.Event.(vk.Confirmation); ok {
		w.Write([]byte(foundCfg.Confirmation))
		return
	}

	go func() {
		select {
		case <-p.ctx.Done():
			log.Printf("Warning: Event would be lost because it was not processed before ctx.Done, but ok was already sent to VK")
		case p.dest <- event:
		}
	}()

	w.Write([]byte("ok\n"))
}

// Poll conforms to Poller interface
func (p *CallbackPoller) Poll(ctx context.Context, b *Bot, dest chan<- vk.CallbackEvent) {
	p.ctx = ctx
	p.dest = dest

	if p.Listen == "" {
		<-ctx.Done()
		return
	}

	srv := http.Server{
		Addr:    p.Listen,
		Handler: p,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Printf("Server unexpectedly stopped: %v", err)
		}
	}()

	<-ctx.Done()
	srv.Shutdown(ctx)
	return
}
