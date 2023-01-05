package main

import "fmt"

func main() {
	// Arrays
	//  - Son inmutables
	//  - El compilador necesita si o si saber el tamaño que va a tener el array

	// Declaraciones
	var x [3]int
	fmt.Println(x)

	var k [3][2]float64
	fmt.Println(k)

	// Asignacion de valor
	x[1] = 25
	fmt.Println(x)

	// Acceder a un valor por su indice
	fmt.Println(x[1])

	// Declarar e inicializar Arrays (indicando el tamaño)
	y := [5]int{1, 2, 3, 4, 5}
	fmt.Println(y)

	// Declarar e inicializar Arrays usando elipse (...)
	// (... toma como tamaño la cantidad de elementos indicados como literal entre llaves)
	j := [...]int{1, 2, 3, 4}
	fmt.Println(j)

	// Declarar e inicializar Arrays (se pueden colocar en múltiples líneas, el último elemento debe tener coma al final)
	i := [...]int{
		1,
		2,
		3,
		//4,
	}
	fmt.Println(i)

	// Pueden contener elementos de cualquier Tipo
	f := [...]float64{1.2, 1.5, 8.3}
	fmt.Println(f)

	g := [3]string{"1.2", "1.5", "8.3"}
	fmt.Println(g)

	// Funcion len() indica el tamaño de un Array (o de un string tambien)
	fmt.Println(len(f))

	// Imprimir el ultimo elemento de un Array
	fmt.Println(f[len(f)-1])

	// Comparar Arrays
	a := [3]int{1, 2, 3}
	b := [...]int{1, 2, 3}
	c := [3]int{1, 2, 4}

	fmt.Println(a == b)
	fmt.Println(a == c)

	// Es posible inicializar solo algunos de los elementos
	d := [4]int{1, 2}
	fmt.Println(d)

}
