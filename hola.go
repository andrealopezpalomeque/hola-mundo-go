package main

import (
	"fmt"
)

func main() {
	//fmt.Println("Hola Mundo")
	//fmt.Println(quote.Go())

	//? Declaracion de variables -> se pueden declarar fuera de la funcion main

	//var name string = "Juan" //esta variable almacena cadena de caracteres
	
	var firstName, lastName string 
	var age int

	//? Asignacion de variables
	firstName = "Juan"
	lastName = "Perez"
	age = 25

	//? Otra forma de declarar variables
	var (
		firstName2, lastName2 string
		age2 int
	)

	firstName2 = "oscar"
	lastName2 = "lopez"
	age2 = 30

	//? Otra forma de declarar variables
	var primer, segundo, edad = "robert", "rodriguez", 23

	fmt.Println(firstName, lastName, age)
	fmt.Println(firstName2, lastName2, age2)
	fmt.Println(primer, segundo, edad)

	// ! OTRA FORMA DE DECLARAR SIN LA PALABRA VAR
	//? := -> se usa para declarar variables dentro de una funcion
	//? = -> se usa para asignar valores a variables ya declaradas
	comida := "pizza" //? SOLAMENTE dentro de una funcion

	fmt.Println(comida)
	

}