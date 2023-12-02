package main

import "github.com/gowon-irc/go-gowon"

type handler struct{}

func newHandler() *handler {
	return &handler{}
}

func (h handler) Msg() string {
	return "handling"
}

func (h handler) Handle(m gowon.Message) (string, error) {
	return h.Msg(), nil
}
