package main

import (
	"fmt"
)

func main() {
	varInt := 5
	p := &varInt
	fmt.Println("Valur of p:", *p)
	// Changing pointer value will reflect in varInt
	*p = *p / 5
	fmt.Println("Valur of p:", *p)
	fmt.Println("Valur of varInt:", varInt)
}
