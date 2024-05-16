/*
Ejemplo Como declarar variables
En este ejemplo intento
como usar struct basico, y crear otros tipos de datos
*/
package main

import (
	"bytes"
	"fmt"
)

type mystruct struct { // esto es como si fuera una Class constructora
	nombre   string
	apellido string
}

// var miarrayint []int //al crear array hay que definir tanto el tipo como el tamaño en este caso le ponemos 2
var miarrayint [6]int
var miarraystring [6]string

var miarrayconindice map[string]string

func main() {

	/* es posible necesitar instanciar variables de otros paquetes*/
	var stdOut bytes.Buffer
	/* asi se instancia una struct y se agrega valor*/
	var mivar = mystruct{nombre: "Luis", apellido: "Martelo"}
	//como instanciar un struct vacio Nota no funciona porque no se esta usando
	var mivarvacia = mystruct{}
	miarrayint[0] = 97
	miarrayint[4] = 2
	miarraystring[0] = "texto"
	miarraystring[2] = "1"
	fmt.Printf("Soy la var 'mivar' con tipo: %T\n", mivar)
	fmt.Println("Soy la var 'mivar' y contengo el valor de:", mivar)
	fmt.Println("Soy la var 'mivarvacia':", mivarvacia)
	miarrayconindice = make(map[string]string) //asi instanciamos un map
	miarrayconindice["miindice"] = "mitexto"   /// falta instanciar
	fmt.Println("Soy la var 'mivarconndice' y contengo:", miarrayconindice)
	fmt.Println(miarrayint)
	fmt.Println(miarraystring)
	fmt.Printf("Soy la varible stdOut: %v, y soy de tipo: %T\n", stdOut, stdOut)
}
