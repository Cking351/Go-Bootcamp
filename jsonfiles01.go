package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Users struct {
	Users []User `json:"users"`
}

type User struct {
	Name string `json:"name"`
	Type string `json:"type"`
	Age int	    `json:"Age"`
	Social Social`json:"social"`
}

type Social struct {
	Facebook string `json:"facebook"`
	Twitter string `json:"twitter"`
}

func main() {
	// open our json file
	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Successfully Opened users.json")
	// defer the closing of our file so that we can parse it later on
	defer jsonFile.Close()

	// read json file
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we init our users array
	var users Users

	// unmarshal our bytearray
	json.Unmarshal(byteValue, &users)

	// iterate through every user
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
        	fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
        	fmt.Println("User Name: " + users.Users[i].Name)
        	fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
    }

}
