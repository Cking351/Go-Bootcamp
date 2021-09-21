/* Pointer rules:

* operator - or dereferecing operator, used to declare pointer variable and access the value stored in the address
& operator - or address operator, used to return the address of a variable or to access the address of a variable to a pointer
A pointer holds the memory address of a value.
The type, *T is a pointer to a T value. Its zero value is nil.
If we want a method to modify its receiver (what is passed to it), we must use a pointer!
If we do not want to alter some data, the value of the receiver can be passed instead of a pointer to its receiver. In this case, the receiver is treated more like an argument.
Pointers are not safe for concurrent access. When you use multiple threads, you will get inconsistent results if you are changing and reading data from the same places in memory.
Unlike C, Go has no pointer arithmetic.*/

package main

import "fmt"

func main() {
	var q int = 42
	var p *int // Declare the pointer
	p = &q // Init the pointer
	fmt.Println(p) // Memory address (0x20e010)

}
