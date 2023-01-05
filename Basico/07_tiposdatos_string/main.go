package main

import (
	"fmt"
	"strconv"
)

func main() {

	// Tipo de Dato String
	// Los Strings en Go, son una secuencia de Bytes (Slice de Bytes)
	// Son Indexables
	// Son Inmutables

	var cadena string
	fmt.Println(cadena)
	cadena = "Aguas Saborizadas"
	fmt.Println(cadena)
	fmt.Println(len(cadena)) // 17

	// Acceder a una posicion del string (los indices comienzan desde 0)
	fmt.Println(cadena[16]) // Resultado es el Byte, es decir, el UTF8 "115" que es la "s" final en "Aguas Saborizadas"

	// Subcadenas (Posicion INICIAL:Posicion FINAL), la posicion final no es incluyente
	fmt.Println(cadena[0:4]) // Agua
	fmt.Println(cadena[2:4]) //   ua
	fmt.Println(cadena[:])

	//cadena[2] = "l"

	// Concatenacion
	cadena = cadena + " de Manzana!"
	fmt.Println(cadena)

	cadena += " Oh yeah!"
	fmt.Println(cadena)

	// Cadenas Multilinea (se toman tal cual con los espacios incluidos)
	// (CUIDADO! En Go, no se pueden definir cadenas entre comillas simples)
	cadena = `
	<html>
	    <head>
	        <meta charset="utf-8">
	        <title></title>
	    </head>
	    <body>
	    </body>
	</html>
	`
	fmt.Println(cadena)

	// Escapado de comillas dobles o indicacion de tabulaciones
	cadena = "(Barra invertida) \\ (tab) \t (comillas dobles) \"25\""
	fmt.Println(cadena)

	// Conversion de nro entero a string para poder concatenarlo
	edad := 29
	cadena = "La edad es " + strconv.Itoa(edad)
	fmt.Println(cadena)
	fmt.Println("Edad", edad)

}
