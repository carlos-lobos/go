package main

import "fmt"

// Panic y Recover

func imprimir() {
	fmt.Println("Hola Alejandro")

	// Difiero una funcion anonima para recobrar cualquier error
	defer func() {
		cadena := recover()
		fmt.Println(cadena, "la funcion anonima diferida recobro el error.")
	}()

	// panic(msg): cortaria la ejecucion directamente,
	// pero como se usa la funcion recover() de forma diferida, se continua
	panic("ERROR:")

	// (Cualquier codigo debajo del panic no se va a ejecutar)
}

func main() {
	imprimir()
	fmt.Println("Hola Main")
}
