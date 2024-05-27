package main

/*
Ejemplo de punteros
En este ejemplo intento mostrar
como usar puntero con funciones
es el inicio de la programacion orientada a objeto que
se puede crear en golang
*/
import "fmt"

type user []string //En Go podemos crear nuestras propias plantillas usando la palabra reservada 'type', en este caso creamos una plantilla tipo slice(Array) cuyos elementos seran del tipo string

//Ej:
//var luis user{"luis","martelo"} //declaramos una varible de tipo user y lo asignamos los elementos de tipo string

func (p user) agregar(name string) { //Asi creamos un metodo ligado a un tipo de dato, en este caso al tipo 'user', en el que nos pedirá ingresar un tipo de dato 'string'
	luis := user{} //prueba personal, no hace parte del ejercicio en si
	fmt.Println(luis)
	p = append(p, name)
	fmt.Println("Line 21:", p)
}

func add(p *user, s string) { //Asi es como creamos una funcion comun Go, en este caso nos pide ingresar dos tipos de datos distintos uno del tipo '*user' y otro del tipo 'string' (*user se refiere a que trabajaremos con la direccion en memoria de una varible de tipo 'user')
	*p = append(*p, s)
	fmt.Println(*p)
}

func (p *user) addPtr(name string) {
	*p = append(*p, name)
}

func main() {
	/*
		Este tipo de declaracion de varible ':=' solo funciona dentro de alguna
		funcion, usar el metodo 'var x int = 9' por fuera instead
	*/
	luis := user{"luis"}            //varible tipo user normal
	fmt.Println(luis)               //[luis]
	world := &user{"John", "Steve"} //world es un puntero que apunta a la direccion en memoria que almacena el valor [John Steve] podriamos decir que ese valor no es representado por variable alguna
	add(&luis, "mart")              //Don´t work, it only create a copy of the var
	//&luis.addPtr("martelo") //asi no funciona, debemos crear otra varible para acceder a la dirreccion en memoria de la varible
	mLuis := &luis //creamos una varible que almacena la ubicacion en memoria de la varible 'luis'
	fmt.Println(mLuis)
	luis.agregar("hernandez")
	fmt.Printf("%v\n%v\n", world, luis)
	world.addPtr("Mike") // worked
	fmt.Println(world)
}

/*CONCLUCION: todas las funciones que reciben argumentos crean una copia de estos
y trabajan en tales copias, NO modifican la varible original pasada como argumento,
si queremos modificar la varible original debemos pasarle a la funcion el numero de
la ubicaion en memoria de la varible y trabajar con ese parametro, como en el metodo addPtr()
*/
