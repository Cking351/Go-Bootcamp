package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {

	fileServer := http.FileSrver(http.Dir("/public"))
	http.Handle("/", fileServer)

	httpHandleFunc("hello", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello there!\n")
	})

	fmt.Printf("starting server on port 8044\n")
	log.Fatal(http.ListenAndServe(":8044", nil))
}
