package main

import (
	"os"
	"text/template"
)

type Persona struct {
	Nombre string
	Edad   int
	Pais   string
}

func main() {
	persona := Persona{"Carlos", 43, "Argentina"}

	t := template.New("person")              // "person" es un nombre cualquiera
	t, err := t.ParseGlob("templates/*.txt") // acá parsea todos los templates en el directorio
	if err != nil {
		panic(err)
	}
	err = t.ExecuteTemplate(os.Stdout, "residentes", persona) // imprime por salida estándar
	if err != nil {
		panic(err)
	}
}
