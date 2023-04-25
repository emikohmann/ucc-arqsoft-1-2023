package items

import "errors"

type Item struct {
	Title       string
	Description string
	Price       float64
	HasTaxes    bool
}

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

func Division(a int, b int) (int, error) {
	if b == 0 {
		err := errors.New("b no puede ser cero")
		return -1, err
	}
	return a / b, nil
}
