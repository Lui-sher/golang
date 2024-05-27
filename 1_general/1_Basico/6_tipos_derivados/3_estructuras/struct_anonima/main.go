package main

/*
Ejemplo struct
En este ejemplo intento mostrar como trabajar
con estructuras anomimas
*/
//https://repl.it/@steevehook/anonymous-structs#main.go
import (
	"encoding/json" //  Se utiliza para trabajar con datos en formato JSON.
	"fmt"
)

func main() {
	//ejemplo de un struct anomino
	res := struct { // llamamos a un 'struct' anonimo cuando no lo creamos como plantilla para instanciar su modelo
		Name string `json:"name_of"` //etiqueta para el manejo del json
		Age  int    `json:"age_of"`
	}{ //con su declaracion de datosss
		Name: "Steve",
		Age:  25,
	}
	fmt.Println(res)
	res.Name = "Luis"               // reasignamos el valor de Nombre del struct res
	bs, _ := json.Marshal(res)      // Al usar ( _ ) le estamos diciendo al compilador que ignore el segundo resultado de la funcion Marshal()
	bs2, error := json.Marshal(res) //asi seria de forma normal
	/*
		Marshal es una funcion del paquete json que
		retorna dos valores, el primero es un slice que
		representa el contenido de la varible
		ingresada, en bytes y el segundo el error en
		caso de que lo halla
		json.Marshal(v) => ([]byte, error)
		Marshal() retorna una copia del struct en formato json
	*/
	fmt.Println(bs) // Imprime el resultado en formato byte
	fmt.Println(bs2, error)
	fmt.Println(string(bs)) //Imprime el resultado de byte en formato string
}
