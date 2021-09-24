package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
)

func main() {

	// define our URL as a string
	url := "https://api.spacexdata.com/v3/capsules"
	// the http method we wish to send
	method := "GET"

	// for control over HTTP client headers,
	// redirect policy, and other settings,
	// creates a Client
	// a Client is an HTTP client
	client := &http.Client{}

	// build a new request
	req, err := http.NewRequest(method, url, nil)

	// error checking
	if err != nil {
		fmt.Println(err)
		return
	}

	// Send the request via a client
	// Do sends an HTTP request and
	// returns an HTTP response
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return 
	}

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
