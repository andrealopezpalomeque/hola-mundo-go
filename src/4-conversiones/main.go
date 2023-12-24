package main

import (
	"fmt"
	"strconv"
)

func main() {
	// var integer16 int16 = 50
	// var integer32 int32 = 100

	//fmt.Println(integer16 + int16(integer32)) //convierto integer32 a int16
	//fmt.Println(int32(integer16) + integer32) //convierto integer16 a int32

	//convierto string a int
	 s:= "100"
	 i, _ := strconv.Atoi(s) //? devuelve un int y un error // _ es para ignorar el error

	 fmt.Println(i)

	//convierto int a string
	n := 42
	s = strconv.Itoa(n) //? devuelve un string
	fmt.Println(s + s) //obtengo una concatenacion de strings
	

	



	
	
}