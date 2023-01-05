package main

import "fmt"

func main() {

	// Slices (Segmentos)
	//  - No se debe indicar el tamaño
	//  - Se pueden aplicar todas las caracteristicas de los arrays
	//  - Son secciones de arrays, trabajan con arrays más grandes por debajo
	//  - Cada elemento es un puntero a un elemento de un array

	// Declarar Slices
	var j []int
	fmt.Println("Slice j: ", j)
	fmt.Println("Longitud de j: ", len(j))
	fmt.Println("Capacidad de j: ", cap(j))

	// Declarar Slices incializando algunos valores
	x := []int{1, 2, 3}
	fmt.Println("Slice x: ", x)
	fmt.Println("Longitud de x: ", len(x))
	fmt.Println("Capacidad de x: ", cap(x))

	// Declarando Slices con make(), indicando la longitud
	y := make([]int, 5)
	fmt.Println("Slice y: ", y)
	fmt.Println("Longitud de y: ", len(y))
	fmt.Println("Capacidad de y: ", cap(y))

	// Declarar Slices con make(), indicando la longitud y la capacidad
	k := make([]int, 5, 10)
	fmt.Println("Slice k: ", k)
	fmt.Println("Longitud de k: ", len(k))
	fmt.Println("Capacidad de k: ", cap(k))

	// Definimos un array con 10 elementos de tipo int
	var ar = [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("Array ar: ", ar)

	// Definimos dos slice de tipo []int
	var a, b []int
	fmt.Println("Slice a: ", a)
	fmt.Println("Slice b: ", b)

	// 'a' apunta del 3ro al 5to elemento en el array ar.
	a = ar[2:5] // el ultimo elemento no es incluido
	// ahora 'a' tiene los elementos ar[2],ar[3] y ar[4]
	fmt.Println("Slice a: ", a)

	// 'b' es otro slice del array ar
	b = ar[3:5]
	// Ahora 'b' tiene ar[3] y ar[4]
	fmt.Println("Slice b: ", b)

	b[0] = 25
	fmt.Println("Asignamos b[0] = 25")
	fmt.Println("Slice b: ", b)   // [25, 4]
	fmt.Println("Slice a: ", a)   // [2, 25, 4]
	fmt.Println("Array ar: ", ar) // [0,1,2,25,4,5 ...]
	fmt.Println("Cap(a)", cap(a))
	fmt.Println("Cap(b)", cap(b))

}
