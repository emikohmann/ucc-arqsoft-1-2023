package main

import (
	"encoding/json"
	"fmt"
	"go-api/items"
	"errors"
)

func division(a int, b int) (int, error) {
	if b == 0 {
	return -1, errors.New("b no puede ser cero")
	}
	return a / b, nil
	}


func main() {


	fmt.Println("Hello, world")

	var a int = 10
	var b int = 0

	

	div, err := division(a, b)
	if err != nil {
	fmt.Println("Error:", err.Error())
	return
	}
	fmt.Println("Division:", div)


	// Puedo acceder a la constante StoreName de package items porque está exportada (empieza con mayúscula)
	fmt.Println(items.StoreName)

	// LLamo a GetItem() del package items y obtengo un *Item que lo asigno en la variable item
	item := items.GetItem()

	// Convierto el item a un []byte en formato JSON
	bytes, _ := json.MarshalIndent(item, "", "    ")

	// Convierto el []byte a string para mostrarlo por terminal
	jsonString := string(bytes)

	fmt.Println(jsonString)
}
