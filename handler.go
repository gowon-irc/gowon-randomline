package main

import "github.com/gowon-irc/go-gowon"

type handler struct {
	text string
}

func newHandler() *handler {
	return &handler{}
}

func (h *handler) Msg() string {
	return h.text
}

func (h *handler) handle(m gowon.Message) (string, error) {
	return h.Msg(), nil
}

type handlerBuilder struct {
	handler handler
}

func newBuilder() *handlerBuilder {
	return &handlerBuilder{
		handler: handler{},
	}
}

func (b *handlerBuilder) inputText(text string) *handlerBuilder {
	b.handler.text = text
	return b
}

func (b *handlerBuilder) build() handler {
	return b.handler
}
