package main

import "fmt"

//Creación de un struct: colección de campos
type trabajador struct {
	nombre string
	sueldo float64
}

func main() {

	//Definimos un trabajador
	trabajador1 := trabajador{nombre: "Juan", sueldo: 1345.56}

	//Sacamos su representación string por pantalla
	fmt.Println(trabajador1)

	//Es posible acceder a los campos mediante el operador punto
	if trabajador1.sueldo > 1200 {
		fmt.Println("Sueldo superior al de un programador junior")
	}

	//Además podemos manipular el valor de los campos
	trabajador1.nombre = "Jesus"
	fmt.Println(trabajador1)
}
