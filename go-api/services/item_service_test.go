package services

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestClient struct{}

func (testClient TestClient) GetItem(id int64) (ItemMercadoLibre, error) {
	return ItemMercadoLibre{
		Title: "Mocked item",
		Price: 50.00,
	}, nil
}

func TestBuildItem(t *testing.T) {
	// prepare
	itemML := ItemMercadoLibre{
		Title: "Test item",
		Price: 100.00,
	}

	// act
	item := buildItem(1, itemML)

	// assert
	assert.Equal(t, itemML.Title, item.Name)
	assert.NotZero(t, item.Price)

	// if item.Name != itemML.Title {
	// 	t.Error("item name doesn't match")
	// }
}

// IMPLEMENTAR EL TEST PARA GetItem (tarea clase)
func TestGetItem(t *testing.T) {
	//prepare
	var id int64 = 1356058509
	ItemClient = TestClient{}
	//act
	item, err := GetItem(id)
	if err != nil {
		println("ERROR")
	}

	//assert
	assert.Equal(t, item.ID, id)
}
