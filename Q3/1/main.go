package main

import (
	"io"
	"os"
)

func main() {
	base, err := os.Open("file.txt")
	if err != nil {
		panic(err)
	}
	defer base.Close()

	target, err := os.Create("Newfile.txt")
	if err != nil {
		panic(err)
	}
	defer target.Close()

	io.Copy(target, base)
}
