package main

import "fmt"

// Afuera de la funcion no se puede usar := para declarar variables
var razaGoku = "Saiyajin"

func main() {
	// Declaracion y uso de variable tipo entera
	var numero int
	numero = 25
	fmt.Println(numero)
	// (forma corta)
	numero2 := 40
	fmt.Println(numero2)

	// Declaracion y uso de variables tipo cadena
	var nombre string
	nombre = "Elizabet"
	// (forma corta)
	nombre2 := "Charly"
	fmt.Println(nombre + " - " + nombre2) // Concatenación

	// Asignación múltiple
	numero, nombre = 20, "Alberto"
	fmt.Println(numero)
	fmt.Println(nombre)

	// Asignacion multiple para realizar enroque de variables
	fmt.Println(nombre + " - " + nombre2)
	nombre, nombre2 = nombre2, nombre
	fmt.Println(nombre + " - " + nombre2)

	// Mix de declaración y asignación múltiple
	nombre3, nombre := "Miguel", "Angel"
	fmt.Println(nombre + " - " + nombre3)

	// Declaraciones explicitas poseen valores por defecto
	var num int    // == 0
	var nom string // == cadena vacia ""
	fmt.Println(num)
	fmt.Println(nom)

	// declarando multiples variables
	var (
		Dios               = "Goku"
		enemigo1, enemigo2 = "Babidi", "Cell"
	)
	var numero4 = 22
	fmt.Println(Dios, enemigo1, enemigo2, numero4)

	// Uso de variables globales
	fmt.Println("La raza de " + Dios + " es " + razaGoku)
}
