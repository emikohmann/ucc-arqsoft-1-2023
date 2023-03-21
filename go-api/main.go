package main

import (
	"fmt"
	"go-api/items"
)

func main() {

	myitems, err := items.GetItems(-1)
	if err != nil {
		fmt.Println(err)
		return
	}
	println(myitems)

	/*

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
	*/

}
