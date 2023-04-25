package controllers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParams(t *testing.T) {
	assert.Equal(t, "itemID", paramItemID)
}
