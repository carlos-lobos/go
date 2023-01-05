package main

import "fmt"

func main() {

	// Dato de tipo Bool (posibles valores: true, false)
	var resultado bool

	resultado = 4 < 7
	fmt.Println(" 4 < 7 - ", resultado)
	fmt.Println("***************************************")
	resultado = (4 < 7) && (2 > 1)
	fmt.Println(" (4 < 7) && (2 > 1) - ", resultado)
	fmt.Println("***************************************")
	resultado = (4 > 7) || (2 > 1)
	fmt.Println(" (4 > 7) || (2 > 1) - ", resultado)
	fmt.Println("***************************************")
	fmt.Println("!resultado", !resultado)

}
