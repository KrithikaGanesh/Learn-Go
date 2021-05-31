package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

const url = "http://services.explorecalifornia.org/json/tours.php"

func main() {

	resp, err := http.Get(url)
	checkError(err)
	defer resp.Body.Close()

	bytes, err := ioutil.ReadAll(resp.Body)
	checkError(err)

	content := string(bytes)
	tours := toursFromJson(content)
	var total int64 = 0
	for _, tour := range tours {
		price, _ := strconv.ParseInt(tour.Price, 10, 64)
		total += price
	}
	fmt.Println("Total price:", total)
}

func toursFromJson(content string) []Tour {
	tours := make([]Tour, 0, 20)
	decoder := json.NewDecoder(strings.NewReader(content))
	_, err := decoder.Token()
	checkError(err)

	var tour Tour
	for decoder.More() {
		err := decoder.Decode(&tour)
		checkError(err)
		tours = append(tours, tour)
	}
	return tours
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type Tour struct {
	Name, Price string
}
