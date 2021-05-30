package main

import (
	"fmt"
	"sort"
)

func main() {
	// string array initialization
	var countries [3]string
	countries[0], countries[1] = "India", "USA"
	fmt.Println(countries)

	// int array declaration
	var values = [3]int{3, 2, 1}
	fmt.Println(values)

	// int slice declarations
	var integers = []int{3, 2, 1}
	integers = append(integers, -1)
	fmt.Println(integers)

	integers = append(integers[2:])
	fmt.Println(integers)

	sort.Ints(integers)
	fmt.Println(integers)

}
