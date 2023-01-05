package main

import "fmt"

// Variadic Functions

func sumar(numeros ...int) int {
	resultado := 0
	for _, numero := range numeros {
		resultado += numero
	}
	return resultado
}

// Solo el ultimo parametro puede ser indicado como variadic
// (esta funcion copia el comportamiento de fmt.Println() para concatenar strings)
func print(cadena string, cadenas ...string) {
	for _, c := range cadenas {
		cadena += " "
		cadena += c
	}
	fmt.Println(cadena)
}

func main() {

	fmt.Println(sumar(1, 2, 3, 4))
	fmt.Println(sumar(10, 65, 48, 75))
	fmt.Println(sumar(22, 25))
	fmt.Println(sumar(25))
	fmt.Println(sumar())

	// Puede enviarse un slice como argumento de un parametro variadic, agregando "..." pegado al parametro
	// esto hace que c/elemento del slice sea un parametro en la funcion
	numeros := []int{
		25,
		56,
		32,
	}
	fmt.Println(sumar(numeros...))

	print("Hola", "Mundo", "con", "Go")
}
