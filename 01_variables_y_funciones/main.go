package main

import (
	"errors"
	"fmt"
)

var hey = "Hey."

func main() {
	testFuncionesMulti()
}

func tiposBasicosDeDatos() {
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

/*
	Desde la siguiente función hacemos llamadas a la función dividir
	Importamos el paquete errors para notificar que no se puede dividir por cero
*/
func testFuncionesMulti() {

	resultado, error := dividir(43, 0)

	if error == nil {
		fmt.Println("Resultado de la division", resultado)
	} else {
		fmt.Println(error)
	}
}

func dividir(a float64, b float64) (float64, error) {

	if b == 0 {
		return 0, errors.New("Es incongruente dividir por cero")
	}
	return a / b, nil
}
