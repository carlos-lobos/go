package main

import (
	"fmt"
	"time"
)

// Ejemplo: Calcula el cuadrado de los primeros 5 nros, cerrando los channels
//          y evitando tener que validar cada vez usando el for + range.
// (usando channels unbuffered de tipo int)

func main() {
	numero := make(chan int)
	cuadrado := make(chan int)

	//Generamos los numeros
	go func() {
		for x := 0; x < 5; x++ {
			numero <- x
		}
		close(numero)
	}()

	//Lo elevamos al cuadrado
	go func() {
		for x := range numero {
			cuadrado <- x * x
		}
		close(cuadrado)
	}()

	//Imprimimos en main
	for x := range cuadrado {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}

}
