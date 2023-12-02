package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewHandler(t *testing.T) {
	h := newHandler()

	assert.NotNil(t, h)
}

func TestHandlerBuilderInput(t *testing.T) {
	b := newBuilder().inputText("handling input").build()

	out := b.Msg()

	assert.Equal(t, "handling input", out)
}
