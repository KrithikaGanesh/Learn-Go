package main

import (
	"fmt"
)

// Global variable
const globalVar int = 1

func main() {
	// Global variable
	fmt.Printf("Value %v, Type %T\n", globalVar, globalVar)
	/*
	 * Declaring variable with type
	 */

	// integer
	var localIntVar int = 2
	fmt.Printf("Value %v, Type %T\n", localIntVar, localIntVar)

	// float
	var localFloatVar float64 = 2.35
	fmt.Printf("Value %v, Type %T\n", localFloatVar, localFloatVar)

	// string
	var localStringVar string = "Hello World!"
	fmt.Printf("Value %v, Type %T\n", localStringVar, localStringVar)

	/*
	 * Declaring variable without type
	 */

	// int
	localIntVar2 := 2
	fmt.Printf("Value %v, Type %T\n", localIntVar2, localIntVar)

	// float
	localFloatVar2 := 2.35
	fmt.Printf("Value %v, Type %T\n", localFloatVar2, localFloatVar2)

	// string
	localStringVar2 := "Hello World!"
	fmt.Printf("Value %v, Type %T\n", localStringVar2, localStringVar2)

}
