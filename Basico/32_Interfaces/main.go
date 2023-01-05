package main

import "fmt"

// Interface Usuario
type Usuario interface {
	Nombre() string
	Email() string
}

// Estructura Persona
type Persona struct {
	nombre string
	email  string
	edad   int
}

// Método Nombre() para una Persona
func (p Persona) Nombre() string {
	return p.nombre
}

// Método Email() para una Persona
func (p Persona) Email() string {
	return p.email
}

// Como se ve, Persona implementa los 2 métodos necesarios para que se pueda considerar que cumple con la interface Usuario
// En Go no es necesario declarar que Persona implementa Usuario

// Estructura Moderador ...
type Moderador struct {
	Persona
	Foro string
}

// Método AbrirForo para un Moderador
func (m Moderador) AbrirForo() {
	fmt.Println("Abrir un Foro")
}

// Método CerrarForo para un Moderador
func (m Moderador) CerrarForo() {
	fmt.Println("Cerrar un Foro")
}

// Estructura Administrador
type Administrador struct {
	Persona
	Sección string
}

// Método CrearForo para un Administrador
func (a Administrador) CrearForo() {
	fmt.Println("Abrir un Foro")
}

// Método EliminarForo para un Administrador
func (a Administrador) EliminarForo() {
	fmt.Println("Cerrar un Foro")
}

// Funcion Presentarse, recive cualquier variable que pueda representar una interface Usuario
// (Solo se tiene acceso a los métodos declarados en el contrato de la interface)
func Presentarse(p Usuario) {
	fmt.Println("Nombre: ", p.Nombre())
	fmt.Println("Email: ", p.Email())
}

func main() {

	alejandro := Persona{"Alejandro", "a@hmail.com", 29}
	Presentarse(alejandro)
	juan := Moderador{Persona{"Juan", "j@hmail.com", 46}, "Juegos"}
	pedro := Administrador{Persona{"Pedro", "p@hmail.com", 25}, "PC"}
	Presentarse(juan)
	Presentarse(pedro)

	var i Usuario
	// i encampsula a alejandro que es de tipo Persona
	i = alejandro
	fmt.Println("i: ", i)
	fmt.Println("i: ", i.Email())
	fmt.Println("i: ", i.Nombre())

	// i, tambien puede encampsular a juan o pedro, que son de tipos diferentes,
	// pero que al componerse de Persona, tambien satisfacen a la interface Usuario ...
	i = juan
	fmt.Println("i: ", i)
	fmt.Println("i: ", i.Email())
	fmt.Println("i: ", i.Nombre())
	// ... pero no es posible usar los métodos propios de juan o pedro,
	// solo son alcanzables los métodos declarados en el contrato de la interface

}
