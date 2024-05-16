/*
Ejemplo de operadores aritmeticos
En este ejemplo intento mostrar 2 cosas; Como hace operaciones matematicas y a que me referia cuando decia que golang es fuertemente tipado
*/
package main

import "fmt"

func main() {

	var a, b, c int

	a = 1
	b = 2
	// ¿Qué tipo de variable fue asignado a "d"?
	d := a + b //Mi R. "d" se auto asignó como tipo int ya que el resultado de la suma 'a + b' es 3
	fmt.Printf("tipo de dato d %T\n", d)
	//¿"c" no tiene decimales?
	c = a / b //Mi R. c no tine decimales porque fue declarado con el tipo int, por ende solo guardará los enteros antes del punto decimal

	fmt.Printf("datos de variable d = %v\n", d)
	fmt.Printf("datos de variable a + b= %v\n", a+b)
	fmt.Printf("datos de variable c = %v\n", c)
	fmt.Println("-----------ahora con punto flotante ")

	var af, bf, cf float64

	af = 1.234
	bf = 2.6
	df := af + bf
	fmt.Printf("%T\n", df)
	//cf=a/bs
	cf = af / bf
	fmt.Printf("datos de variable df = %v\nsoy de tipo: %T\n", df, df)
	fmt.Printf("datos de variable af+bf = %v\nsoy de tipo: %T\n", af+bf, af+bf)
	fmt.Printf("datos de variable cf = %v\nsoy de tipo: %T\n", cf, cf)

	//¿Qué pasaria si "d" que es entero le asignamos una division de float64?
	// Mi R. Go no nos permite asignar valores de un tipo a otro tipo por lo tanto 'd = af / bf' nos marcaria un error
	//d = af / bf

	//Otro Ejempo de operaciones entre tipos de datos:
	//p := d / af  //ERROR: invalid operation: d / af (mismatched types int and float64)

}
