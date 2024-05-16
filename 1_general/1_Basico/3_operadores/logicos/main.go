/*
Ejemplo de operadores logico
En este ejemplo
se muestra los operadores logicos y operaciones de corridas de bit
*/

package main

import "fmt"

func main() {
	x := true  //bool
	y := false //bool
	fmt.Println(x || y)
	fmt.Println(x && y)
	fmt.Println(x && !y)
	fmt.Println(0 ^ 5)
	fmt.Println("-----------Corridas de bits ")
	fmt.Println(1 << 1) //10 ->
	fmt.Println(1 << 2) //100
	fmt.Println(1 << 3) //1000
	fmt.Println(1 << 4) //10000
	// todos los numeros iniciales estan expresados en decimal por ende 1 dec es convertido a binario y luego le agregamos el numero de '0's  indicados

	fmt.Println("-----------corridas de bits ¿por que esta mal?") //porque la entradra y la salida es en decimal, la conversion o binario solo ocurre internamente

	fmt.Println(10 >> 1)    // 10 = 1010(bin) -> ( >> 1 bit) -> 101(bin) -> 5
	fmt.Println(100 >> 2)   //100 = 1100100(bin) -> ( >> 2 bit) -> 11001(bin) = 25
	fmt.Println(1000 >> 3)  //125
	fmt.Println(10000 >> 4) //625
	//10 en binario en '1010' luego le restamos un '0' a la derecha y quedaria '101' que en decimal es le numero '5'

	fmt.Println("-----------corridas de bits ")
	fmt.Println(2 >> 1)  // 2 = 10(bin) -> ( >> 1 bit) -> 1(bin) = 1
	fmt.Println(4 >> 2)  // 4 = 100(bin) -> ( >> 2 bit) -> 1(bin) = 1
	fmt.Println(8 >> 3)  // 1
	fmt.Println(16 >> 4) // 1

	fmt.Println("-----------que numero necesitamos para que la salida sea 1? ")

	fmt.Println(1 >> 1) // 2
	fmt.Println(1 >> 2) //4
	fmt.Println(1 >> 3) //8
	fmt.Println(1 >> 4) //16

	////////////////////////////////////////
	/*  16|8|4|2|1
	    0|0|0|0|0 =  0+0+0+0+0 = 0
	    0|0|0|0|1 =  0+0+0+0+1 = 1
	    0|0|0|1|0 =  0+0+0+2+0 = 2
	    0|0|0|1|1 =  0+0+0+2+1 = 3
	    0|0|1|0|0 =  0+0+4+0+0 = 4
	    0|0|1|0|1 =  0+0+4+0+1 = 5
	    0|0|1|1|0 =  0+0+4+1+0 = 6
		0|0|1|1|1 =  0+0+4+2+1 = 7
		0|1|0|1|0 =  0+8+0+2+0 = 10


	    1|1|0|0|0 =  16+8+0+0+0 = 24
	         .
	         .
	         .
	    1|0|0|0|0 =  16+0+0+0+0 = 16
	         .
	         .
	         .
	    1|1|1|1|1 =  16+8+4+2+1 =31

	*/
}
