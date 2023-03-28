package services

import (
	"go-api/domain"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"encoding/json"
)

// JSON - TO GO para transformar estructuras json en structuras .go

type ItemMercadoLibre struct {
	Title 	string 	`json:"title"`
	Price	float64 `json:"price"`

}


func GetItem(id int64) (domain.Item, error) {

	endpoint := fmt.Sprintf("/items/MLA%d", id)

	url := fmt.Sprintf("https://api.mercadolibre.com%s", endpoint)

	response, err := http.Get(url)
	if err != nil {
		
		return domain.Item{}, err
	}
	if response.StatusCode != http.StatusOK {
		return domain.Item{}, errors.New(fmt.Sprintf("unexpected status code %d", response.StatusCode))
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		
		return domain.Item{}, err
	}

	var itemML ItemMercadoLibre
	json.Unmarshal(bytes, &itemML)

	return domain.Item{
		ID: id,
		Name: itemML.Title,
		Price: itemML.Price,
	}, nil

}