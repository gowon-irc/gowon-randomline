package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildHandler(t *testing.T) {
	h := newBuilder().build()

	out := h.Msg()

	assert.Equal(t, "", out)
}

func TestBuildHandlerWithInput(t *testing.T) {
	lines := []string{"handling input"}
	b := newBuilder().inputText(lines).build()

	out := b.Msg()

	assert.Equal(t, "handling input", out)
}
