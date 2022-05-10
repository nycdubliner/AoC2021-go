package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	// Hello World, the web server. 

	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	otherHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Some Other Text!\n")
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/other", otherHandler)
	log.Println("Listening for requests at http://localhost:8000/hello")
	log.Fatal(http.ListenAndServe(":8000", nil))
}