package main

/*
Parados sobre el directorio del proyecto, ejecutar:
+ Para crear el archivo de configuracion del modulo por defecto (go.mod)
    shell> go mod init main
+ Para instalar gorilla mux
	shell> go get github.com/gorilla/mux
*/

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

/* Probar que funcionen:
+ curl -X GET http://localhost:8080/api/users
+ curl -X POST http://localhost:8080/api/users
+ curl -X PUT http://localhost:8080/api/users
+ curl -X DELETE http://localhost:8080/api/users
*/

func GetUsers(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde el método GET\n")
}
func CreateUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde el método POST\n")
}
func ModUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde el método PUT\n")
}
func DelUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Mensaje desde el método DELETE\n")
}

func main() {
	// Crear nuestro Serve Mux usando Gorilla Mux (Crear nuestro Enrutador).
	// StrictSlash(false): hace que se tomen como rutas diferentes cuando terminan o no con barra (slash).
	//                     Ej: /api/user y /api/user/ seran tomadas como distintas rutas.
	r := mux.NewRouter().StrictSlash(false)

	r.HandleFunc("/api/users", GetUsers).Methods("GET")
	r.HandleFunc("/api/users", CreateUser).Methods("POST")
	r.HandleFunc("/api/users", ModUser).Methods("PUT")
	r.HandleFunc("/api/users", DelUser).Methods("DELETE")

	server := &http.Server{
		Addr:           ":8080",
		Handler:        r,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Println("Listening ...")
	log.Fatal(server.ListenAndServe())
}
