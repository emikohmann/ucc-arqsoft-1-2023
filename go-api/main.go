package main

import (
	"go-api/app"
	"go-api/items"

	"encoding/json"
	"fmt"
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

	//funcion para checkear precio de orden de compra devuelve error si la cant a comprar es 0 o negativa

	fmt.Print(items.Buy(-4, item))

	// grupo-15 //
	app.StartApp()
}

// Hola
