package main

import "github.com/gowon-irc/go-gowon"

type handler struct {
	lines []string
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
		},
	}
}

func (b *handlerBuilder) inputText(lines []string) *handlerBuilder {
	b.handler.lines = lines
	return b
}

func (b *handlerBuilder) build() handler {
	return b.handler
}
