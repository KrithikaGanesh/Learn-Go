package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var result string
	if val := input(); val < 0 {
		result = "Less than 0"
	} else if val > 0 {
		result = "Greater than 0"
	} else {
		result = "Equal to 0"
	}
	fmt.Print(result)
}

func input() int64 {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a int: ")
	input, _ := reader.ReadString('\n')
	// ParseInt takes in string, base, bitSize
	intInput, err := strconv.ParseInt(strings.TrimSpace(input), 10, 64)
	if err != nil {
		panic(err)
	}
	return intInput
}
