package main

import "fmt"

/*
	Dado una lista de valores enteros ordenados podemos realizar
	una búsqueda binaria.

	En el peor caso tendremos que realizar log n en base 2 más 1
	donde n es el tamaño de la lista
*/

func main() {
	listaOrdenada := []int{0, 1, 4, 6, 12, 13, 16, 23, 27, 38, 59, 63, 71, 116, 126, 210}
	resultado := busquedaBinaria(listaOrdenada, 6)

	fmt.Println(resultado)
}

//Precondición: el tamaño de la lista debe ser mayor que cero
func busquedaBinaria(listaOrdenada []int, valor int) bool {

	//En cada llamada observamos la transformación que sufre el slice
	fmt.Println("Lista de valores", listaOrdenada)

	//Hacemos cumplir la precondición
	if len(listaOrdenada) == 0 {
		return false
	}

	//Receta

	//Primero determinamos el elemento mitad redondeando la parte entera hacia abajo
	//En Go la parte decimal se trunca por lo tanto lo anterior siempre se cumple
	indiceMitad := len(listaOrdenada) / 2

	//Comparamos el valor del elemento mitad con el valor que se pide encontrar

	//Si el valor del elemento mitad es igual al valor buscado devolvemos true
	if listaOrdenada[indiceMitad] == valor {
		return true
	}

	//Si el valor es más pequeño que el elemento mitad entonce eliminamos los elementos de la derecha
	//inclyendo al elemento mitad y volvemos a llamar a busquedaBinaria con la nueva lista
	if listaOrdenada[indiceMitad] > valor {
		//desde cero hasta indice mitad no incluido
		return busquedaBinaria(listaOrdenada[:indiceMitad], valor)
	}

	//Y al contrario, cuando el valor es más grande que el elemento mitad entonces descartamos
	//los elementos de la izquierda incluyendo al elemento mitad y volvemos a empezar
	if listaOrdenada[indiceMitad] < valor {
		//desde indice mitad mas uno hasta el final
		return busquedaBinaria(listaOrdenada[indiceMitad+1:], valor)
	}

	//Por defecto devolvemos false
	return false
}

//27
