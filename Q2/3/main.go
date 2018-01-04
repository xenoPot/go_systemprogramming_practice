package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	// json化する元のデータ
	source := map[string]string{
		"Hello": "World",
	}
	// 解答処理
	compresser := gzip.NewWriter(w)
	multiWriter := io.MultiWriter(compresser, os.Stdout)
	encoder := json.NewEncoder(multiWriter)
	encoder.SetIndent("", "  ")
	encoder.Encode(source)
	compresser.Flush()
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
