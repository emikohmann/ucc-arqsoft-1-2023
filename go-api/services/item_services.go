package services

import (
<<<<<<< HEAD
=======
	"encoding/json"
>>>>>>> c0642656318807112e40c33dd6dd105d40bcd27c
	"errors"
	"fmt"
	"go-api/domain"
	"io/ioutil"
	"net/http"
)

<<<<<<< HEAD
func GetItem(id int64) (domain.Item, error) {

	endpoint := fmt.Sprintf("/items/MLA%d", id)

	url := fmt.Sprintf("http://api.mercadolibre.com%s", endpoint)

	response, err := http.Get(url)
	if err != nil {
		return domain.Item{}, err
	}
	if response.StatusCode != http.StatusOK {
		return domain.Item{}, errors.New(
			fmt.Sprintf("unexpected status code %d", response.StatusCode))
	}

	bytes, err := ioutil.ReadAll(response.Body)

	return domain.Item{
		ID:    id,
		Name:  "TV 32 pulgadas",
		Price: 100000,
	}, nil
=======
type Client interface {
	GetItem(id int64) (ItemMercadoLibre, error)
}

var (
	ItemClient Client
)

type ItemMercadoLibre struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

const (
	mercadoLibreItemEndpoint = "/items/MLA%d"
	mercadoLibreBaseURL      = "https://api.mercadolibre.com%s"
)

type MlClient struct{}

func (client MlClient) GetItem(id int64) (ItemMercadoLibre, error) {
	// Build the URL
	endpoint := fmt.Sprintf(mercadoLibreItemEndpoint, id)
	url := fmt.Sprintf(mercadoLibreBaseURL, endpoint)

	// Invoke MercadoLibre API
	response, err := http.Get(url)
	if err != nil {
		return ItemMercadoLibre{}, err
	}

	// Validate API Error
	if response.StatusCode != http.StatusOK {
		return ItemMercadoLibre{}, errors.New(fmt.Sprintf("unexpected status code %d", response.StatusCode))
	}

	// Read response payload bytes
	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ItemMercadoLibre{}, err
	}

	// Convert bytes to custom struct
	var itemML ItemMercadoLibre
	err = json.Unmarshal(bytes, &itemML)
	if err != nil {
		return ItemMercadoLibre{}, err
	}

	return itemML, nil
}

func GetItem(id int64) (domain.Item, error) {
	itemML, err := ItemClient.GetItem(id)
	if err != nil {
		return domain.Item{}, err
	}

	// Map MercadoLibre item to Item
	return buildItem(id, itemML), nil
}

func buildItem(id int64, itemML ItemMercadoLibre) domain.Item {
	return domain.Item{
		ID:    id,
		Name:  itemML.Title,
		Price: itemML.Price,
	}
>>>>>>> c0642656318807112e40c33dd6dd105d40bcd27c
}
