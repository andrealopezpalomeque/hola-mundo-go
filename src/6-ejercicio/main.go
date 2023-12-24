package main

import (
	"fmt"
)


func main() {
	//calculo area triangulo

	//ingresar longitud de triangulo rectangulo

	var base float64
	var altura float64

	fmt.Print("Ingresa la base del triangulo: ")
	fmt.Scanf("%f", &base)

	fmt.Print("Ingresa la altura del triangulo: ")
	fmt.Scanf("%f", &altura)

	//area con precision de 2 decimales
	area := (base * altura) / 2

	fmt.Printf("El area del triangulo es: %.2f\n", area)


}