package main

import "fmt"

func main() {
	// Range

	nombres := []string{
		"Alejandro",
		"Maria",
		"Pedro",
		"Carlos",
	}

	// Range nos da el indice y el valor
	for index, nombre1 := range nombres {
		// Index = 0
		// nombre nombres[index]
		fmt.Printf("El nombre %q esta en el index %d \n", nombre1, index)
	}

	// Podemos desechar el indice indicando "_"
	for _, nombre2 := range nombres {
		fmt.Println(nombre2)
	}

	// Podemos desechar el valor, no incluyendo la 2da variable
	for index1 := range nombres {
		fmt.Println(index1)
	}

	// También podemos recorrer strings
	cadena := "Una buena frase."

	for index2, letra := range cadena {
		fmt.Printf("La letra %q esta en el index %d \n", letra, index2)
	}

}
