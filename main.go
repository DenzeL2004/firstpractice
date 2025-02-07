package main

import "fmt"

func main() {
	intVar := 42
	fmt.Printf("Value int variable: %d. Type: %T\n", intVar, intVar)

	floatVar := 3.14
	fmt.Printf("Value float variable: %g. Type: %T\n", floatVar, floatVar)

	boolVar := true
	fmt.Printf("Value bool variable: %t. Type: %T\n", boolVar, boolVar)

	strVar := "forty two"
	fmt.Printf("Value string variable: %s. Type: %T\n", strVar, strVar)
}
