package main

/* EJEMPLO 1 HELLO WORLD
En ejemplo quiero que veas 2 cosas,
1 las zonas de programa hecho en go con el simple
ejemplo de un hola mundo.
2 como se crean comentarios dentro del codigo
*/
/*primera linea de package */

/*zona de importaciones*/
import "fmt"

/*
tambien si queremos importar uno o varios paquetes a utilizar podemos declararlos asi:
	import (
		"fmt"
		"pepito"
		"fulano"
	)
*/

/*FIN DE LA zona de importaciones*/

// asi se hace un comentario de una linea

/* asi se
hace un comentarios
de varias lineas
*/

/* funcion principal*/
func main() { //nombre de la funcion principal, simpre debe ser 'main'
	/* asi se imprimi en pantalla*/
	fmt.Println("Hello, World!")
}

/* debajo de la funcion main tambien puedes escribir codigo en los siguiente ejemplo te enseño como*/

/*
Para compilar nuestro archivo, en la terninal desde la ubicacion del archivo,
ejecutamos el comando 'go build main.go' , con esto se creará un archivo main.exe,
pero si lo que queremos es compilar y ejecutar nuestro archivo, debemos usar el
comando 'go run main.go'
*/
