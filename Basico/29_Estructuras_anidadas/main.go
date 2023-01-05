package main

import "fmt"

// Persona contiene los campos nombre y apellido
type Persona struct {
	Nombre   string
	Apellido string
}

// Estudiante contiene el campo persona y carrera
type Estudiante struct {
	Persona
	Carrera string
}

// Profesor tiene los campos Estudiante y Carrera
type Profesor struct {
	Estudiante
	Carrera string
}

func main() {
	// Declaramos una variable del tipo Estudiante
	matias := Estudiante{
		Persona{
			Nombre:   "Matías",
			Apellido: "Rodriguez",
		},
		"Informática",
	}
	fmt.Println("Matías: ", matias)

	// Accediendo a los campos
	fmt.Println("Nombre: ", matias.Nombre)
	fmt.Println("Apellido: ", matias.Apellido)
	fmt.Println("Carrera: ", matias.Carrera)

	//Declaramos una clase de tipo Profesor (en múltiples líneas)
	juan := Profesor{
		Estudiante{
			Persona{
				"Juan",
				"López",
			},
			"Contabilidad",
		},
		"Informática",
	}
	fmt.Println("Juan: ", juan)

	// Accediendo a los campos
	fmt.Println("Nombre: ", juan.Nombre)
	fmt.Println("Apellido: ", juan.Apellido)
	fmt.Println("Carrera: ", juan.Carrera)

	fmt.Println("Carrera Estudiantes: ", juan.Estudiante.Carrera)

	// Declarando una variable de tipo Profesor (en una sola linea)
	manuel := Profesor{Estudiante{Persona{"Manuel", "Peralta"}, "Ing. Industrial"}, "Informatica"}
	fmt.Println("Manuel: ", manuel)

	// Declarando una variable de tipo Profesor (asignando valores luego)
	var jose Profesor
	jose.Nombre = "José"
	jose.Apellido = "Contreras"
	jose.Estudiante.Carrera = "Educación"
	jose.Carrera = "Mercadeo"
	fmt.Println("José: ", jose)

}
