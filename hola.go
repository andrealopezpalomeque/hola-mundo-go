package main

import (
	"fmt"
)

//! Declaracion de CONSTANTES ------------------------------------------------------------------------
const Pi float32 = 3.14 // hay que declarar e inicializar la constante, despues no se puede cambiar su valor

const (
	x = 100
	y = 0b1010 // binario
	z = 0o12 // octal
	w = 0xFF // hexadecimal
)

const (
	Domingo = iota + 1 // iota -> genera una secuencia de numeros, inicia en 0
	Lunes
	Martes
	Miercoles
	Jueves
	Viernes
	Sabado
)

func main() {
	//fmt.Println("Hola Mundo")
	//fmt.Println(quote.Go())
	// ! DECLARACION DE VARIABLES ------------------------------------------------------------------------

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
	
	//! Declaracion de CONSTANTES ------------------------------------------------------------------------
	fmt.Println(Pi, x, y, z, w)
	fmt.Println(Domingo, Lunes, Martes, Miercoles, Jueves, Viernes, Sabado)
		
}