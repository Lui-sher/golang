/*
Ejemplo ciclos de for
En este ejemplo intento mostrar muchas formas de usar el "for"
*/
package main

import "fmt"

func main() {
	var i int //cuando se declara una varible esta inicia en cero por defecto
	var j string
	fmt.Printf("soy i: %d, y yo soy j: %q\n", i, j) //valor de i = 0 y j = ""
	//NOTA: j es un string en blanco por eso en la salida no vemos el elemento que contiene cuando usamos fmt.Println(j)

	fmt.Println("-------------Primer ejemplo-----------------")
	for i <= 10 {
		fmt.Println("Valor de i:", i)
		i++
	}
	fmt.Println("-------------Segundo ejemplo-----------------")
	i = 0
	for ; i < 10; i++ {
		fmt.Println("Valor de i:", i)
	}

	fmt.Println("-------------Tercer ejemplo-----------------")

	for i = 0; i < 10; i++ {
		fmt.Println("Valor de i:", i)
	}

	fmt.Println("-------------Cuarto ejemplo-----------------")

	for i = 0; i < 10; {
		fmt.Printf("Valor de i: %d", i)
		if i == 6 {
			fmt.Printf(" sumaremos 3...\n")
			i = i + 3
			continue //ejemplo para ver el continue
			//Nota: El continue hace que la ejecucion salte al siguiente ciclo (OJO: NO CANCELA for completo )
		}
		fmt.Printf("...\n")
		i++

	}
	fmt.Println("-------------Quinto ejemplo-----------------")
	for i = 0; i < 10; {
		fmt.Printf("Valor de i: %d", i)
		if i == 6 {
			fmt.Printf(" sumaremos 3\n")
			i = i + 3
			break //ejemplo para ver el break
			// El break CANCELA toda la ejecucion del for
		}
		fmt.Printf("...\n")
		i++
	}

	fmt.Println("-------------Sexto ejemplo-----------------")

	for i := preFor(); condicion(i); i = postFor(i) {
		//1. le asignamos a la varible 'i' la funcion preFor() por ende, 'i' ahora es la funcion preFor
		//2. Ejecutamos la funcion condicion(i) pasandole la varible i creada en el paso 1 es lo mismo que hacer => condicion(preFor())
		// 	2.1 En preFor(): se imprime "prefor i" y retorna el entero 0 lo equivalente => condicion(0)
		//  2.2 En condicion(0): se imprime "condicion i" y se retorna el booleano 'true' (i < 10 => i vale 0 => 0 < 10 => true ) lo que corresponde a evaluar la condicion (si es true se ejecuta el codigo, else se termina el ciclo for)
		// 3. al ser true, se ejecuta el codigo dentro del for

		fmt.Println("Ejecutamos el ciclo: ", i)
		if i == 7 {
			fmt.Printf(" En i = 7 se activa un condicional if y ejecutamos el comando 'break', forzando la terminacion del FOR...\n")
			break /// este ejemplo es para usar el break
		}

		// 4. Se ejecuta el incremento de 'i', en este caso es la funcion postFor(i) => postFor(preFor())
		//   4.1 En postFor(preFor()): igual que en el paso 2.1 => postFor(0)
		//   4.2 En postFor(0): se imprime "postFor sumemos i" then -> 0++ = 1 then -> retorna el entero 1
		// 5. En i = postFor(i) => i = 1
		// 6. Se repite el ciclo continuando desde el paso 2 que evalua la condicion
	}

}
func preFor() int {
	fmt.Println("Inicializamos en preFor() la varible i =", 0)
	return 0
}
func postFor(i int) int {
	i++
	fmt.Println("postFor sumemos 1 a i y continuemos con el siguente ciclo i = ", i)
	fmt.Println()
	return i
}
func condicion(i int) bool {
	boo := i < 10
	fmt.Printf("Iniciamos el ciclo evaluando la condicion (%d < 10) -> %v\n", i, boo)
	return (i < 10)
}
