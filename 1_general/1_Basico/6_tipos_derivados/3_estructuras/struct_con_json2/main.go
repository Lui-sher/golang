package main

/*
Ejemplo struct
En este ejemplo intento mostrar

en este ejercicio se le pide al estudiante que busque y entienda
el ejercicio
porque se usan en las notaciones de json valores como "-", omitempty
o porque se declara tipos de datos como string solo en unos de los items
 ademas ese item es float64 (doble asignacion de tipo de datos?)
*/
//https://repl.it/@steevehook/json-struct-tags#main.go
//struct con notas de json s
import (
	"encoding/json"
	"fmt"
)

type person struct {
	Name string `json:"name"`
	//aqui un ejemplo del uso de omitempty s
	Hobby string `json:"hobby,omitempty"` // ("omitempty") omite esta etiqueta si está vacia
	Email string `json:"-"`               // ("-") se omite esta etiqueta siempre, esté o no vacia
	//ejemplo de doble asignacion de tipos
	Money float64 `json:"money,string"` // ("string") en formato json simpre será string
}

func main() {
	p1 := person{Name: "John", Hobby: "Music"}
	bs1, _ := json.Marshal(p1)
	fmt.Println(string(bs1))

	p2 := person{Name: "Jane", Email: "e@d.com", Money: 5.88}
	bs2, _ := json.Marshal(p2)
	fmt.Println(string(bs2))

	p3 := person{Name: "Steve", Money: 200}
	bs3, _ := json.Marshal(p3)
	fmt.Println(string(bs3))
}
