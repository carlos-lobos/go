package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// Enteros CON signo
	var entero8 int8   // 8-bit (-128 a 127)
	var entero16 int16 // 16-bit (-32768 a 32767)
	var entero32 int32 // 32-bit (-2147483648 a 2147483647)
	var entero64 int64 // 64-bit (-9223372036854775808 a 9223372036854775807)
	fmt.Println(entero8, entero16, entero32, entero64)

	// Enteros SIN signo
	var uentero8 uint8   // 8-bit (0 a 255)
	var uentero16 uint16 // 16-bit (0 a 65535)
	var uentero32 uint32 // 32-bit (0 a 4294967295)
	var uentero64 uint64 // 64-bit (0 a 18446744073709551615)
	fmt.Println(uentero8, uentero16, uentero32, uentero64)

	// Aliases
	var enteroByte byte // == uint8
	var enteroRune rune // == int32
	fmt.Println(enteroByte, enteroRune)

	// Tipos dependientes de la plataforma
	var enteroUint uint       // segun arquitectura 32 o 64 bits
	var enteroInt int         // segun arquitectura 32 o 64 bits
	var enteroUintptr uintptr // entero SIN signo, lo suficientemente grande como para contener el valor de un puntero
	fmt.Println(enteroUint, enteroInt, enteroUintptr)

	// Conversion entre tipos
	entero32 = 10
	entero64 = 20
	fmt.Println(entero32 + int32(entero64))

	enteroRune = 30
	fmt.Println(entero32 + enteroRune)

	enteroInt = 50
	fmt.Println(entero32 + int32(enteroInt))

	// Tamaño en bytes de los numeros
	// 4 bytes (4*8bits = 32 bits)  ;  8 bytes (8*8bits = 64 bits)
	fmt.Println(unsafe.Sizeof(entero32), unsafe.Sizeof(enteroInt))
}
