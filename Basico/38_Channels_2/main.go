package main

import (
	"fmt"
	"time"
)

// Ejemplo: Calcula el cuadrado de un nro en un bucle infinito logrado a traves de gorutines
// (usando channels unbuffered de tipo int)

func main() {
	numero := make(chan int)
	cuadrado := make(chan int)

	//Generamos los numeros
	go func() {
		for x := 0; ; x++ {
			numero <- x
		}
	}()

	//Lo elevamos al cuadrado
	go func() {
		for {
			x := <-numero
			cuadrado <- x * x
		}
	}()

	//Imprimimos en main
	for {
		fmt.Println(<-cuadrado)
		time.Sleep(1 * time.Second)
	}

}
