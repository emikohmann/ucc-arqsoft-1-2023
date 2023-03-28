package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"go-api/domain"
	"io/ioutil"
	"net/http"
)

type ItemMercadoLibre struct {
	ID    string  `json:"id"`
	Title string  `json:"title"`
	Price float64 `json:"price"`
}

func GetItem(id int64) (domain.Item, error) {

	endpoint := fmt.Sprintf("/items/MLA%d", id)

	url := fmt.Sprintf("https://api.mercadolibre.com%s", endpoint)

	response, err := http.Get(url)
	if err != nil {
		return domain.Item{}, err
	}
	if response.StatusCode != http.StatusOK {
		return domain.Item{}, errors.New(
			fmt.Sprintf("unexpected status code: %d", response.StatusCode))
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return domain.Item{}, err
	}

	var itemML ItemMercadoLibre
	err = json.Unmarshal(bytes, &itemML)
	if err != nil {
		return domain.Item{}, err
	}

	return domain.Item{
		ID:    id,
		Name:  itemML.Title,
		Price: itemML.Price,
	}, nil
}
