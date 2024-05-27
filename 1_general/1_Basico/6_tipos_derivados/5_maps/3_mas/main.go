package main

/*
Ejemplo map
En este ejemplo intento mostrar
trabajar estructuras con maps y maps con estructuras
*/
import (
	"fmt"
)

type Stats struct {
	cnt      int
	category map[string]Events
}

type Events struct {
	cnt   int
	event map[string]*Event
}

type Event struct {
	value int64
}

func main() {

	stats := new(Stats)
	stats.cnt = 33
	stats.category = make(map[string]Events)
	fmt.Println("1.:", stats)
	e, f := stats.category["aa"] //En Go al intentar acceder a una clave en un map este te permite usar otra varible para asignarle un valor booleano, para el manejo de errores, como en este caso: la clave ["aa"] no existe por ende a la varible 'f' se le asigna el valor de false, en caso contrario la varible 'e' tomaria el valor del elemento que contine la clave["aa"]
	fmt.Println(f, e)            //false
	if !f {
		fmt.Println("f es false!: ", f)

		e = Events{} //ya que no existe la clave "aa" creamos el valor primero
	}
	e.cnt = 66 //Almacenamos el int 66 en e.cnt

	e.event = make(map[string]*Event)           //inicializamos el map dentro de Events{}
	stats.category["aa"] = e                    //creamos la clave y valor que no existia anteriormente
	stats.category["aa"].event["bb"] = &Event{} //creamos inizializamos y creamos la clave "bb" (aun no le asignamos valor)
	stats.category["aa"].event["bb"].value = 99

	fmt.Println(stats)
	fmt.Println(stats.cnt, stats.category["aa"].event["bb"])
}
