package main

import "fmt"

// Punto contiene los valores x, y
type Punto struct {
	x, y int
}

// Punto3D contiene los valores x,y,z y un puntero a la misma estructura Punto3D (recursivamente)
type Punto3D struct {
	x, y, z int
	*Punto3D
}

// OpPunto es una estructura vacia, podria contener los metodos para realizar las operaciones con puntos
// type OpPunto struct {
// }

func main() {
	// Una estructura se inicializa con los valores por defecto de cada uno de sus atributos
	p := Punto{}
	fmt.Println("p: ", p)

	// Declaracion de una estructura con puntero a la misma estructura
	p2 := Punto3D{
		5,
		6,
		4,
		&Punto3D{
			6,
			4,
			6,
			nil, // para terminar de enlazar recursivamente
		},
	}
	fmt.Println("p2: ", p2)
	fmt.Println("p2: ", *p2.Punto3D)

	// Comparando Estructuras
	a := Punto{5, 6}
	b := Punto{7, 4}
	fmt.Println("a == b : ", a == b)

	c := Punto{7, 4}
	fmt.Println("b == c :", b == c)

	// Puede usarse una estructura como indice de un map, porque son comparables
	figuras := make(map[Punto]string)
	figuras[a] = "Hola Mundo"
	fmt.Println("figuras[a] :", figuras[a])

}
