package main

import (
	"os"
	"text/template"
)

type Persona struct {
	Nombre string
	Edad   int
}

// Este seria nuestro template
const tp = `Nombre: {{.Nombre}}, Edad: {{.Edad}}` + "\n"

// Este seria un template con un range y con un if adentro (logica dentro del template)
const tpmulti = "{{range .}}" +
	"{{if .Edad}}" +
	"Nombre: {{.Nombre}}, Edad: {{.Edad}} - Correcto (tiene una edad asignada).\n" +
	"{{else if .Nombre}}" +
	"Nombre: {{.Nombre}}, Edad: {{.Edad}} - INCORRECTO (NO tiene una edad asignada).\n" +
	"{{else}}" +
	"DATO VACIO, no tiene un nombre ni una edad definida.\n" +
	"{{end}}" +
	"{{end}}"

func main() {
	// Ejemplo con una sola persona
	persona := Persona{"Carlos", 43}

	t := template.New("person") // "person" es un nombre cualquiera
	t, err := t.Parse(tp)       // acá parsea nuestro template
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, persona) // imprime por salida estándar
	if err != nil {
		panic(err)
	}

	// Ejemplo con varias personas
	personas := []Persona{
		{"Miguel", 48},
		{"Agustín", 51},
		{"Sergio", 43},
		{"María", 0}, // Se imprimiria con el msg Incorrecto
		{"", 0},      // Se imprime con el msg DATO VACIO
	}

	// Repetimos el codigo anterior (casi identico, pero cambio "tp" por "tpmulti" y "persona" por "personas")
	t, err = t.Parse(tpmulti)
	if err != nil {
		panic(err)
	}
	err = t.Execute(os.Stdout, personas)
	if err != nil {
		panic(err)
	}

}
