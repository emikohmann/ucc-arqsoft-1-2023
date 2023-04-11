package services

import (
	"go-api/domain"
	"go-api/services/clients"
)

var (
	MLClient clients.MLClient
)

// GetItem returns the item information
func GetItem(id int64) (domain.Item, error) {
	itemML, err := MLClient.GetItem(id)
	if err != nil {
		return domain.Item{}, err
	}

	// Map MercadoLibre item to Item
	return buildItem(id, itemML), nil
}

// buildItem converts an MLItem into an Item
func buildItem(id int64, mlItem clients.MLItem) domain.Item {
	return domain.Item{
		ID:    id,
		Name:  mlItem.Title,
		Price: mlItem.Price,
	}
}
