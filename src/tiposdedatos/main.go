package main

import (
	"fmt"
	"math"
)

func main() {
	var integer8 int8 = 127
	var integer uint = 100
	var float float32 = 3.14 //32 numeros pequeños, 64 numeros grandes
	
	

	fmt.Println(integer)
	fmt.Println(integer8)
	fmt.Println(float)
	//para saber los valores minimos y maximos de los tipos de datos -> uso el paquete math
	fmt.Println(math.MinInt64, math.MaxInt64)

	//! Datos BOOLEANOS
	var valueBoolean bool = true // false

	fmt.Println(valueBoolean)

	//! caracteres ESPECIALES
	fullName := "Alex Lopez \t(alias \"pipi\") \n" // \n -> salto de linea, \t -> tabulacion, 
	fmt.Println(fullName)

	//! BYTE -> es un alias para uint8 -> numeros sin signo de 8 bits
	
	var a byte = 'a'
	fmt.Println(a) //imprime el valor de codigo ascii de la letra a, que es 97

	s:= "hola"
	fmt.Println(s[0]) // imprime el valor de codigo ascii de la letra h, que es 104
	
	//! RUNE -> es un alias para int32 y representa un punto de codigo unicode
	var r rune = '♥' // para asignar un valor de tipo rune, se debe usar comillas simples
	fmt.Println(r) // imprime el valor de codigo unicode del corazon, que es 9829
	
} 