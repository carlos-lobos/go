package main

import "fmt"

// Funciones que retornan multiples valores

func multiplicar(numero int) (n1, n2, n3 int) {
	n1 = numero * 10
	n2 = numero * 20
	n3 = numero * 30
	return
}

// Definicion alternativa a la primera
func multiplicar2(numero int) (int, int, int) {
	n1 := numero * 10
	n2 := numero * 20
	n3 := numero * 30
	return n1, n2, n3
}

// Ejemplo de que puede ser con cualquier tipo de valor
func retorno() (string, string) {
	return "Hola", "Mundo"
}

func main() {

	// Formas de usar esta caracteristica

	// 1, usandolo directamente
	fmt.Println(multiplicar(25))
	fmt.Println(multiplicar2(96))

	// 2, guardando cada valor en una variable de forma ordenada
	c1, c2, c3 := multiplicar(65)
	fmt.Println(c1, c2, c3)

	// 4, descartando algunos valores y quedandonos solamente con el que nos interesa
	_, d2, _ := multiplicar(66)
	fmt.Println(d2)

	// Ejemplo de que puede ser con cualquier tipo de valor
	fmt.Println(retorno())
}
