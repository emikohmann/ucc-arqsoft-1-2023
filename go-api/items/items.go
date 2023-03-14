package items

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
