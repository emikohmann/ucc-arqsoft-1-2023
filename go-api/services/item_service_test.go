package services

import (
	"testing"
	"github.com/stretchr/testify/assert"
)
//
type TestClient struct {}

func (TestClient TestClient) GetItem(id int64) (ItemMercadoLibre, error){
	return ItemMercadoLibre{
		Title: "Mocked item",
		Price: 50.00,
	}, nil
}

func TestBuildItem(t *testing.T){

	itemML := ItemMercadoLibre{
		Title: "Test item",
		Price: 100.00,
	}

	// act
	item := buildItem(1, itemML)

	// assert
	// COMPARA QUE SEAN IGUALES
	assert.Equal(t, itemML.Title, item.Name)

	// PARA QUE NO SEA IGUAL A 0
	assert.NotZero(t, item.Price)


	
	//if item.Name != itemML.Title {

	//	t.Fail()


		// EJECUTAR ERROR CON MENSAJE EN LA CONSOLA
		//t.Error("te anda mal papi")
	//}

}

// PARA CORRER TEST "go test ./... -v"