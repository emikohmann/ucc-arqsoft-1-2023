package services

import (
	"errors"
	"fmt"
	"go-api/domain"
	"io/ioutil"
	"net/http"
)

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
}
