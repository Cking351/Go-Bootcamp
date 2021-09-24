package main

import (
	"fmt"
	"encoding/json"
	"log"
	"net/http"
	"net/url"
)

func main() {

	data := url.Values{
		"namme":	{"John Doe"},
		"occupation":   {"gardener"},
	}

	resp, err := http.PostForm("https://httpbin.org/post", data)

	if err != nil {
		log.Fatal(err)
	}


	// decode the response body into a map
	var res map[string]interface{}

	json.NewDecoder(resp.Body).Decode(&res)

	fmt.Println(res["form"])
}
