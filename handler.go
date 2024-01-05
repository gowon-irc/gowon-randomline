package main

import (
	"errors"
	"math/rand"
	"time"

	"github.com/gowon-irc/go-gowon"
)

type handler struct {
	lines    []string
	seed     int64
	position int
	chance   int
	rand     *rand.Rand
}

func (h *handler) increment() {
	h.position = func() int {
		newpos := h.position + 1
		if newpos == len(h.lines) {
			return 0
		}
		return newpos
	}()
}

func (h *handler) Msg() string {
	chance := 100 - h.rand.Intn(100)

	if chance > h.chance {
		return ""
	}

	msg := h.lines[h.position]
	h.increment()
	return msg
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
			lines:    []string{""},
			seed:     time.Now().UnixNano(),
			position: 0,
			chance:   100,
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
			b.handler.lines[i], b.handler.lines[j] =
				b.handler.lines[j], b.handler.lines[i]
		})
	}

	return b
}

func (b *handlerBuilder) setPosition(position int) *handlerBuilder {
	b.handler.position = position
	return b
}

func (b *handlerBuilder) setChance(chance int) *handlerBuilder {
	b.handler.chance = chance
	return b
}

func (b *handlerBuilder) build() (handler, error) {
	if b.handler.position >= len(b.handler.lines) || b.handler.position < 0 {
		return b.handler, errors.New("position exceeds list length")
	}

	rs := rand.NewSource(b.handler.seed)
	b.handler.rand = rand.New(rs)

	return b.handler, nil
}
