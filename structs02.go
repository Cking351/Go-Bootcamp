package main

import (
	"fmt"
	"os/user"
)

type UserData struct {
	na string // user id
	un string // username
	hd string // home directory
	em string // email

}

func main() {
	user, err := user.Current() // returns user and error
	if err != nil {
		panic(err)
	}

	// create a struct named currentUser
	    currentUser := UserData{user.Uid, user.Username, user.HomeDir, "rzfeeser@example.com"}

	    // recall currentUser
    fmt.Println("User id:", currentUser.na)        // recall current user id
    fmt.Println("Username:", currentUser.un)       // recall username
    fmt.Println("Home Directory:", currentUser.hd) // recall homedirectory
    fmt.Println("Email:", currentUser.em)          // recall email
}
