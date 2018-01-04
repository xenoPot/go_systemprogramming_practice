package main

import (
	"encoding/csv"
	"os"
)

func main() {
	data := []string{"hoge", "hage", "bar", "fez"}

	writer := csv.NewWriter(os.Stdout)
	writer.Write(data)
	writer.Flush()
}
