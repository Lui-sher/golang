package main

/*
Ejemplo de struct
En este ejemplo intento mostrar
el trabajo con struct
*/
import "fmt"

type items struct { // En Go, struct lo podriamos ver como un Objeto en JS con clave: valor
	Type  string            // Definimos la clave 'Type' del tipo string
	Items map[string]string // Definimos otra clave 'Items' del tipo map*
}

var mapa = make(map[int]string) // varible 'mapa' en cuyo interior guarda un map definido con 'clave' de tipo 'int' y 'valor' de tipo string, se usa el 'make' para inicializar el 'map'

func main() {
	mapa[5] = "hola"                                 //	asi agregamos elementos al map
	fmt.Println(mapa)                                //
	var items_vacio = items{}                        // Creamos una nueva instancia de items con el nombre items_vacio
	fmt.Println("1. item_vacio:", items_vacio)       //
	items_vacio.Type = "text"                        // Agregamos valor a la clave Type
	fmt.Println("2. item_vacio:", items_vacio)       //
	items_vacio.Items = make(map[string]string)      // Para inicializar el 'map' de una plantilla debemos usar 'make'
	fmt.Println("3. item_vacio:", items_vacio)       //
	items_vacio.Items["item1"] = "item1tipo1"        //agregamos un elemento al map inicializado anteriormente
	fmt.Println("4. item_vacio:", items_vacio)       //
	i := make(map[string]string)                     //creamos e inicializamos la varible 'i' con un map
	fmt.Println("5. i:", i)                          //
	i["item1"] = "item1tipo1"                        //le agregamos elementos a 'i'
	fmt.Println("6. i:", i)                          //
	i["item2"] = "item1tipo1"                        //
	fmt.Println("7. i:", i)                          //
	i["item3"] = "item1tipo1"                        //
	fmt.Println("8. i:", i)                          //
	i["item1"] = "item1tipo1xxxxxxxxx"               //reescribimos un elemento ya creado en i
	fmt.Println("9. i:", i)                          //
	items_con_algo := items{Type: "tipo1", Items: i} // instanciamos el struct creado y lo inicializamos con los datos respectivos
	fmt.Println("items_vacio:", items_vacio)
	fmt.Println("items_con_algo:", items_con_algo)
}

/*
	CONCLUCION: los 'struct' son estructuras de datos que
	manejan 'clave' 'valor' en donde la clave es un nombre
	como si fuera una varible
	que identifica el elemento almacenado, este elemento
	puede ser definido con cualquier tipo de dato, mientras
	que los 'map' son estructuras con 'clave' 'valor' igualmente
	pero sus 'claves' puede ser definida de un solo tipo en toda
	la estructura al igual que el 'valor' solo puede almacenar
	un solo tipo de dato definido en su creacion

	Ej:

	type clase struct {
		Pepito      int
		Algo        string
		Lo_que_sea  func()
	}

	var mapa = make(map[int]string)
	y/o dentro de una funcion podemos usar la variante
	mapa := make(map[int]string)
*/
