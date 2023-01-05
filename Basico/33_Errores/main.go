package main

import (
	"errors"
	"fmt"
)

// Declaracion de los Errores
var (
	ErrorUsuarioNoValido  = errors.New("El usuario no es valido")
	ErrorUsuarioEnProceso = errors.New("Usuario en proceso de registro")
	ErrorPorDefecto       = errors.New("Algo paso xd...")
)

func baneado(usuario string) (err error) {
	ban := false
	switch usuario {
	case "miguel":
		ban = true
	case "carlos":
		ban = false
	case "juan":
		return ErrorUsuarioNoValido
		// Una alternativa es usar el metodo Errorf del paquete fmt
		// return fmt.Errorf("El usuario no es valido")
	case "pedro":
		return ErrorUsuarioEnProceso
		// return fmt.Errorf("Usuario en proceso de registro")
	default:
		return ErrorPorDefecto
		// return fmt.Errorf("Algo paso xd...")
	}

	if !ban {
		fmt.Printf("El usuario %s no esta baneado \n", usuario)
	} else {
		fmt.Printf("El usuario %s esta baneado \n", usuario)
	}

	return nil
}

func checkError(err error) {
	if err != nil {
		fmt.Println("Error: ", err)
	}
}

func main() {

	err := baneado("miguel")
	checkError(err)

	err = baneado("juan")
	checkError(err)

	err = baneado("carlos")
	checkError(err)

	err = baneado("pedro")
	checkError(err)

	err = baneado("pololo")
	checkError(err)

}
