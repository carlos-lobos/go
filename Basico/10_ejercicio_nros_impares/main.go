package main

import "fmt"

func main() {

	// Declaración de las variables a utilizar
	var numero1 int
	var numero2 int
	var contador int // Contador de Números Impares

	encabezado := `
	****************************
	Contador de Numeros Impares
	****************************
	`

	// Imprimimos el encabezado
	fmt.Println(encabezado)

	// Solicitamos los 2 numeros
	fmt.Println("Digita el Primer numero (número desde)")
	fmt.Scanln(&numero1) // Lectura de entrada por teclado desde la consola

	fmt.Println("Digita el Segundo numero (número hasta)")
	fmt.Scanln(&numero2)

	// Realizamos un bucle tomando como inicio y fin los numeros digitados
	for i := numero1; i <= numero2; i++ {
		// Evaluamos si el numero es impar
		if i%2 != 0 {
			// incrementamos el valor de la variable contador en 1
			contador++
			// Imprimimos el numero impar
			fmt.Printf("%d es un numero impar, continuando ... \n", i)
		}
	}

	encabezado = `
	*********************
	Resultado del Conteo
	*********************
	`
	// Imprimimos el encabezado
	fmt.Println(encabezado)
	// Imprimimos los resultados
	fmt.Printf("Entre el %d y el %d hay %d numeros impares.\n", numero1, numero2, contador)

}
