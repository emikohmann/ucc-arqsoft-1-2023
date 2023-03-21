package items

import "fmt"
type Item struct {
	Title       string
	Description string
	Price       float64
	Grade int
	HasTaxes    bool
}

// HOLAAAA
// Constante exportada (empieza con mayúscula)
const StoreName = "Tienda Oficial Adidas"

// GetItem retorna un puntero a Item (*Item)
func GetItem() *Item {

	// Con & accedo a la dirección de memoria del Item, entonces equivale a un puntero
	return &Item{
		Title:       "Zapatillas Adidas",
		Description: "Descripcion Zapatillas Adidas Talle 40",
		Price:       30000,
		HasTaxes:    false,
	}
}

func ErrorTesting(valor int) ([]*Item, error, error) {

	var err error = nil
	var err1 error = nil
	if (valor <= 0){
		err= fmt.Errorf("La nota %d es menor a 1", valor)
	}
	if (valor >10){
		err1= fmt.Errorf("La nota %d es mayor que 10", valor)
	}
	notas := []*Item{}
	return notas, err, err1
}