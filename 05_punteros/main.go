package main

import "fmt"

/*
	El uso de los punteros en Go es sintácticamente igual que en C
*/

func main() {

	//la variable i pertenece al ámbito de la función main
	i := 0

	//Pasamos un copia del valor de i, pero NO a la variable i
	//por tanto la dirección de memoria que aloja el valor de i
	//queda inalterada dentro del cuerpo de la función incrementar
	incrementar(i)

	//Por tanto la salida por pantalla es el valor original de i: 0
	fmt.Println("i:", i)

	//Mediante el operador & pasamos la dirección de memoria que aloja el valor de i
	//luego cualquier operación que se haga sobre esa dirección de memoria en el
	//cuerpo de la función incrementarPorReferencia afectará al valor original de i
	incrementarPorReferencia(&i)

	//Luego ahora i sí vale 1
	fmt.Println("i:", i)
}

func incrementar(x int) {
	x++
}

//El asterisco es el operador de indirección lo que significa
//que extrae del valor asociado a la dirección de memoria que se envía como parámetro
func incrementarPorReferencia(x *int) {
	*x++
}
