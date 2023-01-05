package main

import "fmt"

// Type

// Declarar nuestro propio tipo

// Dinero Tipo declarado por nosotros, en base al tipo int
type Dinero int

// Declaramos el metodo String() para el tipo Dinero
// (pisamos el metodo String() por defecto de int por el nuestro)
func (d Dinero) String() string {
	return fmt.Sprintf("$%d", d)
}

func main() {
	var sueldo Dinero
	sueldo = 25000
	// Println() al querer imprimir la variable sueldo, llama implicitamente al metodo String()
	fmt.Println("Sueldo: ", sueldo)

	aumento := 10000
	// Para sumar aumento al sueldo, hay que convertirlo al tipo Dinero
	sueldo += Dinero(aumento)

	fmt.Println("Sueldo + Aumento: ", sueldo)
}
