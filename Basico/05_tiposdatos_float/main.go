package main

import "fmt"

func main() {
	// Datos de Tipo Float
	var f32 float32     // 6 dígitos decimales de precisión
	var f64 float64     // 15 dígitos decimales de precisión (se recomienda usar este)
	var c64 complex64   // Numero complejo con partes real e imaginaria float32
	var c128 complex128 // Numero complejo con partes real e imaginaria float64

	fmt.Println("f32, f64, c64, c128 = ", f32, f64, c64, c128)
	f32 = 0.167
	f64 = 0.167
	fmt.Println("f32 * 51.365636 = ", f32*51.365636)
	fmt.Println("f64 * 51.365636 = ", f64*51.365636)
	c64 = complex(5, 6)
	fmt.Println("c64 = ", c64)
	fmt.Println("c64 * 73835.11 = ", c64*73835.11)
	c128 = complex(5, 6)
	fmt.Println("c128 = ", c128)
	fmt.Println("c128 * 73835.11 = ", c128*73835.11)
	fmt.Println("parte real - c128 * 73835.11 = ", real(c128*73835.11))
	fmt.Println("parte imaginaria - c128 * 73835.11 = ", imag(c128*73835.11))

}
