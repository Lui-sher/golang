package main

/*
Ejemplo slice
En este ejemplo intento mostrar
como trabajar con slice
*/
import (
	"fmt"
	"strings"
)

func main() {
	// For slice[ i : j : k ] the
	// Length:   j - i
	// Capacity: k - i

	people := []string{ //slice tipo string
		"John",
		"Steve",
		"Jane",
	} //len -> 3, cap -> 3

	slice := make([]string, 0, 5)
	slice = append(slice, "Juan", "Esteban", "Juana")
	fmt.Printf("slice tamaño: %v cap: %v  \n", len(slice), cap(slice))
	inspect("slice: ", slice)
	slice = append(slice, "Pedro", "Pablo")
	inspect("slice: ", slice)
	fmt.Printf("slice tamaño: %v cap: %v  \n", len(slice), cap(slice))
	fmt.Println("----------------desborde de la capacidad del slice-------------------")
	//desbordando memoria
	slice = append(slice, "Elephant")
	inspect("slice: ", slice)
	fmt.Printf("slice tamaño: %v cap: %v  \n", len(slice), cap(slice))

	array := append(slice[0:0:0], slice...)
	inspect("Direccion en memoria de los items de array: ", array)

	inspect("Direccion en memoria de los items de people: ", people)
	sliceOfPeople := people[0:]
	inspect("Direccion en memoria de los items de sliceOfPeople: ", sliceOfPeople) // los elementos son los originales(misma direccion en memoria)
	peopleClone := append(people[:0:0], people...)
	inspect("Direccion en memoria de los items de peopleClone: ", peopleClone) //los elementos han sido copiados(otra direccion en memoria)

	// Crear un slice con capacidad inicial de 3
	original := make([]int, 3)
	original[0] = 1
	original[1] = 2
	original[2] = 3

	// Guardar una referencia al array original
	originalArray := original

	// Mostrar la dirección de memoria de los elementos originales
	fmt.Printf("Original[0]: %p\n", &original[0])
	fmt.Printf("Original[1]: %p\n", &original[1])
	fmt.Printf("Original[2]: %p\n", &original[2])

	// Agregar un nuevo elemento al slice, causando una reubicación del array subyacente
	expanded := append(original, 4)

	// Mostrar la dirección de memoria de los elementos después de append
	fmt.Printf("Expanded[0]: %p\n", &expanded[0])
	fmt.Printf("Expanded[1]: %p\n", &expanded[1])
	fmt.Printf("Expanded[2]: %p\n", &expanded[2])
	fmt.Printf("Expanded[3]: %p\n", &expanded[3])

	// Modificar el nuevo elemento en el slice expandido
	expanded[3] = 42

	fmt.Println("Original slice:", original) // Output: [1 2 3]
	fmt.Println("Expanded slice:", expanded) // Output: [1 2 3 42]

	// Mostrar el array original para confirmar que sigue existiendo con las mismas direcciones
	fmt.Printf("OriginalArray[0]: %p\n", &originalArray[0])
	fmt.Printf("OriginalArray[1]: %p\n", &originalArray[1])
	fmt.Printf("OriginalArray[2]: %p\n", &originalArray[2])

	fmt.Println("Array original:", originalArray) // Output: [1 2 3]

}

func inspect(label string, people []string) {
	f := strings.Repeat("%p | ", len(people))
	f = label + f + "\n"
	var args []interface{}
	for i := range people {
		args = append(args, &people[i])
	}
	fmt.Printf(f, args...)
}

/*
CONCLUSION:
los slice se crean apartir de arrays originales, le podemos definir los elementos, y la capacidad mediante la
ecuacion:
		slice[ i : j : k ]
		donde: Length:   j - i
		y Capacity: k - i

	La capacidad de un array define la cantidad de elementos maximo que puede tener un slice, en el momento en que
se agregen mas elementos que superen la capacidad se generá un nuevo slice con mas capacidad, a lo que conlleva
que cambien la direccion en memeoria de todos los elementos dentro del array es como si se clonara

	Los slice creados apartir de otros conservaran la misma direccion en memoria de todos sus elementos.
Al usar el operador (...) podemos clonar un array lo que significa que los elementos del nuevo array no tendran la misma
direccion en memoria que el original
*/
