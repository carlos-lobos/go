package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Creamos un nuevo enrutador
	myServeMux := http.NewServeMux()

	// Indicamos las rutas a manejar
	myServeMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hola Mundo</h1>")
	})

	myServeMux.HandleFunc("/prueba", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hola Mundo desde /prueba</h1>")
	})

	myServeMux.HandleFunc("/usuario", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Hola Usuario</h1>")
	})

	// Ponemos a escuchar el servidor en el puerto 8080 indicándole nuestro enrutador
	http.ListenAndServe(":8080", myServeMux)
}
