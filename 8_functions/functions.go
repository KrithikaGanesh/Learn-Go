package main

import (
	"fmt"
)

func main() {
	sum := addAllValues(1, 2, 3, 4, 5)
	fmt.Println(sum)
}

func addAllValues(values ...int) int {
	sum := 0
	for _, n := range values {
		sum += n
	}
	return sum
}
