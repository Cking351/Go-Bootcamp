package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	// the handlefun registeres the handler function for the given url pattern
	http.HandleFunc("/", HelloHandler)
	fmt.Println("Server listening on port 8089")
	log.Fatal(http.ListenAndServe(":8089", nil)) // listens on TCP network address
}

// descibes how to respond to the HTTP request
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, there\n")
}

