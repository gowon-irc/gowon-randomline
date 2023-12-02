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
	b := newBuilder().input(lines).build()

	out := b.Msg()
	assert.Equal(t, "handling input", out)
}

func TestBuildHandlerWithShuffle(t *testing.T) {
	lines := []string{"a", "b", "c"}
	b := newBuilder().input(lines).setSeed(0).shuffle(true).build()

	out := b.Msg()
	assert.Equal(t, "b", out)
}
