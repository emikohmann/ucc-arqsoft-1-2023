package main

import (
	"encoding/json"
	"fmt"
	"go-api/items"
)

func main() {

	fmt.Println("Hello, world")

	// Puedo acceder a la constante StoreName de package items porque está exportada (empieza con mayúscula)
	fmt.Println(items.StoreName)

	// LLamo a GetItem() del package items y obtengo un *Item que lo asigno en la variable item
	item := items.GetItem()

	// Convierto el item a un []byte en formato JSON
	bytes, _ := json.MarshalIndent(item, "", "    ")

	// Convierto el []byte a string para mostrarlo por terminal
	jsonString := string(bytes)

	fmt.Println(jsonString)
	
	// Se crea logica de error de división por cero
	result, err := division(7, 1)
	if err != nil {
		fmt.Println("Error:", err.Error())
		os.Exit(1)
	} else {
		fmt.Println("El resultado de la división es: ", result)
	}
	
	//grupo13
}

// función división a prueba de error por división de cero
func division(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("No se puede dividir por cero")
	}
	return a / b, nil
}
