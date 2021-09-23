/* RZFeeser | Alta3 Research
   switch - case and default  */

package main

import (
    "fmt"
    "runtime"
    "strings"
)

func main() {

	var gove string = runtime.Version()
           // init gover                
    switch {
    case strings.Contains(gove, "go1.17"):                 // if matches "go1.17", do the code below then STOP
        fmt.Println("You are using the latest version of GoLang")
    case strings.Contains(gove, "go1.16"), strings.Contains(gove, "go1.15"):
        fmt.Println("You are using an older, but acceptable version of GoLang")
    default:
        fmt.Println("I cannot make a recommendation.")
    }
}

