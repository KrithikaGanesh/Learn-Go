package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "Hello world!"
	writeToFile(content)
	defer readFile("./testfile.txt")
}

func writeToFile(content string) {
	file, err1 := os.Create("./testfile.txt")
	checkError(err1)
	length, err2 := io.WriteString(file, content)
	checkError(err2)
	fmt.Printf("%v chars written\n", length)
	defer file.Close()
}

func readFile(filename string) {
	data, err := ioutil.ReadFile(filename)
	checkError(err)
	fmt.Println("Read from file:", string(data))
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
