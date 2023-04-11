package app

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPaths(t *testing.T) {
	assert.Equal(t, "/items/:itemID", pathGetItem)
}
