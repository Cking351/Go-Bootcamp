package main

import "fmt"

type Author struct {
	name string
	branch string
	books int
	salary int
}

// method with a reciever
// of author type
func (a Author) show() {
    fmt.Println("Author's Name: ", a.name)
    fmt.Println("Branch Name: ", a.branch)
    fmt.Println("Published articles: ", a.books)
    fmt.Println("Salary: ", a.salary)
}

func (a *Author) bookup() {
	a.books++
}


func main() {
	res := Author {
		name:   	"RZFeeser",
		branch:		"Pennsylvania",
		books:		14,
		salary:		34000,
	}

	res.show()
	res.bookup()
	res.show()
}
