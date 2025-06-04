package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

// add takes two integers and returns their sum.
func add(a int, b int) int {
	return a + b
}

// readFile reads the content of a file and returns it as a string.
func readFile(filename string) string {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatalf("Failed to read file: %s", err)
	}
	return string(data)
}

func main() {
	// Example usage of add
	result := add(5, 7)
	fmt.Printf("The sum is: %d\n", result)

	// Example usage of readFile
	fileContent := readFile("example.txt")
	fmt.Println("File content:")
	fmt.Println(fileContent)
}
