package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func readFile(path string) string {
	file, err := os.Open(path)
	if err != nil {
		log.Printf("Error: %s", err)
		return ""
	}

	b, err := io.ReadAll(file)
	if err != nil {
		log.Printf("Error: %s", err)
		return ""
	}

	return string(b)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>This is an example web app written in Go.</h1>")
	})

	http.HandleFunc("/text", func(w http.ResponseWriter, r *http.Request) {
		data := readFile("./text")
		fmt.Fprint(w, data)
	})

	log.Println("The web app starts listening at port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
