package main

/*
Ejemplo struct
En este ejemplo intento mostrar struct con funciones

*/
import "fmt"

//https://repl.it/@steevehook/type-aliases#main.go

type human struct { //struct constructor
	name string
	age  int
}

func (h *human) setName(n string) { //creamos el metodo setName ligado a las instancias de human, h es un puntero que apunta a las variable que invoca el metodo
	h.name = n
	fmt.Println("line 19:", &h.name) //mostramos la direccion en memoria de la variable 'name'
}
func (h *human) setAge(a int) {
	h.age = a
}

//asi se puede copiar una structura y heredar sus metodos
type student = human

func info(h human) {
	fmt.Printf("Hi my name is: %s and I'm: %d\n", h.name, h.age)
}

func main() {
	h := human{name: "John", age: 25}
	luis := human{}
	luis.name = "pepe"
	s := student{}
	luis.setName("Luis")
	p := &luis.name
	fmt.Println("Linea 39, p:", p)
	luis.setAge(39)
	s.setName("Jane")
	s.setAge(23)
	info(h)
	info(s)
	fmt.Println(luis)
	info(luis)
}
