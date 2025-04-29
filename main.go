package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
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

func toHtmlList(strs []string) string {
	var res string = "<ul>"
	for _, str := range strs {
		res += fmt.Sprintf("<li>%s</li>", str)
	}
	res += "</ul>"
	return res
}

func queryPg() string {
	// urlExample := "postgres://username:password@localhost:5432/database_name"
	conn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	rows, err := conn.Query(context.Background(), "select name from test")
	if err != nil {
		fmt.Fprintf(os.Stderr, "QueryRow failed: %v\n", err)
		os.Exit(1)
	}

	users := []string{}
	for rows.Next() {
		var name string
		rows.Scan(&name)
		users = append(users, name)
	}

	return toHtmlList(users)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "<h1>This is an example web app written in Go.</h1>")
	})

	http.HandleFunc("/text", func(w http.ResponseWriter, r *http.Request) {
		data := readFile("./text")
		fmt.Fprint(w, data)
	})

	http.HandleFunc("/pg", func(w http.ResponseWriter, r *http.Request) {
		data := queryPg()

		fmt.Fprint(w, fmt.Sprintf("<h1>Users in Postgres</h1>%s", data))
	})

	log.Println("The web app starts listening at port :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
