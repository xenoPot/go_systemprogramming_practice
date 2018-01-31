package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	newFile, createErr := os.Create("Newfile")
	if createErr != nil {
		panic(createErr)
	}
	defer newFile.Close()

	buffer := make([]byte, 1024)
	_, readErr := io.ReadFull(rand.Reader, buffer)
	if readErr != nil {
		panic(readErr)
	}

	newFile.Write(buffer)
}
