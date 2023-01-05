package main

import (
	"fmt"
	"strings"
)

func correr1(r rune) rune {
	return r + 1
}

func main() {

	cadena := "012345678"

	/* Las funciones anonimas son funciones declaradas in-situ (en el lugar), sin nombre

	 	La funcion Map() del paquete "strings" recive una funcion y trabaja con el tipo rune==int32
		Esto es porque para cada caracter de una cadena, lo que se guarda es el Código numérico de la tabla UTF8
	*/
	cadena = strings.Map(func(r rune) rune {
		return r + 1
	}, cadena)
	fmt.Println(cadena)

	// Lo mismo se podria lograr declarando una funcion normalmente, con nombre
	cadena = strings.Map(correr1, cadena)
	fmt.Println(cadena)
}
