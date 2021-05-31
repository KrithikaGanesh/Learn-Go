package main

import (
	"fmt"
	"sort"
)

func main() {
	state_capitals := make(map[string]string)
	fmt.Println(state_capitals)
	state_capitals["California"] = "Sacremento"
	state_capitals["Arizona"] = "Phoenix"
	state_capitals["Colorado"] = "Denver"

	for k, v := range state_capitals {
		fmt.Printf("%v: %v\n", k, v)
	}

	states := make([]string, len(state_capitals))

	i := 0
	for k := range state_capitals {
		states[i] = k
		i += 1
	}

	fmt.Println(states)
	sort.Strings(states)
	fmt.Println(states)

	for i := range states {
		fmt.Println(states[i])
		fmt.Println(state_capitals[states[i]])
	}

}
