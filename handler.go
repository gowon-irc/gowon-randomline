package main

import "github.com/gowon-irc/go-gowon"

type handler struct {
	text string
}

func newHandler() *handler {
	return &handler{}
}

func (h *handler) Input(s string) {
	h.text = s
}

func (h *handler) Msg() string {
	return h.text
}

func (h *handler) Handle(m gowon.Message) (string, error) {
	return h.Msg(), nil
}
