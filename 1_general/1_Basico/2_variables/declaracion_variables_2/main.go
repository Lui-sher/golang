/*
Ejemplo Como declarar variables
En este ejemplo intento
mostrar como puedes averiguar que tipo de dato tiene una variable
usando el paquete fmt

para la clase en video se arreglaran errores al final del archivo*/

package main

import "fmt"

func main() {

	var variable_entera int
	var variable_string string
	variable_entera = 5
	variable_string = "texto a imprimir en pantalla \n"
	boolean := true

	/*
		'Printf' permite seleccionar lo que deseamos imprimir respecto a la variable selecinada, si no especificamos lo que deseamos,
		este permite imprimir texto, o cualquier variable tipo texto
	*/

	fmt.Printf("1. Tipo de datos: %T\n", variable_string)                                      // '%T' nos muestra el tipo de la variable, NO agrega salto de linea al final / string
	fmt.Printf("2. Soy la varible 'variable_entera' y soy de tipo: %T\n", variable_entera)     // '%T' nos muestra el tipo de la variable / int
	fmt.Println("3. Soy la varible 'variable_entera' asignada con el valor:", variable_entera) //Println imprime el valor de la variable, y agrega un salto de linea al final
	/*porque no se imprime uno debajo del otro que esta faltando?*/
	fmt.Printf("4. Tipo de datos: %T\n", variable_string) // le agregamos "\n" para generar un salto de linea al final
	fmt.Printf("5. Tipo de datos: %v\n", boolean)         // "%v" nos muestra el valor guardado en la variable, y le agregamos "\n" para generar un salto de linea al final y se vea mejor presentado
	/*para mas detalles visitar https://pkg.go.dev/fmt#pkg-overview */
}
