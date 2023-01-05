package main

import (
	"fmt"
	"os"
)

// Defer

func primera() {
	fmt.Println("Primera")
}
func segunda() {
	fmt.Println("Segunda")
}

func main() {

	// Abrimos el archivo
	f, err := os.Open("texto.txt")
	// Verificamos que no haya ocurrido ningún error
	if err != nil {
		panic(err)
	}
	defer f.Close() // Diferimos la ejecucion del cierre del archivo para el final

	// Creamos un Slice para almacenar los que leemos del archivo
	data := make([]byte, 50)
	c, err := f.Read(data)
	// Verificamos que no haya ocurrido ningún error
	if err != nil {
		panic(err)
	}
	// Imprimimos lo leído
	fmt.Printf("Cantidad de byte leidos: %d\n Texto leido: \n%q \nerror: %v\n\n", c, data, err)

	// Ejemplo de que los defer se van sumando en una pila (LIFO)
	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}

	// (Recien aca se ejecutaria el Close() de archivo)
}
