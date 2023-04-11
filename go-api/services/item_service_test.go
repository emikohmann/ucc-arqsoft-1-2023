package services

import (
	"github.com/stretchr/testify/assert"
<<<<<<< HEAD
	"testing"
)

type TestClient struct{}

func (testClient TestClient) GetItem(id int64) (ItemMercadoLibre, error) {
	return ItemMercadoLibre{
		Title: "Mocked item",
		Price: 50.00,
	}, nil
=======
	"go-api/services/clients"
	"testing"
)

func init() {
	// Here we define that for testing we will use MockClient as MLClient
	MLClient = clients.MockClient{}
>>>>>>> master
}

func TestBuildItem(t *testing.T) {
	// prepare
<<<<<<< HEAD
	itemML := ItemMercadoLibre{
=======
	var itemID int64 = 1
	itemML := clients.MLItem{
>>>>>>> master
		Title: "Test item",
		Price: 100.00,
	}

	// act
<<<<<<< HEAD
	item := buildItem(1, itemML)

	// assert
	assert.Equal(t, itemML.Title, item.Name)
	assert.NotZero(t, item.Price)

	// if item.Name != itemML.Title {
	// 	t.Error("item name doesn't match")
	// }
}

func TestGetItem(t *testing.T) {

	item, err := GetItem(12345)
	assert.Nil(t, err)
	assert.Equal(t, "Mocked item", item.Name)

}
=======
	item := buildItem(itemID, itemML)

	// assert
	assert.Equal(t, itemID, item.ID)
	assert.Equal(t, itemML.Title, item.Name)
	assert.Equal(t, itemML.Price, item.Price)
	assert.NotZero(t, item.Price)
}

func TestGetItem(t *testing.T) {
	item, err := GetItem(12345)
	assert.Nil(t, err)
	assert.Equal(t, "Mocked item title", item.Name)
	assert.Equal(t, 50.00, item.Price)
}
>>>>>>> master
