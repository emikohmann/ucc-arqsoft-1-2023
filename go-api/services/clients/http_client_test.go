package clients

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestURLs(t *testing.T) {
	assert.Equal(t, "/items/MLA%d", mercadoLibreItemEndpoint)
	assert.Equal(t, "https://api.mercadolibre.com%s", mercadoLibreBaseURL)
}
