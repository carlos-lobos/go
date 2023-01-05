package main

import (
	"fmt"
	"strconv"
	"time"
)

// Creando channels con buffer

func enviarMensaje(out chan<- string, numero int) {
	for {
		out <- "Mensaje : " + strconv.Itoa(numero)
		fmt.Println("Enviando mensaje func: ", numero)
	}
}

func imprimir(in <-chan string) {
	for x := range in {
		fmt.Println(x)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// Crear channel de tipo string con un buffer que albergará hasta 5 mensajes
	ch := make(chan string, 5)
	for i := 1; i < 5; i++ {
		// Creamos 4 go-rutinas que trabajaran con el mismo channel
		go enviarMensaje(ch, i)
	}
	// Vamos imprimiendo
	imprimir(ch)
}
