package main

import (
	"fmt"
	"unsafe"
	"strconv"
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

	const myFirstConst = "a12"
	fmt.Println("My const is:", myFirstConst)

	const myIntConst int = 12
	fmt.Println("My int const is:", myIntConst)

	var myInt8 int8
	fmt.Printf("My 8 bit var: %d\n", myInt8)

	myInt8Bytes := unsafe.Sizeof(myInt8)
	myInt8Bits := unsafe.Sizeof(myInt8) * 8

	fmt.Printf("Type: %T, bits: %d, bytes: %d\n", myInt8, myInt8Bits, myInt8Bytes)

	var myOnlyPositiveInt uint = 13
	fmt.Printf("My OnlyPositiveInt var: %d\n", myOnlyPositiveInt)

	var myStringFromNumbers = fmt.Sprintf("my 8 bytes is %d and my 8 bits is %d", myInt8Bytes, myInt8Bits)
	fmt.Printf("My StringFromNumbers is: %s\n", myStringFromNumbers)

    intFromString, _ := strconv.ParseInt("1333", 0, 64)
	fmt.Printf("My intFromString var: %d\n", intFromString)
}