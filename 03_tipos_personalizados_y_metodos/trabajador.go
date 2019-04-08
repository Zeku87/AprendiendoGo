package main

import "fmt"

//con la palabra clave type podemos crear tipos personalizados
//cuyo comportamiento es el mismo que el tipo escrito más a la derecha
//por ejemplo el tipo trabajadores se comporta igual que un slice de string
type trabajadores []string

//A continuación veremos que definiremos métodos referidos
//al tipo trabajadores extendiendo así su funcionalidad

//Esta función solo puede ser accedida por el tipo trabajadores
//y esto lo indicamos tras la palabra func como (t trabajadores)
//t trabajadores se denomina recibidor en Go
func (t trabajadores) mostrar() {
	var count = len(t)
	for index := 0; index < count; index++ {
		fmt.Println(t[index])
	}
}
