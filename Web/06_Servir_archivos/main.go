package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

// Estas 3 son funciones que manejan cada una de las rutas por separado
/*func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola Mundo</h1>")
}*/
func prueba(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola Mundo desde /prueba</h1>")
}
func usuario(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hola Usuario</h1>")
}

// Otra forma de hacer lo mismo: Definimos nuestro propio Handler
type mensaje struct {
	msg string
}

func (m mensaje) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, m.msg)
}

func main() {
	// Creamos un nuevo enrutador
	myServeMux := http.NewServeMux()

	// Indicamos las rutas a manejar

	// myServeMux.HandleFunc("/", home)
	fs := http.FileServer(http.Dir("public"))
	myServeMux.Handle("/", fs)

	myServeMux.HandleFunc("/prueba", prueba)
	myServeMux.HandleFunc("/usuario", usuario)

	// De esta forma se usa cuand tenemos nuestro propio handler definido
	mesg := mensaje{
		msg: "<h1>Hola, soy un mensaje nuevo</h1>",
	}
	myServeMux.Handle("/hola", mesg)

	// Ponemos a escuchar el servidor en el puerto 8080 indicándole nuestro enrutador
	// http.ListenAndServe(":8080", myServeMux)

	server := &http.Server{
		Addr:         ":8080",
		Handler:      myServeMux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		// << se le conoce como operador binario
		MaxHeaderBytes: 1 << 20, // == 1 MB == 1048576 Bytes [ 1*pow(2,20) == 1 multiplicado por 2, 20 veces ]ng ...")
	}

	log.Println("Listening ...")
	log.Fatal(server.ListenAndServe())
}
