package tinytime

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	tt := New(1585750374)
	assert.Equal(t, TinyTime(1585750374), tt)
}

func TestUnix(t *testing.T) {
	tt := Unix(1, 0)
	assert.Equal(t, TinyTime(1), tt)

	tt = Unix(86400, 0)
	assert.Equal(t, TinyTime(86400), tt)

	tt = Unix(2678400, 0)
	assert.Equal(t, TinyTime(2678400), tt)

	tt = Unix(31536000, 0)
	assert.Equal(t, TinyTime(31536000), tt)
}
