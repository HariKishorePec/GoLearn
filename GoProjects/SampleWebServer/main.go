package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Book struct {
	Name   string `json:"book_name"`
	Author string `json:"author"`
	Pages  int    `json:"total_pages"`
}

func main() {
	port := ":5000"
	fmt.Println("Second web server...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello from server")
	})

	http.HandleFunc("/about", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Server - About Page")
	})

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		fmt.Fprintln(w, "<h1> Home </h1> <p>Hi There!! </p>")
	})

	http.HandleFunc("/books", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		book1 := Book{
			Name:   "The Psychology of Money",
			Author: "Morgan Housel",
			Pages:  241,
		}

		// res, _ := json.Marshal(book1)
		// fmt.Fprintln(w, string(res))

		json.NewEncoder(w).Encode(book1)
	})

	fmt.Printf("Starting server at port: %s\n", port)
	log.Fatalln(http.ListenAndServe(port, nil))

}
