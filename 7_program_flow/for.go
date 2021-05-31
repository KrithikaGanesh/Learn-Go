package main

import (
	"fmt"
)

func main() {
	var countries = []string{"India", "US", "Singapore"}
	fmt.Println(countries)

	for i := range countries {
		fmt.Println(countries[i])
	}

	for _, country := range countries {
		fmt.Println(country)
	}

	for i := 0; i < len(countries); i++ {
		fmt.Println(countries[i])
	}

	// there is no while, for can be used like while
	count := 0
	for count < 10 {
		fmt.Println(count)
		count++
	}

	for count < 1000 {
		fmt.Println(count)
		count++
		if count > 15 {
			goto end
		}
	}
end:
	fmt.Println("Done")
}
