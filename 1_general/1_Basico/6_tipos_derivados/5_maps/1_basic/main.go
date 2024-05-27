package main

/*
Ejemplo map
En este ejemplo intento mostrar
como declara un map y revisar si esta nulo o no
*/
import "fmt"

func main() {
	var app map[string]int     //creamos un map pero al no estar inicializado (en esta etapa es nil)
	app = make(map[string]int) //así inicializamos el map (ya no es nil pero aun no tiene elementos almacenados)
	fmt.Println(app)
	if app == nil {
		fmt.Println("app variable is nil")
	} else {
		fmt.Println("app variable is non-nil")
	}
	app["first"] = 98 //asi agregamos elementos al map (respetando el tipo definidos inicialmente)
	fmt.Println(app)
}

/*
CONCLUSION: podemos crear map pero si o si para poder usarlos hay que
inicializarlo primero y luego ir agregando los elementos. tambien
podemos inicializarlo y agregar elementos en la creacion:
var app = map[string]int {
	"second":90,
	"cuatro":34,
}                                //En este caso no es necesario usar la funcion make
o (dentro de una funcion):
app := make(map[string]int)
*/
