package main

/*
Ejemplo punteros
En este ejemplo intento mostrar como usar de forma basica
los punteros
signo & -> permite obtener la direccion en memoria de las varibles
Ej:
	var x int = 5
	var p *int = &x		//varible p creada del tipo puntero (*) hacia una variable del tipo entero (int), apuntando a la direccion en memoria de la varible x
	ó
	p := &x 			//Varible p asignada e inicializada con el tipo puntero hacia la direccion en memoria de x
El signo * -> permite ver y editar lo almacenado en la direccion de memoria
Ej:
	fmt.Println(*p)     // imprime en console el valor de 5
	fmt.Println(p)		// imprime en consola la direccion de memoria de x
	*p = 300			// reasigna el contenido almacenado en la direccion de memoria p o sea el valor almacenado en x (ahora x valdrá 300)

como podemos ver la direccion de memoria de la variable original y
la direccion de memoria del puntero
*/

import "fmt"

func main() {
	mi_var := 100      //varible asignada al tipo int e inicializada con el valor 100
	puntero := &mi_var //puntero apunta a una variable

	fmt.Println("Elemento almacenado en 'mi_var':", mi_var)
	fmt.Println("Dirección en memoria de mi_var:", &mi_var)
	fmt.Println("Dirección en memoria que apunta el 'puntero':", puntero)
	fmt.Println("Elemento almacenado en memoria que apunta el 'puntero':", *puntero)
	fmt.Println("Dirección en memoria que ocupa el 'puntero':", &puntero)
	sumaVariableNormal(mi_var)
	fmt.Println("Elemento almacenado en 'mi_var':", mi_var) //mi_var  no ha cambiado

	//como cambiamos el valor de la variable?
	sumaPorReferencia(&mi_var)                              //mandamos la direccion en memoria de las varible mi_var
	fmt.Println("Elemento almacenado en 'mi_var':", mi_var) //mi_var ha cambiado

	//(no se recomienda hacer esto, ya que es dificil de leer a simple vista)
	sumaPorReferenciaunPuntero(&puntero)
	fmt.Println("Elemento almacenado en 'mi_var':", mi_var) //mi_var ha cambiado
}

func sumaVariableNormal(a int) {
	fmt.Println("           creamos la varible 'a'")
	a = a + 10
	fmt.Println("Suma Normal, a = ingresado + 10 =", a)
	fmt.Println("Dirección en memoria que ocupa la variable 'a':", &a)
}

func sumaPorReferencia(a *int) { //a es creada como un puntero
	fmt.Println("           creamos una nueva varible 'a'")
	*a = *a + 2 //le sumamos 10 al elemento almacenado en la direccion de memoria 'a'
	fmt.Println("Dirección en memoria que ocupa la variable 'a':", &a)
	fmt.Println("Dirección que almacena el puntero 'a':", a)
}

func sumaPorReferenciaunPuntero(a **int) { //significa que guardaremos la direccion en memoria del puntero que guarda la direccion en memoria de la varible mi_var
	fmt.Println("           creamos una nueva varible 'a'")
	**a = **a + 10 //a es otro puntero que apunta al 'puntero' que guarda mi_var y le sumremos 10 a ese elemento
	fmt.Println("Dirección que ocupa variable 'a': ", &a)
	fmt.Println("Dirección almacenada en 'a':", a) // es la misma que direccion en memoria del 'puntero'
}
