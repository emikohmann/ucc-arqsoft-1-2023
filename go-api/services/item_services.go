package services

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"go-api/domain")


	type ItemMercadoLibre struct{
		ID int64 `json: "id"`
		Name string `json: "name"`
		Price float64 `json: "price"`
	}

func GetItem(id int64) (domain.Item, error){

	endpoint:= fmt.Sprintf("/items/MLA%d", id)
	url := fmt.sprintf("https://api.mercadolibre.com%s", endpoint)

	response, err := http.Get(url)
	if err != nil{

		return domain.Item{}, err
	}

	if response.StatusCode != http.StatusOK{
		return domain.Item{}, errors.New(fmt.Sprintf("unexpected status code %d", response.StatusCode))
	}

	bytes, err := ioutil.ReadAll(response.Body)
	if err!=nil{
		return domain.Item{}, err
	}

	var itemML ItemMercadoLibre
	json.Unmarshal(bytes, &itemML)

	return domain.Item{
		ID:		id,
		Name:	itemML.Title,
		Price:	itemML.Price,
	}, nil


}

