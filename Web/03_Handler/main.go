package main

import (
	"fmt"
	"net/http"
)

// Estas 3 son funciones que manejan cada una de las rutas por separado
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola Mundo</h1>")
}
func prueba(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola Mundo desde /prueba</h1>")
}
func usuario(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola Usuario</h1>")
}

func main() {
	// Creamos un nuevo enrutador
	myServeMux := http.NewServeMux()

	// Indicamos las rutas a manejar
	myServeMux.HandleFunc("/", home)
	myServeMux.HandleFunc("/prueba", prueba)
	myServeMux.HandleFunc("/usuario", usuario)

	// Ponemos a escuchar el servidor en el puerto 8080 indicándole nuestro enrutador
	http.ListenAndServe(":8080", myServeMux)
}
