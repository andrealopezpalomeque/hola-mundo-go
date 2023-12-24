package main

import (
	"fmt"
)

func main() {


	name := "Golang"
	version := 1

	// %s para string
	// %d para int
	// %f para float
	// %v para cualquier tipo de dato (que no conocemos)
	// %T para conocer el tipo de dato

	fmt.Printf("Hola este lenguaje se llama %s y estoy utilizando la version %d\n", name, version)

	//! pedir que el usuario ingrese datos

	var name2 string
	var version2 int

	fmt.Print("Ingresa el nombre del lenguaje: ")
	fmt.Scanf("%s", &name2)

	fmt.Print("Ingresa la version del lenguaje: ")
	fmt.Scanf("%d", &version2)

	fmt.Printf("Hola este lenguaje se llama %s y estoy utilizando la version %d\n", name2, version2)
	fmt.Printf("El tipo de dato de version2 es: %T\n", version2)
}