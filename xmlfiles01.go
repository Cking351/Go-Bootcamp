package main

import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
)

type Notes struct {
	To string `xml:"to"`
	From string `xml:"from"`
	Subject string `xml:"subject"`
	Body string `xml:"body"`
}

func main() {
	data, _ := ioutil.ReadFile("avengers.xml") // open file

	note := &Notes{}

	_ = xml.Unmarshal([]byte(data), &note)

    fmt.Println(note.To)                    // display value of "to"
    fmt.Println(note.From)                  // display value from "from"
    fmt.Println(note.Subject)               // display value from "subject"
    fmt.Println(note.Body)                  // display value from "body"
}
