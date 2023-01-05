package main

import "fmt"

func main() {

	//-- Estructuras de Control --//
	// Condicionales If

	a := 5
	if a < 5 {
		fmt.Println("a es menor que 5")
	} else if a > 5 {
		fmt.Println("a es mayor que 5")
	} else {
		fmt.Println("a es igual 5")
	}

	a = 5
	if a = a + 1; a < 5 {
		fmt.Println("a es menor que 5")
	} else if a > 5 {
		fmt.Println("a es mayor que 5")
	} else {
		fmt.Println("a es igual 5")
	}

	fmt.Println("Fin del programa")

}
