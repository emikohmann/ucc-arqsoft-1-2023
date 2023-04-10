package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPort(t *testing.T) {
	assert.Equal(t, ":8080", port)
}
