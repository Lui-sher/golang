/*
Ejemplo goto
Muestro como saltar a una etiqueta con goto
*/
package main

import "fmt"

func main() {
	var i int

CICLO1: //permite hacer saltos en el codigo, Muy util
	fmt.Println("estamos fuera del for")
	for i < 10 {
		if i == 6 {
			i += 3 //i = i + 3
			fmt.Println("Saltando a etiqueta CICLO con i = i + 3")
			goto CICLO2 //goto nos hace saltar al CICLO2
		}
		fmt.Printf("Valor de i: %d\n", i)
		i++
	}
CICLO2:
	fmt.Printf("ciclo 2 Valor de i: %d\n", i)
	if i == 9 {
		fmt.Printf("Valor de i: %d\n", i)
		i += 3
		fmt.Println("Saltando a etiqueta CICLO con i = i + 3")
		goto CICLO1 //goto nos hace saltar al CICLO1
	}
	fmt.Printf("terminamos\n")
}
