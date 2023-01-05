package main

import "fmt"

// La funcion recibe un puntero de tipo int
func incrementar(numero *int) {
	// para hacer referencia al valor tenemos que usar * delante
	*numero++
	fmt.Println("Valor de Numero dentro de la funcion incrementar: ", *numero)
}

func main() {

	// a es de tipo int
	a := 25
	fmt.Println("Valor de a: ", a)
	fmt.Println("Dirección de a: ", &a)

	// b es un puntero de a, por lo que es de tipo *int
	b := &a

	// Se imprime el contenidos de b
	fmt.Println("Dirección a la que apunta b: ", b)
	// Se imprime el contenido de de la variable a, que es donde apunta b
	fmt.Println("Valor al que apunta b: ", *b)

	// b = 25 // Esto daria error, ya que b es del tipo puntero a int (*int), no de tipo int

	// Le asignamos un nuevo valor a (a) a traves de b
	*b = 32
	fmt.Println("Valor de a: ", a)

	// Incrementamos a
	a++
	// Imprimimos *b
	fmt.Println("Valor al que apunta b: ", *b)

	// // Valor 0 de los punteros es nil
	//
	// if b != nil {
	// 	fmt.Println("b es diferente de nil")
	// }

	// los punteros son comparables
	c := &a
	if b == c {
		fmt.Println("b y c son iguales")
	}

	// Para crear un puntero a int sin apuntarlo todavia a una variable especifica, utilizamos new()
	// En realidad Go crea una variable de tipo int y la inicializa en 0, entonces apunta el puntero a esa variable
	d := new(int) // *int
	fmt.Println("Dirección de d: ", d)
	fmt.Println("Valor de d: ", *d)

	// Ahora si asignamos a d el valor de b (asignacion entre punteros, estamos asignando direcciones de memoria)
	d = b
	fmt.Println("Dirección de d: ", d)
	fmt.Println("Valor al que apunta d: ", *d)

	fmt.Println("Valor de a: ", a)
	fmt.Println("Valor al que apunta b: ", *b)
	fmt.Println("Valor al que apunta c: ", *c)

	// Usar punteros en funciones
	numero := 2
	fmt.Println("Numero antes de incrementar: ", numero)
	// pasamos la direccion donde apunta numero, no la variable numero (porque si pasamos la variable se pasa una copia)
	incrementar(&numero)
	fmt.Println("Numero despues de incrementar: ", numero)

}
