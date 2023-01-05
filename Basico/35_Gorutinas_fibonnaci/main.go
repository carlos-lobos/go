package main

import (
	"fmt"
	"time"
)

func main() {
	// Ira imprimiendo un spinner en la terminal
	go animacion(100 * time.Millisecond)

	// Declaración de una constante
	const n = 45

	// La funcion fib() devolverá el nro que se encuentra en la posición n
	// la funcion tarda para "n" grandes porque debe llamarse a si misma "n" veces
	resultado := fib(n)
	fmt.Printf("\rFibonacci(%d) = %d\n", n, resultado)

	// La función comentada abajo es mas rápida para mostrar todos los numeros de la sucesion de fibonnaci,
	// desde la posición "0" hasta la "n", porque recuerda los valores anteriores y los usa,
	// en lugar de tener que calcularlos cada vez.

	// print_all_fib(n)
}

func animacion(retraso time.Duration) {
	for {
		for _, r := range `\|/-` {
			fmt.Printf("\r%c", r)
			time.Sleep(retraso)
		}
	}
}

func fib(x int) int {
	if x < 2 {
		return x
	}
	return fib(x-1) + fib(x-2)
}

func print_all_fib(x int) {
	resultado := 0
	r1 := 0 // Resultado anterior sub 1
	r2 := 0 // Resultado anterior sub 2

	for i := 0; i <= x; i++ {
		if i == 1 {
			r1 = resultado
		} else if i >= 2 {
			r2 = r1
			r1 = resultado
			resultado = r1 + r2
			fmt.Printf("\rFibonacci(%d) = %d\n", i, resultado)
			continue
		}

		// La funcion fib() devolverá el nro que se encuentra en la posición n
		// (aca solo la uso para los resultados de las posiciones 0 y 1)
		resultado = fib(i)
		fmt.Printf("\rFibonacci(%d) = %d\n", i, resultado)
	}
}
