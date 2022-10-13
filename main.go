package main

import (
	"io"
	"log"
	"net/http"
)

func main() {
	helloHandler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Hello, world!\n")
	}

	postArticleHandler := func(w http.ResponseWriter, requ *http.Request) {
		io.WriteString(w, "Posting Articles...\n")
	}

	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/article", helloHandler)

	log.Println("server start at port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
