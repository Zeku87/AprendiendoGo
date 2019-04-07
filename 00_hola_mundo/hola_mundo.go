//Paquete ejecutable
//Denominamos al paquete con el nombre "main". 
//De esta forma el compilador de Go sabe que este fichero es ejecutable
//Además debe tener una función main definida
package main

//Si el nombre del paquete fuera diferente de "main" entonces
//este fichero no genera ningún ejecutable tras su compilación.


//Importamos todas las funcionalidades implementadas en el paquete "fmt"
//"fmt" es parte de la biblioteca estándar de Go
//Permite sacar información por la salida estándar 
//"fmt" nos permite formatear la entrada y la salida estándar
import "fmt"

func main(){
	fmt.Println("hola mundo")
}