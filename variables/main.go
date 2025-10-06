package main

import (
	"fmt"
)

func main() {
	var myIntVar int = 23
	fmt.Println("My variable is:", myIntVar)

	var myPositiveIntVar uint = 32
	fmt.Println("My positive variable is:", myPositiveIntVar)

	var myStringVar string = "My string variable"
	fmt.Println("My string variable is:", myStringVar)

	var myBooleanVar bool = true
	fmt.Println("My boolean variable is:", myBooleanVar)

	fmt.Println("My boolean variable address (pointer) is:", &myBooleanVar)

	myIntVar2 := 12
	fmt.Println("My dynamic int typed variable is:", myIntVar2)

	myStringVar2 := "test"
	fmt.Println("My dynamic string typed variable is:", myStringVar2)

}