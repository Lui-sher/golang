package main

import "fmt"

type person struct {
}

type items struct {
}

func main() {
	// type & value MUST BE NIL
	p := adult(5)
	fmt.Println(p)
	switch p.(type) {

	case nil:
		fmt.Println("Es una variable tipo nil")
	case int:
		fmt.Println("Es una variable tipo int")
	case float64:
		fmt.Println("Es una variable tipo float64")
	case int64:
		fmt.Println("Es una variable tipo int64")
	case string:
		fmt.Println("Es una variable tipo string")
	case person:
		fmt.Println("Es una variable tipo person")
	default:
		fmt.Println("No es ninguno de los tipos anteriores")

	}
	/*
		if p != nil {
			fmt.Println(":(, I'm still a kid")
		} else {
			fmt.Println("He-he, finally got adult")
		}
	*/
}

func adult(n int) interface{} {
	//var res *person = nil
	res := person{}

	if n <= 0 {
		return nil // type: nil | value: nil
	}
	if n > 0 && n < 5 {
		return "String"
	}
	if n == 5 {
		return true
	}
	return res // type: *person | value: nil
}
