package main

import (
	"fmt"
	"strconv"
)

func main() {

	//Un slice es un array que crece y decrece dinámicamente
	//por lo que no tendremos que establecer un tamaño fijo
	//Además posee funciones prediseñadas que nos permiten manipularlo
	//Es comparable al ArrayList de Java
	//Los elementos han de ser del mismo tipo
	cartas := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	palos := []string{"basto", "oro", "espada"}

	//la función append añade un nuevo elemento al final del slice
	palos = append(palos, "copa")

	//permiten su representación por la salida estándar como un string
	fmt.Println(cartas)
	fmt.Println(palos)

	//mediante [] podemos apuntar al elemento deseado
	//los elementos de cartas son de tipo int y los convertimos a string
	//para concatenar su valor con el resto de la salida
	fmt.Println(strconv.Itoa(cartas[3]) + " de " + palos[3])

	//Iteramos usando un for. En esencia es un foreach
	//La variable i indica la posición actual
	//La variable carta sostiene el valor del elemento actual
	for i, carta := range cartas {
		fmt.Println(carta)
	}

	//Ahora anidamos dos estructuras for para sacar todas las cartas por pantalla
	for i := range palos {
		for j := range cartas {
			fmt.Println(strconv.Itoa(cartas[j]) + " de " + palos[i])
		}
	}

}
