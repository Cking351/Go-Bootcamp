package main

import (
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	rasp, err := http.Get("http://webcode.me")

	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	_, err = io.Copy(os.Stdout, resp.Body)

	if err != nil {
		log.Fatal(err)
	}
}
