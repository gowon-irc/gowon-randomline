package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	_ = newHandler()
}

func TestHandlerGet(t *testing.T) {
	h := newHandler()
	out := h.Msg()

	assert.Equal(t, "handling", out)
}
