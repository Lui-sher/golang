/*
Ejemplo Como declarar variables
En este ejemplo intento
mostrar todas las formas de declaración que acepta golang

el archivo NO va a correr porque no fue terminado...

te invito a terminarlo usa todas las variables de alguna forma.
Este ejemplo no correra tiene errores
*/

package main

import "fmt"

//import "fmt"

/* funcion principal*/
func main() {

	/*FORMA TRADICIONA*/
	var variable_entera int
	var variable_string string

	/*FORMA ITERATIVA*/
	var entero = 5
	otroentero := 6

	/*FORMA EN GRUPO*/
	var x, y, z int
	var (
		i int
		b bool
		s string
	)

	x1, y2, z3 := 1, 2, 3
	fmt.Println("soy la variable x1 = ", x1) //Println es una funcion dentro del objeto "fmt" que imprime en consola lo que está dentro del parentesis
	fmt.Println("soy la variable y2 = ", y2)
	fmt.Println("soy la variable z3 = ", z3)
	i = 1
	b = true
	s = "un texto cualquiera"
	fmt.Println("soy la variable i = ", i)
	fmt.Println("soy la variable b = ", b)
	fmt.Println("soy la variable s = ", s)

	variable_entera = 5            //asignamos valor a la variable_entera, declarada en la linea 22
	variable_string = "otro texto" //asignamos valor a la variable_string, declarada en la linea 23
	fmt.Println("soy la variable_entera = ", variable_entera)
	fmt.Println("soy la variable_string = ", variable_string)

	/*FORMA ITERATIVA*/
	fmt.Println("soy la variable entero = ", entero)
	fmt.Println("soy la variable otroentero = ", otroentero)

	/*FORMA EN GRUPO*/
	x = 10
	y = 44
	z = 66
	fmt.Println("soy la variable x = ", x)
	fmt.Println("soy la variable y = ", y)
	fmt.Println("soy la variable z = ", z)

	/*
	   para hacer funcionar este archivo debe usar todas la variables,
	   te recomiendo que imprimas en pantallas cada una de las variables declaradas anteriormente,
	   no olvides asignarle algun valor
	*/

	/* NOTA IMPORTANTE SIEMPRE DEBES USAR LAS VARIABLES QUE HAS DECLARADO*/

}
