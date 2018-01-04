package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintf(file, "%d, %s, %f", 1000, "Test", 100.0)
}
