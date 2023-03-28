package services

import "go-api/domain"

func GetItem(id int64) domain.Item {

	return domain.Item{
		ID:    id,
		Name:  "Hola",
		Price: 10000,
	}

}
