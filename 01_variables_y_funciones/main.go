package main

import (
	"fmt"
)

var hey = "Hey."

func main() {

	//En go hay 4 tipos básicos
	//bool, string, int, float64
	//var carta string = "As de espadas"

	//El compilador infiere en tiempo de compilación que carta de es tipo string
	// la secuencia := indica que estamos definiendo una nueva variable
	carta := "As de basto"

	//Por tanto si queremos reasignar el valor de esta variable usamos simplemente =
	carta = "Rey de oros"

	//Al igual que en Java guardamos el valor que devuelve la funcion en una variable
	carta = nuevaCarta()

	fmt.Println(hey + " Mi carta es " + carta)
}

//string indica que la función devuelve una cadena
func nuevaCarta() string {
	return "Tres de copas"
}
