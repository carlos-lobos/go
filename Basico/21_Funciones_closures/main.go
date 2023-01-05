package main

import "fmt"

// Closure

func print(cadena string) {
	fmt.Println(cadena)
}

func print2(cadena string) {
	fmt.Println(cadena)
}
func print3(cadena1, cadena2 string) {
	fmt.Println(cadena1 + cadena2)
}

// Ejemplo de funcion que oficia de wrapper para otra funcion
func print4(fprint func(string), cad string) {
	fprint(cad)
}

// Retorna una funcion que va a retornar un entero
func incrementar() func() int {
	i := 0
	return func() (r int) {
		r = i
		i += 2
		return
	}
}

func incremento() {
	i := 0
	i++
	fmt.Println(i)
}

func main() {

	cadena := "Hola Mundo"

	// Es posible guardar la definicion de una funcion en una variable
	imprimir := print
	imprimir(cadena)

	// Es posible definir una variable de tipo funcion
	imprimir2 := func() {
		fmt.Println(cadena) // cadena esta definida en main, asi que es conocida aqui por estar en el scope global
	}
	imprimir2()

	// Es posible redefinir la variable imprimir, usando otra funcion con la misma firma
	imprimir = print2
	imprimir("Hola Mundo 2")

	// Esto daria error, NO es posible redefinir la variable imprimir, usando otra funcion con una firma diferente
	// imprimir = print3

	// Firmas de las primeras funciones
	fmt.Printf("Funcion print1: %T\n", print)
	fmt.Printf("Funcion print2: %T\n", print2)
	fmt.Printf("Funcion print3: %T\n", print3)
	fmt.Printf("Funcion print3: %T\n", print4)

	// Es posible pasar una funcion por parametro
	print4(print, "Hola Mundo desde Print4")

	// Las funciones son comparables sólamente con nil (funciones vacias)
	var fb func()
	if fb == nil {
		fmt.Println("fb es igual a nil")
	}

	// Las clausuras (closure) mantienen el valor de las variables definidas dentro
	inc := incrementar() // incrementar() retorna una funcion

	fmt.Println("Valor de i: ", inc()) // Recien aqui se ejecuta la funcion por primera vez
	fmt.Println("Valor de i: ", inc())
	fmt.Println("Valor de i: ", inc())
	fmt.Println("Valor de i: ", inc())

	// Para comparar, en esta funcion no se mantiene el valor de las variables definidias dentro
	incremento()
	incremento()
	incremento()

}
