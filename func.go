package main

import (
	"fmt"
)

// declare a type of the variable after the variable
// the return type must be written.

// func the name of the function (argument type, ...) return type{
	
//}

func add(x int, y int) int {
	return x + y
}

func sub(x, y int) int {
	return x - y
}

// func the name of the function(argument type ...) (returned type 1, returened type2){
	// return value1, value2 	
//}

func addAndSub(x, y int) (int, int) {
	return x + y, x - y
}

func main() {
	fmt.Println(add(100, 200))
	fmt.Println(sub(100, 200))
	fmt.Println(addAndSub(100, 200))
}