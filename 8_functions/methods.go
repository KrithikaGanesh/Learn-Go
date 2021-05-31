package main

import (
	"fmt"
)

func main() {
	emp1 := Employee{"Krithika", "Backend"}
	fmt.Println(emp1)
	emp1.team = "Distributed Systems"
	team := emp1.getEmployeeTeam()
	fmt.Println(team)
	fmt.Println(emp1)
}

type Employee struct {
	name string
	team string
}

func (e Employee) getEmployeeTeam() string {
	return e.team
}
