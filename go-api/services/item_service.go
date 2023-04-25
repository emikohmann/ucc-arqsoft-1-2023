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

	return buildItem(id, itemML), nil
}

// buildItem converts an MLItem into an Item
func buildItem(id int64, mlItem clients.MLItem) domain.Item {
	var url string
	if len(mlItem.Pictures) > 0 {
		url = mlItem.Pictures[0].URL
	}
	return domain.Item{
		ID:     id,
		Name:   mlItem.Title,
		Price:  mlItem.Price,
		ImgURL: url,
	}
}
