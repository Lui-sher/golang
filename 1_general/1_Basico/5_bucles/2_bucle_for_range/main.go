/*
Ejemplo ciclos de for
En este ejemplo intento mostrar muchas formas de usar el "for range"
*/
package main

import "fmt"

func main() {

	arreglo := [8]int{0, 1, 4, 6, 10, 9} // declaramos un array de enteros de 8 posiciones pero solo asignamos 6 elementos
	for i, j := range arreglo {          //en este FOR tenemos en cuenta el indice i respecto al tamaño del array (8) y los elementos j (cada item del array)
		fmt.Printf("indice i: %d tiene como valor #%d\n", i, j)
	}
	for i := range arreglo { // en este FOR solo tomamos en cuenta el indice i respecto al tamaño del array (no mostramos los items del array)
		fmt.Printf("Valor de i: %d\n", i)
	}
	for _, j := range arreglo { // Al usar _, estoy indicando que no necesito el índice del elemento, solo me importa el valor j o sea el item almacenado como tal
		fmt.Printf("Valor de i: %d\n", j)
	}

}
