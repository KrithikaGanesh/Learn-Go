package main

import (
	"fmt"
)

func main() {
	emp1 := Employee{"Krithika", "Backend"}
	fmt.Println(emp1)
	fmt.Printf("Name:%v\n Team:%v\n", emp1.name, emp1.team)
	emp1.team = "Distributed systems"
	fmt.Printf("Name:%v\n Team:%v\n", emp1.name, emp1.team)
}

type Employee struct {
	name string
	team string
}
