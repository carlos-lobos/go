package main

import (
	"fmt"
	"math"
)

// Rectangulo tiene los campos ancho y alto
type Rectangulo struct {
	ancho, alto float64
}

// Circulo tiene el campo radio
type Circulo struct {
	radio float64
}

// Metodo area() para la estructura Rectangulo
func (r Rectangulo) area() float64 {
	return r.ancho * r.alto
}

// Metodo area() para la estructura circulo
func (c Circulo) area() float64 {
	return c.radio * c.radio * math.Pi
}

// Ejemplo de incrementar valores pasados por valor
func (r Rectangulo) incByValue(i float64) Rectangulo {
	return Rectangulo{
		r.ancho * i,
		r.alto * i,
	}

}

// Ejemplo de incrementar valores pasados por referencia
func (r *Rectangulo) incByRef(i float64) {
	r.ancho *= i
	r.alto *= i
}

func main() {
	// Declaramos dos rectangulos
	r1 := Rectangulo{12, 2}
	r2 := Rectangulo{9, 4}
	// Calculamos e imprimimos sus areas
	fmt.Println("Area de r1 es: ", r1.area())
	fmt.Println("Area de r2 es: ", r2.area())

	// Declaramos dos circulos
	c1 := Circulo{10}
	c2 := Circulo{25}
	// Calculamos e imprimimos sus areas
	fmt.Println("Area de c1 es: ", c1.area())
	fmt.Println("Area de c2 es: ", c2.area())

	// Incrementar rectangulo (por valor)
	fmt.Println("r1: ", r1)
	r1 = r1.incByValue(2)
	fmt.Println("r1: ", r1)

	// Incrementar rectangulo (por referencia)
	r1.incByRef(5)
	fmt.Println("r1: ", r1)
}
