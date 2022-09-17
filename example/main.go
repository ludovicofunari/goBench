package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(w http.ResponseWriter, req *http.Request) {

	fmt.Fprintf(w, "Hi! I'm goBench! :)\n")
}

func main() {

	http.HandleFunc("/hello", hello)

	log.Println("Starting goBench...")
	http.ListenAndServe(":8080", nil)
}
