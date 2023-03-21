package main

import (
	"fmt"
	"go-api/items"
)

func main() {
	div, err := items.Division(8, 4)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("division: ", div)
	}
		
}



