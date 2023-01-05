package main

import "fmt"

func main() {
	// Operadores aritméticos básicos
	a := 5
	b := 2
	fmt.Println("5 + 2 = ", a+b)
	fmt.Println("5 - 2 = ", a-b)
	fmt.Println("5 * 2 = ", a*b)
	fmt.Println("5 / 2 = ", a/b)
	fmt.Println("5 % 2 = ", a%b)

	fmt.Println("********************")

	// Operadores de comparación
	c := 2
	fmt.Println("5 == 2 = ", a == b)
	fmt.Println("2 == 2 = ", b == c)
	fmt.Println("5 != 2 = ", a != b)
	fmt.Println("2 != 2 = ", b != c)
	fmt.Println("5 < 2 =  ", a < b)
	fmt.Println("5 <= 2 = ", a <= b)
	fmt.Println("2 <= 2 = ", b <= c)
	fmt.Println("5 > 2 = ", a > b)
	fmt.Println("5 >= 2 = ", a >= b)
	fmt.Println("2 >= 2 = ", b >= c)

	fmt.Println("********************")

	// Operadores en las asignaciones
	/*
		a += b || a = a + b
		a -= b || a = a - b
		a *= b || a = a * b
		a /= b || a = a / b
		a %= b || a = a % b
	*/
	fmt.Println("5 += 2 = ", a)

	fmt.Println("********************")

	// Operadores lógicos
	/*
		&& (AND)
		|| (OR)
		!  (Negación)
	*/
	fmt.Println("&& (AND)")
	fmt.Println("true && true = ", true && true)
	fmt.Println("true && false = ", true && false)
	fmt.Println("false && true = ", false && true)
	fmt.Println("false && false = ", false && false)
	fmt.Println("|| (OR)")
	fmt.Println("true || true = ", true || true)
	fmt.Println("true || false = ", true || false)
	fmt.Println("false || true = ", false || true)
	fmt.Println("false || false = ", false || false)
	fmt.Println("! (Negación)")
	fmt.Println("!true = ", !true)
	fmt.Println("!false = ", !false)

	fmt.Println("********************")

	// Jerarquía de los operadores
	/*
		()
		* / %
		+ -
		== != < <= > >=
		&&
		||
	*/

	fmt.Println("5-3*5 = ", 5-3*5)
	fmt.Println("(5-3)*5 = ", (5-3)*5)

	// Operadores de incremento y decremento
	/*
		En Go, estos operadores son sentencias,
		lo que significa que hay que usarlos en una linea separada,
		no es posible utilizarlos dentro de una expresión.
	*/
	fmt.Println("a = ", a)
	a++
	fmt.Println("a++ = ", a)
	a--
	fmt.Println("a-- = ", a)
}
