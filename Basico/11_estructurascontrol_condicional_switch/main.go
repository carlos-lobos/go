package main

import (
	"fmt"
)

func main() {

	//-- Estructuras de Control --//
	// Condicionales Switch

	// Imprimir el nombre del dia de la semana que el usuario digita

	var dia int
	fmt.Println("Digita el número del día de la semana (inicia 1=lunes)")
	fmt.Scanln(&dia)

	switch dia {
	case 1:
		fmt.Println("Usted digitó Lunes")
	case 2:
		fmt.Println("Usted digitó Martes")
	case 3:
		fmt.Println("Usted digitó Miercoles")
	case 4:
		fmt.Println("Usted digitó Jueves")
	case 5:
		fmt.Println("Usted digitó Viernes")
	case 6:
		fmt.Println("Usted digitó Sabado")
	case 7:
		fmt.Println("Usted digitó Domingo")
	default:
		fmt.Println("Digitó un número inválido")
	}

	fmt.Println("********************")

	// En Go no es necesario usar "break" para cortar la ejecucion cuando un caso se cumple,
	// Al contrario, se usa "fallthrough" para forzar a que se sigan evaluando los case posteriores
	numero := 26

	// Además, se puede colocar la variable y condicion a evaluar dentro de cada case y no necesariamente al lado del switch
	switch {
	case numero > 23:
		fmt.Println("Es Mayor que 23")
		fallthrough
	case numero > 24:
		fmt.Println("Es Mayor que 24")
	case numero > 25:
		fmt.Println("Es Mayor que 25")
	default:
		fmt.Println("Ninguna de las anteriores")
	}

}
