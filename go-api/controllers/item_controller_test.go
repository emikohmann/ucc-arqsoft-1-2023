package controllers

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParams(t *testing.T) {
	assert.Equal(t, "itemID", paramItemID)
}
