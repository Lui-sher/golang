package main

//ejemplo sacado de un video en internet
//https://repl.it/@steevehook/nil-vs-empty-slices#main.go
//https://www.youtube.com/watch?v=IB2JzpXaXo4
/*
muestra la diferencia entre structs y slices cuando estan vacios y nuloss
*/
import (
	"encoding/json"
	"fmt"
)

type people []struct { //creamos una plantilla slice[] que almacena elementos tipo "struct"(Objeto)
	N string `json:"name"` //al manejarlo como json tendrá la etiqueta name
	A string `json:"apellido"`
}

type response struct { // esta es otra plantilla struct(objeto)
	Items people `json:"items"` // los elementos que podrá almacenar la variable Items es del type people si o si
}
type variable []int //plantilla tipo variable slice[]

func main() {
	// slice zero value = nil
	p := people{ // p ahora es una instancia de people con sus respectivos valores
		{N: "John", A: "Frias"}, //construimos varias instancias de la misma plantilla dentro de p al usar {}
		{A: "Steve"},            //esto es igual a la linea de arriba
	}

	fmt.Println(p) //p es un slice(array) de elementos struct{Name}
	bs1, _ := json.Marshal(response{Items: p})
	bs2, _ := json.Marshal(response{})                                        // Items = null //instanciamos response, pero no lo inicializamos
	bs3, _ := json.Marshal(response{Items: people(nil)})                      // Items = null //instanciamos response pero lo nicializamos con nil, es lo mismo que la linea anterior
	bs4, _ := json.Marshal(response{Items: people{}})                         // Items = empty //instanciamos response y lo inicializamos people hollow
	bs5, _ := json.Marshal(response{Items: people{{N: "PEPE"}, {N: "LOLO"}}}) // Items = empty //instanciamos response y lo inicializamos con una instancia de people inicializado a su vez con Name: PEPE
	x := variable{1, 5, 8, 9}
	fmt.Printf("%s\n%s\n%s\n%s\n%s\n", string(bs1), string(bs2), string(bs3), string(bs4), string(bs5))
	fmt.Printf("Slice: %x tamaño: %v capacidad: %v", x, len(x), cap(x))
}

/*
CONCLUCION:
	los struct los podemos ver como Objetos.js
	los slice como array.js
	los map como array.js con clave(type): valor(type)
	los type son plantilla como constructores.js
	Al crear una plantilla debemos tener en cuenta la estructura que elegimos, de esta dependendera lo que deseemos
	Ej:
	type ejemplo []string        //plantilla de tipo slice con elementos tipo string

	type ejemmplo2 struct{
		Clave string
	}                            //plantilla tipo struct con valor de calve tipo string

	type ejemplo3 [] struct{
		Clave  string
		Clave2 int
		Clave3 bool
	}                            //plantilla tipo slice con elementos tipo struct -> con elementos valor de Clave tipo {string, int, bool respectivamente}

*/
