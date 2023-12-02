package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	h := newHandler()

	assert.NotNil(t, h)
}

func TestHandlerInput(t *testing.T) {
	h := newHandler()

	h.Input("handling input")
	out := h.Msg()

	assert.Equal(t, "handling input", out)
}
