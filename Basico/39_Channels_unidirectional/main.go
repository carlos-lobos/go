package main

import (
	"fmt"
	"time"
)

// Partiendo de 38_Channels_2/main3.go:
// Se refectorizaron las funciones anonimas en funciones normales para poder definir los canales unidireccionales,
// indicandole a Go el tipo de channel en la definicion del parametro
// (channel de salida "chan<-" y channel de entrada "<-chan")

func generarNumeros(out chan<- int) {
	for x := 0; x < 5; x++ {
		out <- x
	}
	close(out)
}

func elevarAlCuadrado(in <-chan int, out chan<- int) {
	for x := range in {
		out <- x * x
	}
	close(out)

	// Estas lineas darian error porque:
	// close(in) => no se puede   cerrar       un canal unidireccional en el cual solo recibimos datos.
	// in <- 23  => no se puede   enviar por   un canal unidireccional en el cual solo recibimos datos.
}

func imprimir(in <-chan int) {
	for x := range in {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	numero := make(chan int)
	cuadrado := make(chan int)

	// Generamos los numeros
	go generarNumeros(numero)

	// Lo elevamos al cuadrado
	go elevarAlCuadrado(numero, cuadrado)

	// Imprimimos en main
	imprimir(cuadrado)
}
