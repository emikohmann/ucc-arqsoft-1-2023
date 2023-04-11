package services

import (
	"github.com/stretchr/testify/assert"
	"go-api/services/clients"
	"testing"
)

func init() {
	// Here we define that for testing we will use MockClient as MLClient
	MLClient = clients.MockClient{}
}

func TestBuildItem(t *testing.T) {
	// prepare
	var itemID int64 = 1
	itemML := clients.MLItem{
		Title: "Test item",
		Price: 100.00,
	}

	// act
	item := buildItem(itemID, itemML)

	// assert
	assert.Equal(t, itemID, item.ID)
	assert.Equal(t, itemML.Title, item.Name)
	assert.Equal(t, itemML.Price, item.Price)
	assert.NotZero(t, item.Price)
}

func TestGetItem(t *testing.T) {
	//prepare
	item, err := GetItem(12345)

	//assert
	assert.Nil(t, err)
	assert.Equal(t, "Mocked item title", item.Name)
	assert.Equal(t, 50.00, item.Price)
}
