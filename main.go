package main

import "fmt"

func main() {
	var intVar int = 42
	fmt.Printf("Value int variable: %d. Type: %T\n", intVar, intVar)

	var floatVar float64 = 3.14
	fmt.Printf("Value float variable: %g. Type: %T\n", floatVar, floatVar)

	var boolVar bool = true
	fmt.Printf("Value bool variable: %t. Type: %T\n", boolVar, boolVar)

	var strVar string = "forty two"
	fmt.Printf("Value string variable: %s. Type: %T\n", strVar, strVar)
}
