package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a string: ")
	// always taken in a string value
	strInput, _ := reader.ReadString('\n')
	fmt.Println("Value:", strInput)

	// To take in integer
	fmt.Print("Enter a int: ")
	input, _ := reader.ReadString('\n')
	// ParseInt takes in string, base, bitSize
	intInput, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value: ", intInput)
	}

	// To take in float
	fmt.Print("Enter a float: ")
	input2, _ := reader.ReadString('\n')
	// ParseInt takes in string, base, bitSize
	floatInput, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Value: ", floatInput)
	}
}
