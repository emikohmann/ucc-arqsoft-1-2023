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

// function que se fija el precio de la orden de compra tira error si la cant <=0
func Buy(price float64, item *Item) (float64, error) {
	if price == 0 {
		return -1, errors.New("no se pueden pedir 0 items")
	}
	if price < 0 {
		return -1, errors.New("no se puede pedir items negativos")
	}
	return price * item.Price, nil
}
