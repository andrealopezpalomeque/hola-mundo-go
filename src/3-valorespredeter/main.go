package main

import (
	"fmt"
)

func main() {
	var defaultInt int
	var defaultUint uint
	var defaultFloat float32
	var defaultString string
	var defaultBool bool

	fmt.Println(
		"int", defaultInt,
		"uint", defaultUint, 
		"float", defaultFloat, 
		"String", defaultString,
		"bool", defaultBool,
	)
}