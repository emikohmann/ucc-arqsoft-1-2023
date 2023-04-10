package clients

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestURLs(t *testing.T) {
	assert.Equal(t, "/items/MLA%d", mercadoLibreItemEndpoint)
	assert.Equal(t, "https://api.mercadolibre.com%s", mercadoLibreBaseURL)
}
