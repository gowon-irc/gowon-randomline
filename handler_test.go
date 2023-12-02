package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBuildHandler(t *testing.T) {
	h, err := newBuilder().build()

	assert.Nil(t, err)

	out := h.Msg()

	assert.Equal(t, "", out)
}

func TestBuildHandlerWithInput(t *testing.T) {
	lines := []string{"handling input"}
	h, err := newBuilder().input(lines).build()

	assert.Nil(t, err)

	out := h.Msg()
	assert.Equal(t, "handling input", out)
}

func TestBuildHandlerWithShuffle(t *testing.T) {
	lines := []string{"a", "b", "c"}
	h, err := newBuilder().input(lines).setSeed(0).shuffle(true).build()

	assert.Nil(t, err)

	out := h.Msg()
	assert.Equal(t, "b", out)
}

func TestHandlerIncrement(t *testing.T) {
	tests := map[string]struct {
		lines    []string
		position int
		expected int
	}{
		"one line": {
			lines:    []string{"a"},
			position: 0,
			expected: 0,
		},
		"two lines": {
			lines:    []string{"a", "b"},
			position: 0,
			expected: 1,
		},
		"two lines loop": {
			lines:    []string{"a", "b"},
			position: 1,
			expected: 0,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			h, err := newBuilder().
				input(tc.lines).
				setPosition(tc.position).
				build()

			assert.Nil(t, err)

			h.increment()

			assert.Equal(t, tc.expected, h.position)
		})
	}
}

func TestBuildHandlerPositionError(t *testing.T) {
	tests := map[string]struct {
		position int
	}{
		"exceeds": {
			position: 1,
		},
		"below 0": {
			position: -1,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, err := newBuilder().setPosition(tc.position).build()

			assert.Error(t, err)
		})
	}
}
