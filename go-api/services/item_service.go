package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-api/domain"
	"io/ioutil"
	"net/http"
)

type Client interface {
	GetItem(id int64) (ItemMercadoLibre, error)
}

type ItemMercadoLibre struct {
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

const (
	mercadoLibreItemEndpoint = "/items/MLA%d"
	mercadoLibreBaseURL      = "https://api.mercadolibre.com%s"
)

var (
	ItemClient Client
)

type MlClient struct{}

func (client MlClient) GetItem(id int64) (ItemML ItemMercadoLibre, err error) {
	endpoint := fmt.Sprintf(mercadoLibreItemEndpoint, id)
	url := fmt.Sprintf(mercadoLibreBaseURL, endpoint)
	response, err := http.Get(url)
	if err != nil {
		return ItemMercadoLibre{}, err
	}

	if response.StatusCode != http.StatusOK {
		return ItemMercadoLibre{}, errors.New(fmt.Sprintf("unexpected status code %d", response.StatusCode))
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return ItemMercadoLibre{}, err
	}

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

	return buildItem(id, itemML), nil
}

func buildItem(id int64, itemML ItemMercadoLibre) domain.Item {
	return domain.Item{
		ID:    id,
		Name:  itemML.Title,
		Price: itemML.Price,
	}
}
