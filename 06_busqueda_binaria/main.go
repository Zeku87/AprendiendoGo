package main

import "fmt"

/*
	Dado una lista de valores enteros ordenados podemos realizar
	una búsqueda binaria.

	En el peor caso tendremos que realizar log n en base 2 más 1
	donde n es el tamaño de la lista
*/
func main() {
	listaOrdenada := [10]int{1, 4, 7, 9, 11, 16, 17, 23, 34, 45}
	resultado := busquedaBinaria(listaOrdenada)

	fmt.Println(resultado)
}

func busquedaBinaria(listaOrdenada [10]int) int {

	//Receta

	//Primero determinamos el elemento mitad
	elementoMitad := len(listaOrdenada) / 2

	return 0
}
