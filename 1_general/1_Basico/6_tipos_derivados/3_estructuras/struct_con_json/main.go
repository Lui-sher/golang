package main

/*
Ejemplo struct
En este ejemplo intento mostrar struct con json
y como usar funciones

recuerda que este ejemplo no fue explicado en los videos
porque se intenta que lo analices y te preguntes porque pasan cosas
ejemplo
¿por que en algunos casos la estructura pasa por referencia y en otros no?
¿por que en algunos casos la estructura es retornada y en otros no?
¿puedes mejorar esta estructura?

la estructura esta hecha para guardar items de diferentes formatos
donde el type guarda que item es y el items guarda la estructura nueva
*/
import (
	"encoding/json" //package para trabajar con json
	"fmt"
)

type items struct {
	Type  string            `json:"type"` //etiqueta que le agreaga el string "type" a los elementos que son creados con este metodo
	Items map[string]string `json:"item"` // con esto estamos creando una structura con datos tipo json
}

func itemToken(token string) items { //entrada de la funcion recibe un "string" y retorna un tipo "items"
	i := make(map[string]string)          //inicializamos un mapa vacio
	i["secret"] = token                   //creamos em el mapa una clave = "secret" con el valor ingresado en la funcion token
	res := items{Type: "token", Items: i} // instanciamos la struct items con el type "token" e Items con el mapa creado en la linea anterior
	fmt.Println("soy itemToken:", res)
	return res
}

func itemCredencial(user string, pass string) items {
	i := make(map[string]string)
	i["user"] = user
	i["secret"] = pass
	res := items{Type: "credential", Items: i}
	return res
}

func itemAmazonCredencial(accountid string, user string, pass string) items {
	i := make(map[string]string)
	i["account"] = accountid
	i["user"] = user
	i["secret"] = pass
	res := items{Type: "credentialAmazon", Items: i}
	return res
}

func jsoncode(item items) []byte { //recibe un elemento tipo "items", y retorna un elemento tipo "[]byte"
	bs1, _ := json.Marshal(item)                                     // convertimos el elemento recibido en tipo "[]byte" (serializar el json)
	fmt.Println("soy jsoncode() -> type bs1(string): ", string(bs1)) // imprimimos el tipo "[]byte" pero en formato "string"
	return bs1                                                       // retornamos el elemento tipo "[]byte"
}

func jsondecode(item []byte) items { // recibe un elemento tipo "[]byte" retorna otro tipo "items"
	var x items              // instansiamos el struct tipo "items"
	json.Unmarshal(item, &x) // Unmarshal(arg, arg2) recibe dos argumentos arg = "[]byte" arg2= struct"items", (deserializamos el json)
	// deserializamos el item []byte y lo agregamos a la variable x, si todo es correcto retornará un error=<nil>
	fmt.Printf("soy jsondecode() -> x: %v, deserializado de item recibido\n", x)
	return x
}

func main() {

	x := jsondecode(jsoncode(itemToken("pass")))
	fmt.Print("\ntipo\n")
	fmt.Print(x.Type)
	fmt.Print("\nitem\n")
	fmt.Println(x.Items["secret"])
	jsondecode(jsoncode(itemCredencial("Luis", "1234")))
	jsondecode(jsoncode(itemAmazonCredencial("ABCD", "martelo", "12345")))
}
