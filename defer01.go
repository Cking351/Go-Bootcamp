
package main

import (
	"fmt"
	"os"
)

// expects input, "p" to be a string
// will return os.File object
func createFile(p string) *os.File {
	fmt.Println("Creating file...")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("Writing employee list to file...")
	fmt.Fprintln(f, "Cloud strike")
}

func closeFile(f *os.File) {
	fmt.Println("Closing file...")
	err := f.Close()

	// best practice says to still check for errors
	if err != nil {
		fmt.Fprintln(os.Stderr, "error: %v\n", err)
		os.Exit(1)
	}
}

func main() {

	// create file object
	f := createFile("/tmp/employeeList.txt")
	// After getting back a file object, we defer closing closeFile()
	defer closeFile(f) // closeFile() will be executed at end of the enclosing
			// function -- in this case, main()
	writeFile(f)
}
