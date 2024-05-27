package main

import (
	"fmt"
	"os"
)

func miPrimeraFuncion() {
	fmt.Println("mi funcion")
	fmt.Println("mi funcion")
	fmt.Println("mi funcion")
	fmt.Println("mi funcion")
	fmt.Println("mi funcion")
	fmt.Println("mi funcion")
}

func funcionConParametros(a string) { //definimos el parametro de entrada
	fmt.Println("mi funcion: ", a)
}
func funcionConParametrosYreturn(a int, b int) int { //definimos los parametros de entrada (a, b) y de salida (tipo int)
	return (a + b)
}

func funcionConmultiplesParametros(arg ...string) []string { //definimos parametro de entrada (copia tipo string) salida tipo array con elementos tipo string
	return arg
}
func funcionMultipleReturn(arg ...string) ([]string, string, string) { //definimos entrada  y multiples salidas
	return arg, "err", "xxxx"
}

func main() {
	s := funcionConmultiplesParametros("1", "2", "3", "4", "5", "6")
	a, b, c := funcionMultipleReturn("1", "2", "3", "4")
	d, _, _ := funcionMultipleReturn("1", "2", "3", "4") //_ permite ignorar salidas de una funcion

	fmt.Println(s)
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	fmt.Println("vamos a usar una funcion del sistema")
	os.Exit(2) //permite detener la ejecucion del codigo justo en este punto, por lo tanto nada de lo que esté escrito a continuacion se ejecutará

	func() {
		fmt.Println("go routine 1 is done")
	}()

}
