package main

import "fmt"

func main() {

	//-- Estructuras de Control --//
	// Bucles For

	// for true { ... }
	fmt.Println("Nros 1 al 5:")
	i := 1
	for i <= 5 {
		fmt.Println(i)
		i++
	}

	fmt.Println("Nros 1 al 5 (Otra forma):")
	// for initialize; condition; increment { ... }
	for j := 1; j <= 5; j++ {
		fmt.Println(j)
	}

	// Estilo do{ ... }while()
	for {
		fmt.Println("Se hace algo y despues se corta por alguna condicion")
		break
	}

}
