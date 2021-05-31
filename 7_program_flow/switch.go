package main

import (
	"fmt"
)

func main() {
	switch value := 10; value {
	case 10:
		fmt.Println("Case 10")
		fallthrough // false through the next case as well
	default:
		fmt.Println("Not 10")
	}
}
