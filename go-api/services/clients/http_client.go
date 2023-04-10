package clients

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

// HTTPClient is a particular implementation of MLClient (interface)
// And it works calling the Mercado Libre API via HTTP
type HTTPClient struct{}

const (
	mercadoLibreItemEndpoint = "/items/MLA%d"
	mercadoLibreBaseURL      = "https://api.mercadolibre.com%s"
)

func (client HTTPClient) GetItem(id int64) (MLItem, error) {
	// Build the URL
	endpoint := fmt.Sprintf(mercadoLibreItemEndpoint, id)
	url := fmt.Sprintf(mercadoLibreBaseURL, endpoint)

	// Invoke MercadoLibre API
	response, err := http.Get(url)
	if err != nil {
		return MLItem{}, err
	}

	// Validate API Error
	if response.StatusCode != http.StatusOK {
		return MLItem{}, errors.New(fmt.Sprintf("unexpected status code %d", response.StatusCode))
	}

	// Read response payload bytes
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return MLItem{}, err
	}

	// Convert bytes to custom struct
	var itemML MLItem
	err = json.Unmarshal(bytes, &itemML)
	if err != nil {
		return MLItem{}, err
	}

	return itemML, nil
}
