package main

import (
	"math/rand"
	"time"

	"github.com/gowon-irc/go-gowon"
)

type handler struct {
	lines []string
	seed  int64
}

func (h *handler) Msg() string {
	return h.lines[0]
}

func (h *handler) handle(m gowon.Message) (string, error) {
	return h.Msg(), nil
}

type handlerBuilder struct {
	handler handler
}

func newBuilder() *handlerBuilder {
	return &handlerBuilder{
		handler: handler{
			lines: []string{""},
			seed:  time.Now().UnixNano(),
		},
	}
}

func (b *handlerBuilder) input(lines []string) *handlerBuilder {
	b.handler.lines = lines
	return b
}

func (b *handlerBuilder) setSeed(seed int64) *handlerBuilder {
	b.handler.seed = seed
	return b
}

func (b *handlerBuilder) shuffle(shuffle bool) *handlerBuilder {
	if shuffle {
		rs := rand.NewSource(b.handler.seed)
		r := rand.New(rs)
		r.Shuffle(len(b.handler.lines), func(i, j int) {
			b.handler.lines[i], b.handler.lines[j] = b.handler.lines[j], b.handler.lines[i]
		})
	}

	return b
}

func (b *handlerBuilder) build() handler {
	return b.handler
}
