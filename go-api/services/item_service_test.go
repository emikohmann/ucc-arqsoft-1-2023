package services

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBuildItem(t *testing.T) {
	//prepare
	itemML := ItemMercadoLibre{
		Title: "Test Item",
		Price: 100.00,
	}

	//act
	item := buildItem(1, itemML)

	//assert
	assert.Equal(t, itemML.Title, item.Name)
	assert.NotZero(t, item.Price)
}

type TestClient struct {
}

func (testClient TestClient) GetItem(id int64) (ItemMercadoLibre, error) {
	return ItemMercadoLibre{
		Title: "Mocked item",
		Price: 50.00,
	}, nil
}

func TestGetItem(t *testing.T) {
	//prepare
	var id int64 = 1150662888
	ItemClient = TestClient{}

	//act
	item, _ := GetItem(id)

	//assert
	assert.Equal(t, item.ID, id)
	assert.Equal(t, item.Name, "Mocked item")
}
