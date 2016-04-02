package main

import (
	"fmt"
)

type Persona struct {
	nombres   string
	apellidos string
	edad      int
}

func main() {

	var per Persona // declaramos una variable de tipo Persona struct

	per.nombres = "Bryan Diego"    // Asignamos el campo nombres
	per.apellidos = "Perez Delfin" // Asignamos el campo apellidos
	per.edad = 23                  // Asignamos el campo edad

	// Mostramos por consola los campos de la estructura
	fmt.Println("La persona ha sido creada y tiene nombres - apellidos - edad", per.nombres, per.apellidos, per.edad)

}
