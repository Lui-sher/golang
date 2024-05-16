/*
Ejemplo if complejo
En este ejemplo intento mostrar
un if diferente que puedes conseguir en ejemplos
de la documentación oficial y en foros de internet
*/
package main

import (
	"fmt"
	"os"
)

func main() {

	var rutaAbuscar string
	rutaAbuscar = "C:/Users/lui_s/OneDrive/Cursos de Programacion/GO/Curso Golang/golang/1_general/1_Basico/4_condicional/2_if_complejos/main.go" //en la terminal toca si o si usar 'go run main.go'
	//esta es la ruta absoluta, el codigo sabe donde tiene que buscar el archivo 'main.go'

	if _, err := os.Stat(rutaAbuscar); !os.IsNotExist(err) {
		fmt.Println("existe")
	} else {
		fmt.Println("no existe")
	}
	rutaAbuscar = "./luis.go" // este achivo claramente no existe en la ruta
	// Esta es la ruta realativa, el codigo usa como referencia el mismo punto donde se ejecuta el archivo main.go
	if _, err := os.Stat(rutaAbuscar); os.IsNotExist(err) {
		fmt.Println("existe")
	} else {
		fmt.Println("no existe")
	}
}

/* Nota:
os.Stat('ruta') lo que hace es buscar un archivo en la ruta proporcionada, nos retornará dos posibles valores, uno con el contenido del archivo y otro con un error, en caso de no encontrar el archivo

El guion bajo ( _ ) despues del 'if' quiere decir que no nos importa el contenido del archivo buscado en la ruta dada

os.IsNotExist(err) en el segundo caso esta funcion nos lanza un true

*/
